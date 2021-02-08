package mongodb

import (
	"errors"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/config"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schema"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system/resources"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system/util"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"
)

type GenerationContext struct {
	Schema     *schema.Collection
	Collection *CodeGenCollection
}

const (
	TmplCollectionReadme             = "/resources/mongodb-%s/readme.txt"
	TmplCollectionModel              = "/resources/mongodb-%s/model.txt"
	TmplCollectionFilter             = "/resources/mongodb-%s/filter.txt"
	TmplCollectionFilterString       = "/resources/mongodb-%s/filter-string.txt"
	TmplCollectionFilterInt          = "/resources/mongodb-%s/filter-int.txt"
	TmplCollectionStructFilterString = "/resources/mongodb-%s/struct-filter-string.txt"
	TmplCollectionStructFilterInt    = "/resources/mongodb-%s/struct-filter-int.txt"
)

// List of templates for mongoDbGeneration
func filterTmplList(tmplVersion string) []string {
	s := make([]string, 0, 3)
	s = append(s, fmt.Sprintf(TmplCollectionFilter, tmplVersion))
	s = append(s, fmt.Sprintf(TmplCollectionFilterString, tmplVersion))
	s = append(s, fmt.Sprintf(TmplCollectionFilterInt, tmplVersion))
	s = append(s, fmt.Sprintf(TmplCollectionStructFilterString, tmplVersion))
	s = append(s, fmt.Sprintf(TmplCollectionStructFilterInt, tmplVersion))
	return s
}

func readmeTmplList(tmplVersion string) []string {
	s := make([]string, 0, 1)
	s = append(s, fmt.Sprintf(TmplCollectionReadme, tmplVersion))
	return s
}

func modelTmplList(tmplVersion string) []string {
	s := make([]string, 0, 1)
	s = append(s, fmt.Sprintf(TmplCollectionModel, tmplVersion))
	return s
}

func Generate(logger log.Logger, cfg *config.Config, gen *CodeGenCollection) error {

	genFolder, err := gen.GetGeneratedContentPath(cfg.TargetDirectory)
	if err != nil {
		return err
	}

	_ = level.Info(logger).Log("msg", "output directory", "dir", genFolder)

	tmplVersion := "v1"
	if cfg.Version != "" {
		tmplVersion = cfg.Version
	}

	genCtx := GenerationContext{gen.Schema, gen}

	/*
	 * Readme.md
	 */
	if err := emit(logger, genCtx, cfg.ResourceDirectory, genFolder, "readme.md", modelTmplList(tmplVersion), cfg.FormatCode); err != nil {
		return err
	}

	/*
		if t, ok := loadTemplate(cfg.ResourceDirectory, readmeTmplList(tmplVersion)...); ok {
			destinationFile := filepath.Join(genFolder, "readme.md")
			_ = level.Info(logger).Log("msg", "generating text from template", "tmpl", tmplName, "dest", destinationFile)

			if err := ParseTemplateWithFuncMapsProcessWrite2File(t, nil, genCtx, destinationFile, false); err != nil {
				return err
			}
		} else {
			_ = level.Info(logger).Log("msg", "template not present...skipping", "tmpl", tmplName)
		}
	*/

	/*
	 * model.go
	 */
	if err := emit(logger, genCtx, cfg.ResourceDirectory, genFolder, "model.go", modelTmplList(tmplVersion), cfg.FormatCode); err != nil {
		return err
	}

	/*
		if t, ok := loadTemplate(cfg.ResourceDirectory, modelTmplList(tmplVersion)...); ok {
			destinationFile := filepath.Join(genFolder, "model.go")
			_ = level.Info(logger).Log("msg", "generating text from template", "tmpl", tmplName, "dest", destinationFile)

			if err := ParseTemplateWithFuncMapsProcessWrite2File(t, geTemplateUtilityFunctions(), genCtx, destinationFile, cfg.FormatCode); err != nil {
				return err
			}
		} else {
			_ = level.Info(logger).Log("msg", "template not present...skipping", "tmpl", tmplName)
		}
	*/

	/*
	 * filter.go
	 */
	if err := emit(logger, genCtx, cfg.ResourceDirectory, genFolder, "filter.go", filterTmplList(tmplVersion), cfg.FormatCode); err != nil {
		return err
	}

	/*
		if t, ok := loadTemplate(cfg.ResourceDirectory, filterTmplList(tmplVersion)...); ok {
			destinationFile := filepath.Join(genFolder, "filter.go")
			_ = level.Info(logger).Log("msg", "generating text from template", "tmpl", tmplName, "dest", destinationFile)

			if err := ParseTemplateWithFuncMapsProcessWrite2File(t, geTemplateUtilityFunctions(), genCtx, destinationFile, cfg.FormatCode); err != nil {
				return err
			}
		} else {
			_ = level.Info(logger).Log("msg", "template not present...skipping", "tmpl", tmplName)
		}
	*/

	attrs := gen.FindAttributes()
	for _, a := range attrs {
		fmt.Println(a)
	}
	return nil
}

