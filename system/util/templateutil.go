package util

import (
	"bytes"
	"errors"
	"go/format"
	"io/ioutil"
	"os"
	"text/template"
)

func LoadTemplateProcessWrite2File(templateFileName string, templateData interface{}, outputFile string, formatSource bool) error {

	if f, err := ioutil.ReadFile(templateFileName); err != nil {
		return err
	} else {
		if pkgTemplate, err := template.New("css").Parse(string(f)); err != nil {
			return err
		} else {
			if err := ProcessTemplateWrite2File(pkgTemplate, templateData, outputFile, formatSource); err != nil {
				return err
			}
		}
	}

	return nil

}

type TemplateInfo struct {
	Name    string
	Content string
}

func ParseTemplates(templates []TemplateInfo, fMaps template.FuncMap) (*template.Template, error) {
	if len(templates) == 0 {
		return nil, errors.New("no template provided")
	}

	mainTemplate := template.New(templates[0].Name)
	if len(fMaps) > 0 {
		mainTemplate = mainTemplate.Funcs(fMaps)
	}

	if mainTemplate, err := mainTemplate.Parse(templates[0].Content); err != nil {
		return nil, err
	} else {
		for i := 1; i < len(templates); i++ {
			if _, err = mainTemplate.New(templates[i].Name).Parse(templates[i].Content); err != nil {
				return nil, err
			}
		}
	}

	return mainTemplate, nil
}

/*
func ParseTemplateProcessWrite2File(templateContent string, templateData interface{}, outputFile string, formatSource bool) error {

	if pkgTemplate, err := template.New("css").Parse(templateContent); err != nil {
		return err
	} else {
		if err := ProcessTemplateWrite2File(pkgTemplate, templateData, outputFile, formatSource); err != nil {
			return err
		}
	}

	return nil
}

func ParseTemplateWithFuncMapsProcessWrite2File(templateContent string, fMaps template.FuncMap, templateData interface{}, outputFile string, formatSource bool) error {

	if pkgTemplate, err := template.New("css").Funcs(fMaps).Parse(templateContent); err != nil {
		return err
	} else {
		if err := ProcessTemplateWrite2File(pkgTemplate, templateData, outputFile, formatSource); err != nil {
			return err
		}
	}

	return nil
}
*/

func ProcessTemplateWrite2File(pkgTemplate *template.Template, templateData interface{}, outputFile string, formatSource bool) error {

	builder := &bytes.Buffer{}

	if err := pkgTemplate.Execute(builder, templateData); err != nil {
		return err
	}

	var data []byte
	if formatSource {
		var err error
		if data, err = format.Source(builder.Bytes()); err != nil {
			return err
		}
	} else {
		data = builder.Bytes()
	}

	if err := ioutil.WriteFile(outputFile, data, os.ModePerm); err != nil {
		return err
	}

	return nil
}
