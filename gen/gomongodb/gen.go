package gomongodb

import (
	"embed"
	"errors"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/gen/gomongodb/attributes"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schema"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-schematics/schematics"
	"github.com/rs/zerolog/log"
	"path/filepath"
	"strings"
	"text/template"
)

type GeneratorConfig struct {
	Schema               *schema.Schema
	TargetFolder         string
	EntityName           string
	Version              string
	ConflictModeHandling string
	FormatCode           bool
}

//go:embed templates-pogs/*
var templatesPogs embed.FS

//go:embed templates-doc/*
var templatesDoc embed.FS

func GenerateAllEntities(cfg *GeneratorConfig) error {
	const semLogContext = "go-mongo-db::generate-all-entities"

	for _, s := range cfg.Schema.Structs {
		cfg.EntityName = s.Name
		err := GenerateEntity(cfg)
		if err != nil {
			log.Error().Err(err).Msg(semLogContext)
			return err
		}
	}

	return nil
}

func GenerateEntity(cfg *GeneratorConfig) error {
	const semLogContext = "go-mongo-db::generate-entity"

	var err error

	entityDef := cfg.Schema.GetStructByName(cfg.EntityName)
	if entityDef == nil {
		err = errors.New("cannot find struct in schema")
		log.Error().Err(err).Str("struct-name", cfg.EntityName).Msg(semLogContext)
		return err
	}

	if entityDef.IsDocument {
		// Traverse the structure to wire the paths.
		pv := schema.PathFinderVisitor{}
		cfg.Schema.VisitStruct(entityDef.Name, &pv)
	}

	metadata := map[string]interface{}{
		"name":    cfg.EntityName,
		"version": cfg.Version,
	}

	sourceTemplateOptions := []schematics.SourceTemplateOption{
		schematics.SourceWithFuncMap(getFuncMap()),
		schematics.SourceWithMetadata(metadata),
	}

	if cfg.FormatCode {
		sourceTemplateOptions = append(sourceTemplateOptions, schematics.SourceWithFormatCode())
	}

	var modelPkg string
	templates := templatesPogs
	templatseRootFolder := "templates-pogs"
	if entityDef.IsDocument {
		var docModel DocumentGeneratorModel
		docModel, err = NewDocumentGeneratorModel(cfg.Schema, entityDef)
		modelPkg = docModel.Package
		templates = templatesDoc
		templatseRootFolder = "templates-doc"
		sourceTemplateOptions = append(sourceTemplateOptions, schematics.SourceWithModel(&docModel))

		// Visitor snippet...
		visitor := attributes.LogVisitor{}
		for _, a := range docModel.Attributes {
			a.Visit(&visitor)
		}

	} else {
		var structModel StructGeneratorModel
		structModel, err = NewStructGeneratorModel(cfg.Schema, entityDef)
		modelPkg = structModel.Package
		sourceTemplateOptions = append(sourceTemplateOptions, schematics.SourceWithModel(&structModel))
	}

	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return err
	}

	src, err := schematics.GetSource(
		templates, templatseRootFolder,
		sourceTemplateOptions...,
	)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return err
	}

	pathOfStruct := strings.TrimPrefix(modelPkg, cfg.Schema.ModuleName)
	if pathOfStruct == modelPkg {
		err = errors.New("package name is not part of the module")
		log.Error().Err(err).Str("struct-package", modelPkg).Str("module", cfg.Schema.ModuleName).Msg(semLogContext)
		return err
	}

	err = schematics.Apply(
		filepath.Join(cfg.TargetFolder, pathOfStruct),
		src,
		schematics.WithApplyDefaultConflictMode(schematics.ConflictModeOverwrite),
		schematics.WithApplyProduceDiff())
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return err
	}

	return nil
}

func getFuncMap() template.FuncMap {
	fMap := template.FuncMap{
		"camelize": func(s string) string {
			return util.Camelize(s)
		},
		"dasherize": func(s string) string {
			return util.Dasherize(s)
		},
		"classify": func(s string) string {
			return util.Classify(s)
		},
		"decamelize": func(s string) string {
			return util.Decamelize(s)
		},
		"underscore": func(s string) string {
			return util.Underscore(s)
		},
		"uuid": func() string {
			return util.NewUUID()
		},

		"formatIdentifier": func(n string, sep string, casingMode FormatMode, indexesMode FormatMode, indexesFormatMode FormatMode) string {
			return FormatIdentifier(n, sep, casingMode, indexesMode, indexesFormatMode)
		},
		"numberOfArrayIndicesInQualifiedName": func(n string) int {
			return strings.Count(n, "[]") + strings.Count(n, "%s")
		},
		"filterSubTemplateContext": func(attribute attributes.GoAttribute, currentPackage string) map[string]interface{} {
			return map[string]interface{}{
				"Attr":           attribute,
				"CurrentPackage": currentPackage,
			}
		},
		"isIdentifierIndexed": func(n string) bool {
			return strings.Contains(n, "[]")
		},
		"firstToLower": func(n string) string {
			return FirstToLower(n)
		},
		"criteriaMethodSignature": func(p string) string {
			// Do a Camel case conversion with segments separated by '.'
			s := FormatIdentifier(p, ".", "camelCase", "index", "indexIjk")
			// Remove the period.
			return strings.ReplaceAll(s, ".", "")
		},
		"criteriaMethodVarParams": func(p string, withType bool, commaHandling string) string {

			return criteriaMethodVarParams(p, withType, commaHandling)
		},
		"updateMethodSignature": func(p string) string {
			// Do a Camel case conversion with segments separated by '.'
			s := FormatIdentifier(p, ".", "camelCase", "index", "indexIjk")
			// Remove the period.
			return strings.ReplaceAll(s, ".", "")
		},
		"updateMethodVarParams": func(p string, withType bool, commaHandling string) string {
			return criteriaMethodVarParams(p, withType, commaHandling)
		},
	}

	return fMap
}
