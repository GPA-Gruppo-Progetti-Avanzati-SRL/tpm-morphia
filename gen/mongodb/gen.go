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
	TmplCollectionReadme         = "/resources/mongodb-%s/readme.txt"
	TmplCollectionModel          = "/resources/mongodb-%s/model.txt"
	TmplCollectionFilter         = "/resources/mongodb-%s/filter-methods.txt"
	TmplCollectionCriteria       = "/resources/mongodb-%s/filter.txt"
	TmplCollectionFilterString   = "/resources/mongodb-%s/filter-string.txt"
	TmplCollectionFilterInt      = "/resources/mongodb-%s/filter-int.txt"
	TmplCollectionFilterLong     = "/resources/mongodb-%s/filter-long.txt"
	TmplCollectionFilterBool     = "/resources/mongodb-%s/filter-bool.txt"
	TmplCollectionFilterDate     = "/resources/mongodb-%s/filter-date.txt"
	TmplCollectionFilterObjectId = "/resources/mongodb-%s/filter-object-id.txt"
	TmplCollectionUpdate         = "/resources/mongodb-%s/update.txt"
	TmplCollectionUpdateMethods  = "/resources/mongodb-%s/update-methods.txt"
	TmplCollectionUpdateString   = "/resources/mongodb-%s/update-string.txt"
	TmplCollectionUpdateInt      = "/resources/mongodb-%s/update-int.txt"
	TmplCollectionUpdateLong     = "/resources/mongodb-%s/update-long.txt"
	TmplCollectionUpdateBool     = "/resources/mongodb-%s/update-bool.txt"
	TmplCollectionUpdateDate     = "/resources/mongodb-%s/update-date.txt"
	TmplCollectionUpdateDocument = "/resources/mongodb-%s/update-document.txt"
	TmplCollectionUpdateObjectId = "/resources/mongodb-%s/update-object-id.txt"
)

// List of templates for mongoDbGeneration
func filterTmplList(tmplVersion string) []string {
	s := make([]string, 0, 3)
	s = append(s, fmt.Sprintf(TmplCollectionFilter, tmplVersion))
	s = append(s, fmt.Sprintf(TmplCollectionFilterString, tmplVersion))
	s = append(s, fmt.Sprintf(TmplCollectionFilterInt, tmplVersion))
	s = append(s, fmt.Sprintf(TmplCollectionFilterDate, tmplVersion))
	s = append(s, fmt.Sprintf(TmplCollectionFilterLong, tmplVersion))
	s = append(s, fmt.Sprintf(TmplCollectionFilterBool, tmplVersion))
	s = append(s, fmt.Sprintf(TmplCollectionFilterObjectId, tmplVersion))
	return s
}

