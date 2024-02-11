package attributes

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schema"
)

type ValueTypeAttribute struct {
	GoAttributeImpl
}

func (v *ValueTypeAttribute) GoType() string {
	var t string

	switch v.AttrDefinition.Typ {
	case schema.AttributeTypeString:
		t = AttributeTypeStringGoType
	case schema.AttributeTypeInt:
		t = AttributeTypeIntGoType
	case schema.AttributeTypeLong:
		t = AttributeTypeLongGoType
	case schema.AttributeTypeBool:
		t = AttributeTypeBoolGoType
	case schema.AttributeTypeDate:
		t = AttributeTypeDateGoType
	case schema.AttributeTypeObjectId:
		t = AttributeTypeObjectIdGoType
	case schema.AttributeTypeDocument:
		t = AttributeTypeDocumentGoType
	default:
		t = "-- NotYetImplemented --"
	}
	return t
}

func (v *ValueTypeAttribute) Visit(visitor Visitor) {
	visitor.StartVisit(v)
	visitor.EndVisit(v)
}

func (v *ValueTypeAttribute) StructQualifiedName() string {
	return ""
}
