package attributes

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schema"
	"strings"
)

func NewAttribute(structDef *schema.StructDef, attrDefinition *schema.Field, recurse bool) GoAttribute {

	var a GoAttribute
	switch attrDefinition.Typ {
	case schema.AttributeTypeStruct:
		a = NewStructAttribute(structDef, attrDefinition, recurse)
	case schema.AttributeTypeArray:
		a = NewArrayAttribute(structDef, attrDefinition, recurse)
	case schema.AttributeTypeMap:
		a = NewMapAttribute(structDef, attrDefinition, recurse)
	default:
		a = NewValueTypeAttribute(structDef, attrDefinition)
	}

	return a
}

func NewValueTypeAttribute(parentStruct *schema.StructDef, attrDefinition *schema.Field) GoAttribute {
	v := &ValueTypeAttribute{}
	v.AttrDefinition = attrDefinition

	switch strings.ToLower(attrDefinition.Typ) {
	case schema.AttributeTypeDate:
		v.Imports = make([]string, 0, 1)
		v.Imports = append(v.Imports, "go.mongodb.org/mongo-driver/bson/primitive")

	case schema.AttributeTypeObjectId:
		v.Imports = make([]string, 0, 1)
		v.Imports = append(v.Imports, "go.mongodb.org/mongo-driver/bson/primitive")

	case schema.AttributeTypeDocument:
		v.Imports = make([]string, 0, 1)
		v.Imports = append(v.Imports, "go.mongodb.org/mongo-driver/bson")
	}

	return v
}

func NewStructAttribute(parentStruct *schema.StructDef, attrDefinition *schema.Field, recurse bool) GoAttribute {
	s := &StructAttribute{}
	s.AttrDefinition = attrDefinition

	sn, pkg := attrDefinition.Package()
	s.StructTypeName = sn
	if pkg != "" && pkg != parentStruct.Package {
		s.ExternalPackage = pkg
	}

	if s.ExternalPackage != "" {
		s.Imports = make([]string, 0, 1)
		s.Imports = append(s.Imports, s.ExternalPackage)
	}

	for _, member := range attrDefinition.StructRef.StructDefRef.Attributes {
		child := NewAttribute(attrDefinition.StructRef.StructDefRef, member, recurse)
		s.Children = append(s.Children, child)
	}
	return s
}

func NewArrayAttribute(parentStruct *schema.StructDef, attrDefinition *schema.Field, recurse bool) GoAttribute {
	a := &ArrayAttribute{}
	a.AttrDefinition = attrDefinition
	a.Item = NewAttribute(parentStruct, attrDefinition.Item, recurse)

	sn, pkg := attrDefinition.Package()
	a.StructTypeName = sn
	if pkg != "" && pkg != parentStruct.Package {
		a.ExternalPackage = pkg
	}

	if a.ExternalPackage != "" {
		a.Imports = make([]string, 0, 1)
		a.Imports = append(a.Imports, a.ExternalPackage)
	}

	return a
}

func NewMapAttribute(parentStruct *schema.StructDef, attrDefinition *schema.Field, recurse bool) GoAttribute {
	a := &MapAttribute{}
	a.AttrDefinition = attrDefinition
	a.Item = NewAttribute(parentStruct, attrDefinition.Item, recurse)

	sn, pkg := attrDefinition.Package()
	a.StructTypeName = sn
	if pkg != "" && pkg != parentStruct.Package {
		a.ExternalPackage = pkg
	}

	if a.ExternalPackage != "" {
		a.Imports = make([]string, 0, 1)
		a.Imports = append(a.Imports, a.ExternalPackage)
	}

	return a
}
