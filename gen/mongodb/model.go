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
	AttributeTypeStringGoType = "string"
	AttributeTypeIntGoType    = "int32"
)

type InfoCollectorVisitor struct {
	Attributes []CodeGenAttribute
}

type CodeGenAttribute interface {
	GetDefinition() schema.Field
	GetBsonPropertyName(qualified bool) string
	GetName(qualified bool) string

	GetNumberOfDescendents() int

	GetGoAttributeName() string
	GetGoAttributeType() string
	GetGoAttributeTag() string
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
}

type ValueTypeAttribute struct {
	CodeGenAttributeImpl
}

type StructAttribute struct {
	CodeGenAttributeImpl
	Attributes []CodeGenAttribute
}

type ArrayAttribute struct {
	CodeGenAttributeImpl
	Item CodeGenAttribute
}

func NewCodeGenCollection(aSchema *schema.Collection) (*CodeGenCollection, error) {

	c := &CodeGenCollection{aSchema, nil, 0, nil}

	// va := NamespaceVisitor{}
	numberOfDescendents := 0
	for _, f := range aSchema.Attributes {
		a := NewAttribute(nil, f, false /*, va */)
		c.DirectAttributes = append(c.DirectAttributes, a)
		numberOfDescendents += 1 + a.GetNumberOfDescendents()
	}

	c.totalNumberOfAttributes = numberOfDescendents
	c.Attributes = c.FindAttributes()
	fmt.Printf("number of descendents is: %d\n", numberOfDescendents)
	return c, nil
}

func NewAttribute(parentAttribute CodeGenAttribute, attrDefinition schema.Field, isArrayItem bool /* , visitor NamespaceVisitor */) CodeGenAttribute {

	var a CodeGenAttribute
	switch attrDefinition.Typ {
	case schema.AttributeTypeStruct:
		a = NewStructAttribute(parentAttribute, attrDefinition, isArrayItem /*, visitor */)
	case schema.AttributeTypeArray:
		a = NewArrayAttribute(parentAttribute, attrDefinition, isArrayItem /*, visitor */)
	default:
		a = NewValueTypeAttribute(parentAttribute, attrDefinition, isArrayItem /*, visitor */)
	}

	return a
}

func NewValueTypeAttribute(parentAttribute CodeGenAttribute, attrDefinition schema.Field, isArrayItem bool /*, visitor NamespaceVisitor*/) CodeGenAttribute {
	v := &ValueTypeAttribute{}
	v.AttrDefinition = attrDefinition
	v.ArrayItem = isArrayItem
	v.ParentAttribute = parentAttribute
	v.Tags = attrDefinition.GetTagsAsListOfTag(true, true, []string{"json", "bson"})

	// v.BSONNamespace = visitor.BSONNamespace
	// v.AttributeNamespace = visitor.AttributeNamespace

	if parentAttribute != nil {
		v.BSONNamespace = parentAttribute.GetBsonPropertyName(true)
		v.AttributeNamespace = parentAttribute.GetName(true)
	}
	return v
}

func NewStructAttribute(parentAttribute CodeGenAttribute, attrDefinition schema.Field, isArrayItem bool /*, visitor NamespaceVisitor */) CodeGenAttribute {
	s := &StructAttribute{}
	s.AttrDefinition = attrDefinition
	s.ArrayItem = isArrayItem
	s.ParentAttribute = parentAttribute
	s.Tags = attrDefinition.GetTagsAsListOfTag(true, true, []string{"json", "bson"})

	/*
		s.BSONNamespace = visitor.BSONNamespace
		s.AttributeNamespace = visitor.AttributeNamespace
	*/

	if parentAttribute != nil {
		s.BSONNamespace = parentAttribute.GetBsonPropertyName(true)
		s.AttributeNamespace = parentAttribute.GetName(true)
	}

	/*
		bsonPropertyName := s.AttrDefinition.GetTagNameValue("bson")
		visitor.BSONNamespace = util.AppendToNamespace(visitor.BSONNamespace, bsonPropertyName, ".")
		visitor.AttributeNamespace = util.AppendToNamespace(visitor.AttributeNamespace, s.AttrDefinition.Name, ".")
	*/
	s.NumberOfDescendents = len(attrDefinition.Attributes)
	for _, f := range attrDefinition.Attributes {
		a := NewAttribute(s, f, false /*, visitor */)
		s.Attributes = append(s.Attributes, a)

		s.NumberOfDescendents += a.GetNumberOfDescendents()
	}

	return s
}

func NewArrayAttribute(parentAttribute CodeGenAttribute, attrDefinition schema.Field, isArrayItem bool /*, visitor NamespaceVisitor */) CodeGenAttribute {
	a := &ArrayAttribute{}
	a.AttrDefinition = attrDefinition
	a.ArrayItem = isArrayItem
	a.ParentAttribute = parentAttribute
	a.Tags = attrDefinition.GetTagsAsListOfTag(true, true, []string{"json", "bson"})

	/*
		a.BSONNamespace = visitor.BSONNamespace
		a.AttributeNamespace = visitor.AttributeNamespace
	*/

	if parentAttribute != nil {
		a.BSONNamespace = parentAttribute.GetBsonPropertyName(true)
		a.AttributeNamespace = parentAttribute.GetName(true)
	}

	/*
			bsonPropertyName := a.AttrDefinition.GetTagNameValue("bson")
		visitor.BSONNamespace = util.AppendToNamespace(visitor.BSONNamespace, bsonPropertyName, ".")
		visitor.AttributeNamespace = util.AppendToNamespace(visitor.AttributeNamespace, a.AttrDefinition.Name, ".")
	*/
	a.Item = NewAttribute(a, *attrDefinition.Item, true /*, visitor */)
	a.NumberOfDescendents = a.Item.GetNumberOfDescendents()

	return a
}

func (c *CodeGenCollection) FindAttributes() []CodeGenAttribute {
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
		}
	}

	return infoCollector.Attributes
}

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
		}
	}

	return attrs
}

/*
 * Methods of CodeGenCollection
 */
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

/*
 * CodeGenAttributeImpl
 */
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

func (b *CodeGenAttributeImpl) GetName(qualified bool) string {
	s := b.AttrDefinition.Name
	if qualified {
		return util.AppendToNamespace(b.AttributeNamespace, b.AttrDefinition.Name, ".")
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
	default:
		t = "-- NotYetImplemented --"
	}
	return t
}

/*
 * Struct implmentation.
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
		}
	}

	return infoCollector
}

/*
 * Array implmentation.
 */
func (a *ArrayAttribute) GetGoAttributeType() string {
	return "[]" + a.Item.GetGoAttributeType()
}
