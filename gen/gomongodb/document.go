package gomongodb

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/gen/gomongodb/attributes"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schema"
	"github.com/rs/zerolog/log"
	"path/filepath"
)

type DocumentGeneratorModel struct {
	Package    string
	Definition *schema.StructDef
	Attributes []attributes.GoAttribute
}

func (s *DocumentGeneratorModel) Name() string {
	return s.Definition.Name
}

func (s *DocumentGeneratorModel) PackageName() string {
	return filepath.Base(s.Package)
}

func (s *DocumentGeneratorModel) PackageImports() []string {
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

type AttributesTreeVisitor struct {
	structs map[string]struct{}
	attrs   []attributes.GoAttribute
}

func (av *AttributesTreeVisitor) StartVisit(a attributes.GoAttribute) {
	const semLogContext = "go-attribute-tree-visitor::start-visit"
	log.Trace().Str("name", a.GoName()).Msg(semLogContext)
	// lv.attrs = append(lv.attrs, a)
}

func (av *AttributesTreeVisitor) Visit(a attributes.GoAttribute) bool {
	const semLogContext = "go-attribute-tree-visitor::visit"
	log.Trace().Str("name", a.GoName()).Msg(semLogContext)
	av.attrs = append(av.attrs, a)
	qn := a.StructQualifiedName()
	rc := true
	if qn != "" {
		if len(av.structs) == 0 {
			av.structs = make(map[string]struct{})
			av.structs[qn] = struct{}{}
		} else {
			if _, ok := av.structs[qn]; !ok {
				av.structs[qn] = struct{}{}
			} else {
				rc = false
			}
		}
	}

	return rc
}

func (av *AttributesTreeVisitor) EndVisit(a attributes.GoAttribute) {
	const semLogContext = "go-attribute-tree-visitor::end-visit"
	log.Trace().Str("name", a.GoName()).Msg(semLogContext)
	// lv.attrs = append(lv.attrs, a)
}

func (s *DocumentGeneratorModel) AttributesTree() []attributes.GoAttribute {

	v := AttributesTreeVisitor{}
	// Duplications removal
	for _, a := range s.Attributes {
		if v.Visit(a) {
			a.Visit(&v)
		}

		//attrs = append(attrs, a)
		//children := a.ChildrenAttrs()
		//if len(children) != 0 {
		//	attrs = append(attrs, children...)
		//}
	}

	return v.attrs
}

func NewDocumentGeneratorModel(sch *schema.Schema, entityDef *schema.StructDef) (DocumentGeneratorModel, error) {
	const semLogContext = "go-mongo-db::new-document-generator-model"

	m := DocumentGeneratorModel{
		Definition: entityDef,
		Package:    entityDef.Package,
	}

	for _, a := range m.Definition.Attributes {
		goAttr := attributes.NewAttribute(m.Definition, a, true)
		m.Attributes = append(m.Attributes, goAttr)
	}

	return m, nil
}
