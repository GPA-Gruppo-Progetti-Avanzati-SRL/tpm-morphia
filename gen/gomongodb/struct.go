package gomongodb

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/gen/gomongodb/attributes"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schema"
	"path/filepath"
)

type StructGeneratorModel struct {
	Package    string
	Definition *schema.StructDef
	Attributes []attributes.GoAttribute
}

func (s *StructGeneratorModel) Name() string {
	return s.Definition.Name
}

func (s *StructGeneratorModel) PackageName() string {
	return filepath.Base(s.Package)
}

func (s *StructGeneratorModel) PackageImports() []string {

	var res []string
	var imports map[string]struct{}

	// Duplications removal
	for _, a := range s.Attributes {
		for _, i := range a.PackageImports() {
			if len(imports) == 0 {
				imports = make(map[string]struct{})
			}
			if _, ok := imports[i]; !ok {
				imports[i] = struct{}{}
				res = append(res, a.PackageImports()...)
			}
		}
	}

	return res
}

func NewStructGeneratorModel(schema *schema.Schema, entityDef *schema.StructDef) (StructGeneratorModel, error) {
	const semLogContext = "go-mongo-db::new-struct-generator-model"

	m := StructGeneratorModel{
		Definition: entityDef,
		Package:    entityDef.Package,
	}

	for _, a := range m.Definition.Attributes {
		goAttr := attributes.NewAttribute(m.Definition, a, false)
		m.Attributes = append(m.Attributes, goAttr)
	}

	return m, nil
}
