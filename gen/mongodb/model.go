package mongodb

import (
	"errors"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schema"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system/util"
	"os"
	"path/filepath"
	"strings"
)

const (
	AttributeTypeStringGoType   = "string"
	AttributeTypeIntGoType      = "int32"
	AttributeTypeLongGoType     = "int64"
	AttributeTypeBoolGoType     = "bool"
	AttributeTypeObjectIdGoType = "primitive.ObjectID"
	AttributeTypeDateGoType     = "primitive.DateTime"
)

type InfoCollectorVisitor struct {
	Attributes []CodeGenAttribute
}

type CodeGenAttribute interface {
	GetDefinition() schema.Field
	GetBsonPropertyName(qualified bool) string
	GetName(qualified bool, prefixed bool) string
	GetPrefix(casing string) string
	GetNumberOfDescendents() int
	GetPaths(pathType string) []string

	GetGoPackageImports() []string
	GetGoAttributeName() string
	GetGoAttributeType() string
	GetGoAttributeTag() string
	GetGoAttributeIsZeroCondition() string
	IsArrayItem() bool
}

type CodeGenCollection struct {
	Schema *schema.Collection

	/* List only top level attributes in root structure */
	DirectAttributes []CodeGenAttribute

	/* List all the attributes in the tree */
	totalNumberOfAttributes int
	Attributes              []CodeGenAttribute
}

type CodeGenAttributeImpl struct {
	ParentAttribute     CodeGenAttribute
	AttrDefinition      schema.Field
	NumberOfDescendents int
	AttributeNamespace  string
	BSONNamespace       string
	ArrayItem           bool
	Tags                []schema.Tag
	Prefix              string
	PackageImports      []string
}

type ValueTypeAttribute struct {
	CodeGenAttributeImpl
}

type RefStructAttribute struct {
	CodeGenAttributeImpl
	ReferredStruct CodeGenAttribute
}

type StructAttribute struct {
	CodeGenAttributeImpl
	Attributes []CodeGenAttribute
}

type ArrayAttribute struct {
	CodeGenAttributeImpl
	Item CodeGenAttribute
}

type MapAttribute struct {
	CodeGenAttributeImpl
	Item CodeGenAttribute
}

func NewCodeGenCollection(aSchema *schema.Collection) (*CodeGenCollection, error) {

	c := &CodeGenCollection{aSchema, nil, 0, nil}

	// va := NamespaceVisitor{}
	numberOfDescendents := 0
	for _, f := range aSchema.Attributes {
		a := NewAttribute(nil, f, false, aSchema.Properties.Prefix /*, va */)
		c.DirectAttributes = append(c.DirectAttributes, a)
		numberOfDescendents += 1 + a.GetNumberOfDescendents()
	}

	c.totalNumberOfAttributes = numberOfDescendents
	c.Attributes = c.findAttributes()
	// fmt.Printf("number of descendents is: %d\n", numberOfDescendents)
	return c, nil
}

func NewAttribute(parentAttribute CodeGenAttribute, attrDefinition schema.Field, isArrayItem bool, prefix string /* , visitor NamespaceVisitor */) CodeGenAttribute {

	var a CodeGenAttribute
	switch attrDefinition.Typ {
	case schema.AttributeTypeRefStruct:
		a = NewRefStructTypeAttribute(parentAttribute, attrDefinition, isArrayItem, prefix /*, visitor */)
	case schema.AttributeTypeStruct:
		a = NewStructAttribute(parentAttribute, attrDefinition, isArrayItem, prefix /*, visitor */)
	case schema.AttributeTypeArray:
		a = NewArrayAttribute(parentAttribute, attrDefinition, isArrayItem, prefix /*, visitor */)
	case schema.AttributeTypeMap:
		a = NewMapAttribute(parentAttribute, attrDefinition, isArrayItem, prefix /*, visitor */)
	default:
		a = NewValueTypeAttribute(parentAttribute, attrDefinition, isArrayItem, prefix /*, visitor */)
	}

	return a
}

