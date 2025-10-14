package attributes

import (
	"strings"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schema"
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
		v.Imports = append(v.Imports, "go.mongodb.org/mongo-driver/v2/bson")

	case schema.AttributeTypeObjectId:
		v.Imports = make([]string, 0, 1)
		v.Imports = append(v.Imports, "go.mongodb.org/mongo-driver/v2/bson")

	case schema.AttributeTypeDocument:
		v.Imports = make([]string, 0, 1)
		v.Imports = append(v.Imports, "go.mongodb.org/mongo-driver/v2/bson")

	case schema.AttributeTypeTimestamp:
		v.Imports = make([]string, 0, 1)
		v.Imports = append(v.Imports, "go.mongodb.org/mongo-driver/v2/bson")
	}

	return v
}

func NewStructAttribute(parentStruct *schema.StructDef, attrDefinition *schema.Field, recurse bool) GoAttribute {
	s := &StructAttribute{}
	s.AttrDefinition = attrDefinition

	sn, pkg := attrDefinition.Package()
	s.StructTypeName = sn
	s.Pkg = pkg
	if pkg != "" {
		s.Imports = make([]string, 0, 1)
		s.Imports = append(s.Imports, s.Pkg)
	}
	/*
		if pkg != "" && pkg != parentStruct.Package {
			s.ExternalPackage = pkg
		}
	*/

	/*
		if s.ExternalPackage != "" {
			s.Imports = make([]string, 0, 1)
			s.Imports = append(s.Imports, s.ExternalPackage)
		}
	*/

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
	a.Pkg = pkg
	if a.Pkg != "" {
		a.Imports = make([]string, 0, 1)
		a.Imports = append(a.Imports, a.Pkg)
	}
	/*
		if pkg != "" && pkg != parentStruct.Package {
			a.ExternalPackage = pkg
		}

		if a.ExternalPackage != "" {
			a.Imports = make([]string, 0, 1)
			a.Imports = append(a.Imports, a.ExternalPackage)
		}
	*/
	return a
}

func NewMapAttribute(parentStruct *schema.StructDef, attrDefinition *schema.Field, recurse bool) GoAttribute {
	a := &MapAttribute{}
	a.AttrDefinition = attrDefinition
	a.Item = NewAttribute(parentStruct, attrDefinition.Item, recurse)

	sn, pkg := attrDefinition.Package()
	a.StructTypeName = sn
	a.Pkg = pkg
	if a.Pkg != "" {
		a.Imports = make([]string, 0, 1)
		a.Imports = append(a.Imports, a.Pkg)
	}
	/*
		if pkg != "" && pkg != parentStruct.Package {
			a.ExternalPackage = pkg
		}

		if a.ExternalPackage != "" {
			a.Imports = make([]string, 0, 1)
			a.Imports = append(a.Imports, a.ExternalPackage)
		}
	*/

	return a
}