func emit(logger log.Logger, genCtx GenerationContext, resDir string, outFolder string, generatedFileName string, templates []string, formatCode bool) error {
	if t, ok := loadTemplate(resDir, templates...); ok {
		destinationFile := filepath.Join(outFolder, generatedFileName)
		_ = level.Info(logger).Log("msg", "generating text from template", "dest", destinationFile)

		if err := ParseTemplateWithFuncMapsProcessWrite2File(t, geTemplateUtilityFunctions(), genCtx, destinationFile, formatCode); err != nil {
			return err
		}
	} else {
		_ = level.Info(logger).Log("msg", "unable to load template ...skipping")
		return errors.New("unable to load template ...skipping")
	}

	return nil
}

func loadTemplate(resDirectory string, templatePath ...string) ([]util.TemplateInfo, bool) {

	res := make([]util.TemplateInfo, 0)
	for _, tpath := range templatePath {

		/*
		 * Get the name of the template from the file name.... Hope there is one dot only...
		 * Dunno it is a problem.
		 */
		tname := filepath.Base(tpath)
		if ext := filepath.Ext(tname); ext != "" {
			tname = strings.TrimSuffix(tname, ext)
		}

		var b []byte
		var ok bool
		var e error
		if resDirectory == "" {
			if b, ok = resources.Get(tpath); !ok {
				return nil, false
			}
		} else {
			templateFileName := filepath.Join(resDirectory, tpath)
			if b, e = ioutil.ReadFile(templateFileName); e != nil {
				return nil, false
			}
		}

		res = append(res, util.TemplateInfo{Name: tname, Content: string(b)})
	}

	return res, true
}

func geTemplateUtilityFunctions() template.FuncMap {

	fMap := template.FuncMap{
		"formatIdentifier": func(n string, sep string, casingMode util.FormatMode, indexesMode util.FormatMode) string {
			return util.FormatIdentifier(n, sep, casingMode, indexesMode)
		},
		"numberOfArrayIndicesInQualifiedName": func(n string) int {
			return strings.Count(n, "[]")
		},
		"filterSubTemplateContext": func(attribute CodeGenAttribute, criteriaObjectName string) map[string]interface{} {
			return map[string]interface{}{
				"Attr":              attribute,
				"CriteriaStructRef": criteriaObjectName,
			}
		},
	}

	return fMap
}

func ParseTemplateWithFuncMapsProcessWrite2File(templates []util.TemplateInfo, fMaps template.FuncMap, templateData interface{}, outputFile string, formatSource bool) error {

	if pkgTemplate, err := util.ParseTemplates(templates, fMaps); err != nil {
		return err
	} else {
		if err := util.ProcessTemplateWrite2File(pkgTemplate, templateData, outputFile, formatSource); err != nil {
			return err
		}
	}

	return nil
}