func NewValueTypeAttribute(parentAttribute CodeGenAttribute, attrDefinition schema.Field, isArrayItem bool, prefix string /*, visitor NamespaceVisitor*/) CodeGenAttribute {
	v := &ValueTypeAttribute{}
	v.AttrDefinition = attrDefinition
	v.ArrayItem = isArrayItem
	v.ParentAttribute = parentAttribute
	v.Prefix = prefix
	v.Tags = attrDefinition.GetTagsAsListOfTag(true, true, []string{"json", "bson"})

	switch strings.ToLower(attrDefinition.Typ) {
	case schema.AttributeTypeDate:
		v.PackageImports = make([]string, 0, 1)
		v.PackageImports = append(v.PackageImports, "go.mongodb.org/mongo-driver/bson/primitive")

	case schema.AttributeTypeObjectId:
		v.PackageImports = make([]string, 0, 1)
		v.PackageImports = append(v.PackageImports, "go.mongodb.org/mongo-driver/bson/primitive")
	}

	// v.BSONNamespace = visitor.BSONNamespace
	// v.AttributeNamespace = visitor.AttributeNamespace

	if parentAttribute != nil {
		v.BSONNamespace = parentAttribute.GetBsonPropertyName(true)
		v.AttributeNamespace = parentAttribute.GetName(true, false)
	}
	return v
}

func NewRefStructTypeAttribute(parentAttribute CodeGenAttribute, attrDefinition schema.Field, isArrayItem bool, prefix string /*, visitor NamespaceVisitor*/) CodeGenAttribute {
	v := &RefStructAttribute{}
	v.AttrDefinition = attrDefinition
	v.ArrayItem = isArrayItem
	v.ParentAttribute = parentAttribute
	v.Prefix = prefix
	v.Tags = attrDefinition.GetTagsAsListOfTag(true, true, []string{"json", "bson"})

	switch strings.ToLower(attrDefinition.Typ) {
	case schema.AttributeTypeRefStruct:
		if attrDefinition.StructRef.Package != "" {
			v.PackageImports = make([]string, 0, 1)
			v.PackageImports = append(v.PackageImports, attrDefinition.StructRef.Package)
		}
		/* Introduced the StructReference support info.
		if slashPos := strings.LastIndex(attrDefinition.StructName, "/"); slashPos > 0 {
			v.PackageImports = make([]string, 0, 1)
			v.PackageImports = append(v.PackageImports, attrDefinition.StructName[0:slashPos])
		}
		*/

	}

	// v.BSONNamespace = visitor.BSONNamespace
	// v.AttributeNamespace = visitor.AttributeNamespace

	if parentAttribute != nil {
		v.BSONNamespace = parentAttribute.GetBsonPropertyName(true)
		v.AttributeNamespace = parentAttribute.GetName(true, false)
	}
	return v
}

func NewStructAttribute(parentAttribute CodeGenAttribute, attrDefinition schema.Field, isArrayItem bool, prefix string /*, visitor NamespaceVisitor */) CodeGenAttribute {
	s := &StructAttribute{}
	s.AttrDefinition = attrDefinition
	s.ArrayItem = isArrayItem
	s.ParentAttribute = parentAttribute
	s.Prefix = prefix
	s.Tags = attrDefinition.GetTagsAsListOfTag(true, true, []string{"json", "bson"})

	/*
		s.BSONNamespace = visitor.BSONNamespace
		s.AttributeNamespace = visitor.AttributeNamespace
	*/

	if parentAttribute != nil {
		s.BSONNamespace = parentAttribute.GetBsonPropertyName(true)
		s.AttributeNamespace = parentAttribute.GetName(true, false)
	}

	/*
		bsonPropertyName := s.AttrDefinition.GetTagNameValue("bson")
		visitor.BSONNamespace = util.AppendToNamespace(visitor.BSONNamespace, bsonPropertyName, ".")
		visitor.AttributeNamespace = util.AppendToNamespace(visitor.AttributeNamespace, s.AttrDefinition.Name, ".")
	*/
	s.NumberOfDescendents = len(attrDefinition.Attributes)
	for _, f := range attrDefinition.Attributes {
		a := NewAttribute(s, f, false, prefix /*, visitor */)
		s.Attributes = append(s.Attributes, a)

		s.NumberOfDescendents += a.GetNumberOfDescendents()
	}

	return s
}

func NewArrayAttribute(parentAttribute CodeGenAttribute, attrDefinition schema.Field, isArrayItem bool, prefix string /*, visitor NamespaceVisitor */) CodeGenAttribute {
	a := &ArrayAttribute{}
	a.AttrDefinition = attrDefinition
	a.ArrayItem = isArrayItem
	a.ParentAttribute = parentAttribute
	a.Prefix = prefix
	a.Tags = attrDefinition.GetTagsAsListOfTag(true, true, []string{"json", "bson"})

	/*
		a.BSONNamespace = visitor.BSONNamespace
		a.AttributeNamespace = visitor.AttributeNamespace
	*/

	if parentAttribute != nil {
		a.BSONNamespace = parentAttribute.GetBsonPropertyName(true)
		a.AttributeNamespace = parentAttribute.GetName(true, false)
	}

	/*
			bsonPropertyName := a.AttrDefinition.GetTagNameValue("bson")
		visitor.BSONNamespace = util.AppendToNamespace(visitor.BSONNamespace, bsonPropertyName, ".")
		visitor.AttributeNamespace = util.AppendToNamespace(visitor.AttributeNamespace, a.AttrDefinition.Name, ".")
	*/
	a.Item = NewAttribute(a, *attrDefinition.Item, true, prefix /*, visitor */)
	a.NumberOfDescendents = a.Item.GetNumberOfDescendents()

	return a
}

