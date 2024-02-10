package attributes

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schemaold"
)

type ValueTypeAttribute struct {
	GoAttributeImpl
}

func (v *ValueTypeAttribute) GoType() string {
	var t string

	switch v.AttrDefinition.Typ {
	case schemaold.AttributeTypeString:
		t = AttributeTypeStringGoType
	case schemaold.AttributeTypeInt:
		t = AttributeTypeIntGoType
	case schemaold.AttributeTypeLong:
		t = AttributeTypeLongGoType
	case schemaold.AttributeTypeBool:
		t = AttributeTypeBoolGoType
	case schemaold.AttributeTypeDate:
		t = AttributeTypeDateGoType
	case schemaold.AttributeTypeObjectId:
		t = AttributeTypeObjectIdGoType
	case schemaold.AttributeTypeDocument:
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
