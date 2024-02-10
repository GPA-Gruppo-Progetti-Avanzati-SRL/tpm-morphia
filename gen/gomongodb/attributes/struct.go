package attributes

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"path/filepath"
	"strings"
)

type StructAttribute struct {
	GoAttributeImpl
	Children []GoAttribute
}

/*func (s *StructAttribute) StructQualifiedName() string {
	var qn strings.Builder
	if s.ExternalPackage != "" {
		qn.WriteString(s.ExternalPackage)
		qn.WriteString("/")
	} else {
		qn.WriteString("./")
	}

	qn.WriteString(util.Classify(s.AttrDefinition.StructRef.StructDefRef.Name))
	return qn.String()
}*/

func (s *StructAttribute) GoType() string {
	if s.ExternalPackage != "" {
		var sb strings.Builder
		sb.WriteString(filepath.Base(s.ExternalPackage))
		sb.WriteString(".")
		sb.WriteString(util.Classify(s.AttrDefinition.StructRef.Name))
		return sb.String()
	}

	return util.Classify(s.AttrDefinition.StructRef.Name)
}

func (s *StructAttribute) ChildrenAttrs() []GoAttribute {
	return s.Children
}

func (s *StructAttribute) Visit(visitor Visitor) {
	visitor.StartVisit(s)

	for _, a := range s.Children {
		if visitor.Visit(a) {
			a.Visit(visitor)
		}
	}

	/*
		// Se il riferimento e' esterno allora non lo 'seguo'. L'elemento e' nil.
		if f.Typ == AttributeTypeRefStruct {
			if f.StructRef.Item != nil {
				for i := range f.StructRef.Item.Attributes {
					v.visit(&f.StructRef.Item.Attributes[i])
					f.StructRef.Item.Attributes[i].visit(v)
				}
			}
		} else {
			for i := range f.Attributes {
				v.visit(&f.Attributes[i])
				f.Attributes[i].visit(v)
			}

			if f.Item != nil {
				f.Item.visit(v)
			}
		}
	*/

	visitor.EndVisit(s)
}