func NewMapAttribute(parentAttribute CodeGenAttribute, attrDefinition schema.Field, isArrayItem bool, prefix string /*, visitor NamespaceVisitor */) CodeGenAttribute {
	a := &MapAttribute{}
	a.AttrDefinition = attrDefinition
	a.ArrayItem = isArrayItem
	a.ParentAttribute = parentAttribute
	a.Prefix = prefix
	a.Tags = attrDefinition.GetTagsAsListOfTag(true, true, []string{"json", "bson"})

	/*
		a.BSONNamespace = visitor.BSONNamespace
		a.AttributeNamespace = visitor.AttributeNamespace
	*/

	if parentAttribute != nil {
		a.BSONNamespace = parentAttribute.GetBsonPropertyName(true)
		a.AttributeNamespace = parentAttribute.GetName(true, false)
	}

	/*
			bsonPropertyName := a.AttrDefinition.GetTagNameValue("bson")
		visitor.BSONNamespace = util.AppendToNamespace(visitor.BSONNamespace, bsonPropertyName, ".")
		visitor.AttributeNamespace = util.AppendToNamespace(visitor.AttributeNamespace, a.AttrDefinition.Name, ".")
	*/
	a.Item = NewAttribute(a, *attrDefinition.Item, false, prefix /*, visitor */)
	a.NumberOfDescendents = a.Item.GetNumberOfDescendents()

	return a
}

/*
 * Methods of CodeGenCollection
 */
func (c *CodeGenCollection) findAttributes() []CodeGenAttribute {
	infoCollector := InfoCollectorVisitor{Attributes: make([]CodeGenAttribute, 0, c.totalNumberOfAttributes)}
	for _, a := range c.DirectAttributes {
		infoCollector.Attributes = append(infoCollector.Attributes, a)
		switch a1 := a.(type) {
		case *StructAttribute:
			infoCollector = a1.FindAttributes(infoCollector)
		case *ArrayAttribute:
			if sItem, ok := a1.Item.(*StructAttribute); ok {
				infoCollector = sItem.FindAttributes(infoCollector)
			}
		case *MapAttribute:
			if sItem, ok := a1.Item.(*StructAttribute); ok {
				infoCollector = sItem.FindAttributes(infoCollector)
			}
		}
	}

	return infoCollector.Attributes
}

/*
func (c *CodeGenCollection) findAttributeByStructName(s string) CodeGenAttribute {

	for _, a := range c.Attributes {
		if a.GetDefinition().Typ == schema.AttributeTypeStruct && a.GetDefinition().StructName == s {
				return a
		}
	}

	return nil
}
*/

func (c *CodeGenCollection) GetStructAttributes() []CodeGenAttribute {
	attrs := make([]CodeGenAttribute, 0, 10)
	for _, a := range c.Attributes {
		if a.GetDefinition().Typ == schema.AttributeTypeStruct {
			attrs = append(attrs, a)
		} else if a.GetDefinition().Typ == schema.AttributeTypeArray {
			i := a.(*ArrayAttribute).Item
			if i.GetDefinition().Typ == schema.AttributeTypeStruct {
				attrs = append(attrs, i)
			}
		} else if a.GetDefinition().Typ == schema.AttributeTypeMap {
			i := a.(*MapAttribute).Item
			if i.GetDefinition().Typ == schema.AttributeTypeStruct {
				attrs = append(attrs, i)
			}
		}
	}

	return attrs
}

/*
func (c *CodeGenCollection) GetStructRefAttributes() []CodeGenAttribute {
	attrs := make([]CodeGenAttribute, 0, 10)
	for _, a := range c.Attributes {
		if a.GetDefinition().Typ == schema.AttributeTypeRefStruct {
			referredAttr := c.findAttributeByStructName(a.GetDefinition().StructName)
			if referredAttr != nil {
				a.(*RefStructAttribute).ReferredStruct = referredAttr
				attrs = append(attrs, a)
			}
		}
	}

	return attrs
}
*/