func criteriaTmplList(tmplVersion string) []string {
	s := make([]string, 0, 1)
	s = append(s, fmt.Sprintf(TmplCollectionCriteria, tmplVersion))
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

func updateTmplList(tmplVersion string) []string {
	s := make([]string, 0, 1)
	s = append(s, fmt.Sprintf(TmplCollectionUpdate, tmplVersion))
	return s
}

func updateMethodsTmplList(tmplVersion string) []string {
	s := make([]string, 0, 1)
	s = append(s, fmt.Sprintf(TmplCollectionUpdateMethods, tmplVersion))
	s = append(s, fmt.Sprintf(TmplCollectionUpdateString, tmplVersion))
	s = append(s, fmt.Sprintf(TmplCollectionUpdateInt, tmplVersion))
	s = append(s, fmt.Sprintf(TmplCollectionUpdateLong, tmplVersion))
	s = append(s, fmt.Sprintf(TmplCollectionUpdateBool, tmplVersion))
	s = append(s, fmt.Sprintf(TmplCollectionUpdateDate, tmplVersion))
	s = append(s, fmt.Sprintf(TmplCollectionUpdateDocument, tmplVersion))
	s = append(s, fmt.Sprintf(TmplCollectionUpdateObjectId, tmplVersion))
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
	if err := emit(logger, genCtx, cfg.ResourceDirectory, genFolder, getOutputFilename(gen.GetPrefix("lower"), "readme.md"), readmeTmplList(tmplVersion), false); err != nil {
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
	if err := emit(logger, genCtx, cfg.ResourceDirectory, genFolder, getOutputFilename(gen.GetPrefix("lower"), "model.go"), modelTmplList(tmplVersion), cfg.FormatCode); err != nil {
		return err
	}

	/*
		if t, ok := loadTemplate(cfg.ResourceDirectory, modelTmplList(tmplVersion)...); ok {
			destinationFile := filepath.Join(genFolder, "model.go")
			_ = level.Info(logger).Log("msg", "generating text from template", "tmpl", tmplName, "dest", destinationFile)

			if err := ParseTemplateWithFuncMapsProcessWrite2File(t, getTemplateUtilityFunctions(), genCtx, destinationFile, cfg.FormatCode); err != nil {
				return err
			}
		} else {
			_ = level.Info(logger).Log("msg", "template not present...skipping", "tmpl", tmplName)
		}
	*/

	/*
	 * filter-methods.go
	 */
	if err := emit(logger, genCtx, cfg.ResourceDirectory, genFolder, getOutputFilename(gen.GetPrefix("lower"), "filter-methods.go"), filterTmplList(tmplVersion), cfg.FormatCode); err != nil {
		return err
	}

	/*
	 * filter.go: get generated without prefixing because is common across any collection that might be created.
	 */
	if err := emit(logger, genCtx, cfg.ResourceDirectory, genFolder, getOutputFilename("", "filter.go"), criteriaTmplList(tmplVersion), cfg.FormatCode); err != nil {
		return err
	}

	/*
	 * Update get generated without prefixing because is common across any collection that might be created.
	 */
	if err := emit(logger, genCtx, cfg.ResourceDirectory, genFolder, getOutputFilename("", "update.go"), updateTmplList(tmplVersion), cfg.FormatCode); err != nil {
		return err
	}

	/*
	 * update-methods.go
	 */
	if err := emit(logger, genCtx, cfg.ResourceDirectory, genFolder, getOutputFilename(gen.GetPrefix("lower"), "update-methods.go"), updateMethodsTmplList(tmplVersion), cfg.FormatCode); err != nil {
		return err
	}

	return nil
}

func emit(logger log.Logger, genCtx GenerationContext, resDir string, outFolder string, generatedFileName string, templates []string, formatCode bool) error {
	if t, ok := loadTemplate(resDir, templates...); ok {
		destinationFile := filepath.Join(outFolder, generatedFileName)
		_ = level.Info(logger).Log("msg", "generating text from template", "dest", destinationFile)

		if err := parseTemplateWithFuncMapsProcessWrite2File(t, getTemplateUtilityFunctions(), genCtx, destinationFile, formatCode); err != nil {
			_ = level.Info(logger).Log("msg", "parse template failed", "err", err.Error())
			return err
		}
	} else {
		_ = level.Info(logger).Log("msg", "unable to load template ...skipping")
		return errors.New("unable to load template ...skipping")
	}

	return nil
}

func getOutputFilename(prefix string, baseName string) string {
	if prefix != "" {
		return strings.Join([]string{prefix, baseName}, "-")
	}

	return baseName
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

func getTemplateUtilityFunctions() template.FuncMap {

	fMap := template.FuncMap{
		"formatIdentifier": func(n string, sep string, casingMode util.FormatMode, indexesMode util.FormatMode, indexesFormatMode util.FormatMode) string {
			return util.FormatIdentifier(n, sep, casingMode, indexesMode, indexesFormatMode)
		},
		"numberOfArrayIndicesInQualifiedName": func(n string) int {
			return strings.Count(n, "[]") + strings.Count(n, "%s")
		},
		"filterSubTemplateContext": func(attribute CodeGenAttribute, criteriaObjectName string) map[string]interface{} {
			return map[string]interface{}{
				"Attr":              attribute,
				"CriteriaStructRef": criteriaObjectName,
			}
		},
		"isIdentifierIndexed": func(n string) bool {
			return strings.Contains(n, "[]")
		},
		"firstToLower": func(n string) string {
			return util.FirstToLower(n)
		},
		"criteriaMethodSignature": func(p string) string {
			// Do a Camel case conversion with segments separated by '.'
			s := util.FormatIdentifier(p, ".", "camelCase", "index", "indexIjk")
			// Remove the period.
			return strings.ReplaceAll(s, ".", "")
		},
		"criteriaMethodVarParams": func(p string, withType bool, commaHandling string) string {

			if strings.Contains(p, "[]") || strings.Contains(p, "%s") {

				var sb strings.Builder

				arr := strings.Split(p, ".")

				numIjk := 0
				numStw := 0
				for _, s := range arr {
					if s == "[]" {
						if (numIjk + numStw) > 0 {
							sb.WriteString(", ")
						}
						sb.WriteString("ndx")
						sb.WriteRune(rune('I' + numIjk))
						if withType {
							sb.WriteString(" int")
						}
						numIjk++
					} else if s == "%s" {
						if (numIjk + numStw) > 0 {
							sb.WriteString(", ")
						}
						sb.WriteString("key")
						sb.WriteRune(rune('S' + numStw))
						if withType {
							sb.WriteString(" string")
						}
						numStw++
					}
				}

				if strings.Contains(commaHandling, "before") {
					return ", " + sb.String()
				} else if strings.Contains(commaHandling, "after") {
					return sb.String() + ", "
				}

				return sb.String()
			}

			if strings.Contains(commaHandling, "addonempty") {
				return ", "
			}

			return ""
		},
		"updateMethodSignature": func(p string) string {
			// Do a Camel case conversion with segments separated by '.'
			s := util.FormatIdentifier(p, ".", "camelCase", "index", "indexIjk")
			// Remove the period.
			return strings.ReplaceAll(s, ".", "")
		},
		"updateMethodVarParams": func(p string, withType bool, commaHandling string) string {

			if strings.Contains(p, "[]") || strings.Contains(p, "%s") {

				var sb strings.Builder

				arr := strings.Split(p, ".")

				numIjk := 0
				numStw := 0
				for _, s := range arr {
					if s == "[]" {
						if (numIjk + numStw) > 0 {
							sb.WriteString(", ")
						}
						sb.WriteString("ndx")
						sb.WriteRune(rune('I' + numIjk))
						if withType {
							sb.WriteString(" int")
						}
						numIjk++
					} else if s == "%s" {
						if (numIjk + numStw) > 0 {
							sb.WriteString(", ")
						}
						sb.WriteString("key")
						sb.WriteRune(rune('S' + numStw))
						if withType {
							sb.WriteString(" string")
						}
						numStw++
					}
				}

				if strings.Contains(commaHandling, "before") {
					return ", " + sb.String()
				} else if strings.Contains(commaHandling, "after") {
					return sb.String() + ", "
				}

				return sb.String()
			}

			if strings.Contains(commaHandling, "addonempty") {
				return ", "
			}

			return ""
		},
	}

	return fMap
}

func parseTemplateWithFuncMapsProcessWrite2File(templates []util.TemplateInfo, fMaps template.FuncMap, templateData interface{}, outputFile string, formatSource bool) error {

	if pkgTemplate, err := util.ParseTemplates(templates, fMaps); err != nil {
		return err
	} else {
		if err := util.ProcessTemplateWrite2File(pkgTemplate, templateData, outputFile, formatSource); err != nil {
			return err
		}
	}

	return nil
}
