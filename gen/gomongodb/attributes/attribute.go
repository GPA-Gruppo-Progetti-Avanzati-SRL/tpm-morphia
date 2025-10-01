package attributes

import (
	"errors"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schema"
	"github.com/rs/zerolog/log"
	"strings"
)

const (
	AttributeTypeStringGoType    = "string"
	AttributeTypeIntGoType       = "int32"
	AttributeTypeLongGoType      = "int64"
	AttributeTypeBoolGoType      = "bool"
	AttributeTypeObjectIdGoType  = "primitive.ObjectID"
	AttributeTypeDateGoType      = "primitive.DateTime"
	AttributeTypeTimestampGoType = "primitive.Timestamp"
	AttributeTypeDocumentGoType  = "bson.M"
)

type PathInfo struct {
	Path     string
	BsonPath string
}

type GoAttribute interface {
	StructQualifiedName() string
	Package() string
	Definition() *schema.Field
	BsonPropertyName(qualified bool) string
	PackageImports(currentPkg string) []string
	GoName() string
	GoType(currentPkg string) string
	GoTag() string
	GoIsZeroCondition() string
	ChildrenAttrs() []GoAttribute
	Paths() []PathInfo
	Visit(v Visitor)
}

type GoAttributeImpl struct {
	AttrDefinition *schema.Field
	// ExternalPackage string
	Pkg            string
	StructTypeName string
	Imports        []string
}

func (cg *GoAttributeImpl) Paths() []PathInfo {

	var pis []PathInfo
	for i, p := range cg.AttrDefinition.Paths {
		pis = append(pis, PathInfo{
			Path:     p,
			BsonPath: cg.AttrDefinition.BsonPaths[i],
		})
	}
	return pis
}

type Visitor interface {
	StartVisit(f GoAttribute)
	Visit(f GoAttribute) bool
	EndVisit(f GoAttribute)
}

type LogVisitor struct {
}

func (lv *LogVisitor) StartVisit(a GoAttribute) {
	const semLogContext = "go-attribute-log-visitor::start-visit"
	log.Trace().Str("name", a.GoName()).Msg(semLogContext)
}

func (lv *LogVisitor) Visit(a GoAttribute) bool {
	const semLogContext = "go-attribute-log-visitor::visit"
	log.Trace().Str("name", a.GoName()).Msg(semLogContext)
	return true
}

func (lv *LogVisitor) EndVisit(a GoAttribute) {
	const semLogContext = "go-attribute-log-visitor::end-visit"
	log.Trace().Str("name", a.GoName()).Msg(semLogContext)
}

func (cg *GoAttributeImpl) ChildrenAttrs() []GoAttribute {
	return nil
}

func (cg *GoAttributeImpl) StructQualifiedName() string {

	if cg.StructTypeName == "" {
		return ""
	}

	var qn strings.Builder
	if cg.Pkg != "" {
		qn.WriteString(cg.Pkg /* ExternalPackage */)
		qn.WriteString("/")
	} else {
		qn.WriteString("" /* "./" */)
	}

	qn.WriteString(util.Classify(cg.StructTypeName))
	return qn.String()
}

func (cg *GoAttributeImpl) Package() string {
	return cg.Pkg
}

func (cg *GoAttributeImpl) Definition() *schema.Field {
	return cg.AttrDefinition
}

func (b *GoAttributeImpl) HasOption(o string) bool {
	return b.AttrDefinition.Options != "" && strings.Contains(b.AttrDefinition.Options, o)
}

func (b *GoAttributeImpl) GoIsNOTZeroCondition() string {

	s := ""
	switch b.AttrDefinition.Typ {
	case schema.AttributeTypeObjectId:
		s = fmt.Sprintf("s.%s != primitive.NilObjectID", b.GoName())
	case schema.AttributeTypeInt:
		s = fmt.Sprintf("s.%s != 0", b.GoName())
	case schema.AttributeTypeLong:
		s = fmt.Sprintf("s.%s != 0", b.GoName())
	case schema.AttributeTypeBool:
		s = fmt.Sprintf("s.%s", b.GoName())
	case schema.AttributeTypeDate:
		s = fmt.Sprintf("s.%s != 0", b.GoName())
	case schema.AttributeTypeTimestamp:
		s = fmt.Sprintf("s.%s.T != 0", b.GoName())
	case schema.AttributeTypeArray:
		s = fmt.Sprintf("len(s.%s) != 0", b.GoName())
	case schema.AttributeTypeStruct:
		s = fmt.Sprintf("!s.%s.IsZero()", b.GoName())
	case schema.AttributeTypeString:
		s = fmt.Sprintf("s.%s != \"\"", b.GoName())
	case schema.AttributeTypeMap:
		s = fmt.Sprintf("len(s.%s) != 0", b.GoName())
	case schema.AttributeTypeDocument:
		s = fmt.Sprintf("len(s.%s) != 0", b.GoName())
	}

	return s
}