func (c *CodeGenCollection) GetPackageImports(limitToQueriable bool) []string {

	pkgs := make([]string, 0)
	for _, a := range c.Attributes {
		if a.GetDefinition().Queryable || !limitToQueriable {
			if ap := a.GetGoPackageImports(); len(ap) > 0 {
				pkgs = append(pkgs, ap...)
			}
		}
	}

	if len(pkgs) > 1 {
		pkgs = util.RemoveDuplicates(pkgs)
	}

	return pkgs
}

func (c *CodeGenCollection) GetGeneratedContentPath(mountDir string) (string, error) {

	subFolder := c.Schema.Properties.FolderPath
	if subFolder == "" {
		if c.Schema.Properties.PackageName != "" {
			if util.IsDottedForm(c.Schema.Properties.PackageName) {
				segs := strings.Split(c.Schema.Properties.PackageName, ".")
				subFolder = filepath.Join(segs...)
			} else {
				subFolder = c.Schema.Properties.PackageName
			}
		}
	}

	contentPath := filepath.Join(mountDir, subFolder)
	if _, err := os.Stat(contentPath); os.IsNotExist(err) {
		if err = os.MkdirAll(contentPath, 0755); err != nil {
			return "", err
		}
	}

	return contentPath, nil
}

func (c *CodeGenCollection) GetPackageName() string {

	s := c.Schema.Properties.FolderPath
	if s == "" {
		s = c.Schema.Properties.PackageName
	}

	as := util.SliceSegmentedName(s)
	return as[len(as)-1]
}

func (c *CodeGenCollection) GetPrefix(casing string) string {

	s := c.Schema.Properties.Prefix
	if s != "" {
		switch strings.ToLower(casing) {
		case "lower":
			s = strings.ToLower(s)
		case "capital":
			s = strings.Title(s)
		default:
			s = strings.ToUpper(s)
		}
	}

	return s
}

func (c *CodeGenCollection) GetMorphiaPackage() string {
	return c.Schema.Properties.MorphiaPackage
}

/*
 * CodeGenAttributeImpl
 */
func (b *CodeGenAttributeImpl) HasOption(o string) bool {
	return b.AttrDefinition.Options != "" && strings.Contains(b.AttrDefinition.Options, o)
}

func (b *CodeGenAttributeImpl) GetGoAttributeIsZeroCondition() string {

	s := ""
	switch b.AttrDefinition.Typ {
	case schema.AttributeTypeObjectId:
		s = fmt.Sprintf("s.%s != primitive.NilObjectID", b.GetGoAttributeName())
	case schema.AttributeTypeInt:
		s = fmt.Sprintf("s.%s != 0", b.GetGoAttributeName())
	case schema.AttributeTypeLong:
		s = fmt.Sprintf("s.%s != 0", b.GetGoAttributeName())
	case schema.AttributeTypeBool:
		s = fmt.Sprintf("s.%s", b.GetGoAttributeName())
	case schema.AttributeTypeDate:
		s = fmt.Sprintf("s.%s != 0", b.GetGoAttributeName())
	case schema.AttributeTypeArray:
		s = fmt.Sprintf("len(s.%s) != 0", b.GetGoAttributeName())
	case schema.AttributeTypeStruct:
		s = fmt.Sprintf("!s.%s.IsZero()", b.GetGoAttributeName())
	case schema.AttributeTypeRefStruct:
		s = fmt.Sprintf("!s.%s.IsZero()", b.GetGoAttributeName())
	case schema.AttributeTypeString:
		s = fmt.Sprintf("s.%s != \"\"", b.GetGoAttributeName())
	case schema.AttributeTypeMap:
		s = fmt.Sprintf("len(s.%s) != 0", b.GetGoAttributeName())
	}

	return s
}

func (b *CodeGenAttributeImpl) GetPrefix(casing string) string {

	s := ""
	if b.Prefix != "" {
		switch strings.ToLower(casing) {
		case "lower":
			s = strings.ToLower(b.Prefix)
		case "capital":
			s = strings.Title(b.Prefix)
		default:
			s = strings.ToUpper(b.Prefix)
		}
	}

	return s
}

func (b *CodeGenAttributeImpl) GetGoPackageImports() []string {
	return b.PackageImports
}

func (b *CodeGenAttributeImpl) GetParent() CodeGenAttribute {
	return b.ParentAttribute
}