func (b *GoAttributeImpl) GoIsZeroCondition() string {

	s := ""
	switch b.AttrDefinition.Typ {
	case schema.AttributeTypeObjectId:
		s = fmt.Sprintf("s.%s == primitive.NilObjectID", b.GoName())
	case schema.AttributeTypeInt:
		s = fmt.Sprintf("s.%s == 0", b.GoName())
	case schema.AttributeTypeLong:
		s = fmt.Sprintf("s.%s == 0", b.GoName())
	case schema.AttributeTypeBool:
		s = fmt.Sprintf("!s.%s", b.GoName())
	case schema.AttributeTypeDate:
		s = fmt.Sprintf("s.%s == 0", b.GoName())
	case schema.AttributeTypeTimestamp:
		s = fmt.Sprintf("s.%s.T == 0", b.GoName())
	case schema.AttributeTypeArray:
		s = fmt.Sprintf("len(s.%s) == 0", b.GoName())
	case schema.AttributeTypeStruct:
		s = fmt.Sprintf("s.%s.IsZero()", b.GoName())
	case schema.AttributeTypeString:
		s = fmt.Sprintf("s.%s == \"\"", b.GoName())
	case schema.AttributeTypeMap:
		s = fmt.Sprintf("len(s.%s) == 0", b.GoName())
	case schema.AttributeTypeDocument:
		s = fmt.Sprintf("len(s.%s) == 0", b.GoName())
	}

	return s
}

func (b *GoAttributeImpl) PackageImports(currentPkg string) []string {
	var externalPkgs []string
	for _, pkg := range b.Imports {
		if pkg != currentPkg {
			externalPkgs = append(externalPkgs, pkg)
		}
	}
	return externalPkgs
}

func (b *GoAttributeImpl) BsonPropertyName(qualified bool) string {
	tag, _ := b.AttrDefinition.GetTagValueByName("bson")
	//if qualified {
	//	return util.AppendToNamespace(b.BSONNamespace, s, ".")
	//}

	return tag
}

func (b *GoAttributeImpl) GetPaths(pathType string) []string {

	if len(b.AttrDefinition.Paths) > 0 {
		arr := make([]string, 0, len(b.AttrDefinition.Paths))
		isIndexed := false
		for _, p1 := range b.AttrDefinition.Paths {
			if strings.Contains(p1, "[]") {
				isIndexed = true
				if !strings.Contains(pathType, "omitIndexed") {
					arr = append(arr, p1)
				}
			} else {
				arr = append(arr, p1)
			}
		}

		if isIndexed && strings.Contains(pathType, "preferIndexed") {
			arr = make([]string, 0, len(b.AttrDefinition.Paths))
			for _, p1 := range b.AttrDefinition.Paths {
				if strings.Contains(p1, "[]") {
					arr = append(arr, p1)
				}
			}
		}

		return arr
	}

	return nil
}

func (b *GoAttributeImpl) Name(qualified bool, prefixed bool) string {
	s := b.AttrDefinition.Name
	//if qualified {
	//	s = util.AppendToNamespace(b.AttributeNamespace, s, ".")
	//}

	return s
}

func (b *GoAttributeImpl) String() string {
	return fmt.Sprintf("%s of type: %s and ns: %s", b.AttrDefinition.Name, b.AttrDefinition.Typ)
}

func (b *GoAttributeImpl) GoType(currentPkg string) string {
	panic(errors.New("GoAttributeImpl does not implement GetGoAttributeType"))
}

func (b *GoAttributeImpl) GoName() string {
	return util.Classify(b.AttrDefinition.Name)
}

func (b *GoAttributeImpl) GoTag() string {

	if len(b.AttrDefinition.Tags) > 0 {
		var stb strings.Builder
		stb.WriteRune('`')
		for i, t := range b.AttrDefinition.Tags {
			if i > 0 {
				stb.WriteRune(' ')
			}
			stb.WriteString(t.String())
		}

		stb.WriteRune('`')
		return stb.String()
	}

	return ""
}