func (b *CodeGenAttributeImpl) GetDefinition() schema.Field {
	return b.AttrDefinition
}

func (b *CodeGenAttributeImpl) IsArrayItem() bool {
	return b.ArrayItem
}

func (b *CodeGenAttributeImpl) GetNumberOfDescendents() int {
	return b.NumberOfDescendents
}

func (b *CodeGenAttributeImpl) GetBsonPropertyName(qualified bool) string {
	s := b.AttrDefinition.GetTagNameValue("bson")
	if qualified {
		return util.AppendToNamespace(b.BSONNamespace, s, ".")
	}

	return s
}

/*
func (b *CodeGenAttributeImpl) GetBsonPropertyQualifiedName() string {
	s := b.AttrDefinition.GetTagNameValue("bson")
	return util.AppendToNamespace( b.BSONNamespace, s , ".")
}
*/
func (b *CodeGenAttributeImpl) GetPaths(pathType string) []string {

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

func (b *CodeGenAttributeImpl) GetName(qualified bool, prefixed bool) string {
	s := b.AttrDefinition.Name
	if qualified {
		s = util.AppendToNamespace(b.AttributeNamespace, s, ".")

		if prefixed && b.Prefix != "" {
			s = strings.Join([]string{b.Prefix, s}, "_")
		}

	}

	return s
}

func (b *CodeGenAttributeImpl) String() string {
	return fmt.Sprintf("%s of type: %s and ns: %s", b.AttrDefinition.Name, b.AttrDefinition.Typ, b.AttributeNamespace)
}

func (b *CodeGenAttributeImpl) GetGoAttributeType() string {
	panic(errors.New("CodeGenAttributeImpl does not implement GetGoAttributeType"))
}

func (b *CodeGenAttributeImpl) GetGoAttributeName() string {
	return util.ToCapitalCase(b.AttrDefinition.Name)
}

func (b *CodeGenAttributeImpl) GetGoAttributeTag() string {

	if len(b.Tags) > 0 {
		var stb strings.Builder
		stb.WriteRune('`')
		for i, t := range b.Tags {
			if i > 0 {
				stb.WriteRune(' ')
			}
			stb.WriteString(t.String())
		}
		/*
			for i := 0; i < len(b.AttrDefinition.Tags); i+=2 {
				if i > 0 {
					stb.WriteRune(' ')
				}
				stb.WriteString(b.AttrDefinition.Tags[i])
				stb.WriteString(":\"")
				stb.WriteString(b.AttrDefinition.Tags[i+1])
				stb.WriteString("\"")
			}
		*/
		stb.WriteRune('`')
		return stb.String()
	}

	return ""
}

/*
 * ValueType implmentation.
 */
func (v *ValueTypeAttribute) GetGoAttributeType() string {

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
	default:
		t = "-- NotYetImplemented --"
	}
	return t
}

/*
 * RefStruct Implementation
 */
func (v *RefStructAttribute) GetGoAttributeType() string {
	if v.AttrDefinition.StructRef.Package != "" {
		if pkgNdx := strings.LastIndex(v.AttrDefinition.StructRef.Package, "/"); pkgNdx >= 0 {
			return strings.Join([]string{v.AttrDefinition.StructRef.Package[pkgNdx+1:], v.AttrDefinition.StructRef.StructName}, ".")
		}
	}
	return v.AttrDefinition.StructRef.StructName
}

/*
 * Struct implementation.
 */
func (s *StructAttribute) GetGoAttributeType() string {
	return s.AttrDefinition.StructName
}

func (s *StructAttribute) FindAttributes(infoCollector InfoCollectorVisitor) InfoCollectorVisitor {

	for _, a := range s.Attributes {
		infoCollector.Attributes = append(infoCollector.Attributes, a)
		switch a1 := a.(type) {
		case *StructAttribute:
			infoCollector = a1.FindAttributes(infoCollector)
		case *ArrayAttribute:
			if sItem, ok := a1.Item.(*StructAttribute); ok {
				infoCollector = sItem.FindAttributes(infoCollector)
			}
		case *MapAttribute:
			if sItem, ok := a1.Item.(*StructAttribute); ok {
				infoCollector = sItem.FindAttributes(infoCollector)
			}
		}

	}

	return infoCollector
}

/*
 * Array implementation.
 */
func (a *ArrayAttribute) GetGoAttributeType() string {
	return "[]" + a.Item.GetGoAttributeType()
}

/*
 * Map implementation.
 */
func (a *MapAttribute) GetGoAttributeType() string {
	return "map[string]" + a.Item.GetGoAttributeType()
}
