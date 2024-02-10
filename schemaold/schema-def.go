package schemaold

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system/util"
	"github.com/rs/zerolog/log"
	"io"
	"strings"
)

const (
	AttributeTypeArray     = "array"
	AttributeTypeMap       = "map"
	AttributeTypeStruct    = "struct"
	AttributeTypeString    = "string"
	AttributeTypeInt       = "int"
	AttributeTypeLong      = "long"
	AttributeTypeBool      = "bool"
	AttributeTypeDate      = "date"
	AttributeTypeObjectId  = "object-id"
	AttributeTypeRefStruct = "ref-struct"
	AttributeTypeDocument  = "document"
)

type CollectionDefError struct {
	ctx string
	msg string
}

func (ce *CollectionDefError) Error() string {
	return fmt.Sprintf("[%s] - %s", ce.ctx, ce.msg)
}

type CollectionProps struct {
	FolderPath     string `json:"folder-path,omitempty"`
	Prefix         string `json:"prefix,omitempty"`
	PackageName    string `json:"package-name,omitempty"`
	StructName     string `json:"struct-name,omitempty"`
	MorphiaPackage string `json:"morphia-pkg,omitempty"`
}

type Collection struct {
	Name       string          `json:"name,omitempty"`
	Properties CollectionProps `json:"properties,omitempty"`
	Attributes []Field         `json:"attributes,omitempty"`

	AllAttributes []*Field `json:"-"`
}

type StructReference struct {
	StructName string `json:"struct-name,omitempty"`
	IsExternal bool   `json:"is-external,omitempty"`
	Package    string
	Item       *Field
}

type Field struct {
	Name       string          `json:"name,omitempty"`
	StructName string          `json:"struct-name,omitempty"`
	Typ        string          `json:"type,omitempty"`
	IsKey      bool            `json:"is-key,omitempty"`
	Tags       []string        `json:"tags,omitempty"`
	Attributes []Field         `json:"attributes,omitempty"`
	Item       *Field          `json:"item,omitempty"`
	Queryable  bool            `json:"queryable,omitempty"`
	StructRef  StructReference `json:"struct-ref,omitempty"`
	Paths      []string        `json:"-"`
	BsonPaths  []string        `json:"-"`
	Options    string          `json:"options,omitempty"`
}

type InfoCollectorVisitor struct {
	Attributes []*Field
}

/*type VisitPhase int
const (
  StartVisit VisitPhase = iota
  DoVisit
  EndVisit
)

type VisitorState interface {

}

*/

type Visitor interface {
	startVisit(f *Field)
	visit(f *Field)
	endVisit(f *Field)
}

/*
 * Tag Methods
 */
type Tag struct {
	Name    string
	Value   string
	Options []string
}

func (t *Tag) String() string {

	if t.Value == "" && len(t.Options) == 0 {
		return ""
	}

	var stb strings.Builder
	stb.WriteString(t.Name)

	stb.WriteString(":\"")
	if t.Value != "" {
		stb.WriteString(t.Value)
	}

	if len(t.Options) > 0 {
		stb.WriteRune(',')
		stb.WriteString(strings.Join(t.Options, ","))
	}
	stb.WriteString("\"")

	return stb.String()
}

func NewTag(tn string, s string) Tag {

	t := Tag{Name: tn}
	if s != "" {
		sa := strings.Split(s, ",")
		if len(sa) == 0 {
			panic("NewTag: tag " + s + " cannot be split properly")
		}

		t.Value = sa[0]
		if len(sa) > 1 {
			t.Options = sa[1:]
		}
	}

	return t
}

/*
 * Collection Methods
 */
func (c *Collection) ToJsonString(prefix string, indent string) string {

	if s, e := json.MarshalIndent(c, prefix, indent); e != nil {
		return e.Error()
	} else {
		return string(s)
	}
}

func (c *Collection) findAttributes() []*Field {
	infoCollector := InfoCollectorVisitor{Attributes: make([]*Field, 0)}
	for i := range c.Attributes {
		infoCollector.Attributes = append(infoCollector.Attributes, &c.Attributes[i])
		infoCollector = c.Attributes[i].FindAttributes(infoCollector)
	}

	return infoCollector.Attributes
}

func (c *Collection) visit(v Visitor) {
	for i := range c.Attributes {
		v.visit(&c.Attributes[i])
		c.Attributes[i].visit(v)
	}
}

func (c *Collection) wireReference2Structs(fields []*Field) error {
	refs := make(map[string]*Field)
	for _, f := range fields {
		if f.Typ == AttributeTypeStruct {
			refs[f.StructName] = f
		} else if (f.Typ == AttributeTypeMap || f.Typ == AttributeTypeArray) && f.Item.Typ == AttributeTypeStruct {
			refs[f.Item.StructName] = f.Item
		}
	}

	for _, f := range fields {
		if f.Typ == AttributeTypeRefStruct {
			if item, ok := refs[f.StructRef.StructName]; !ok {
				// Unresolved reference. Can be declared externally to current file.
				// Another approach could be to process two schemas as a bundle but is way to elaborate.
				// Should probably add a dependency in terms of a package import
				if !f.StructRef.IsExternal {
					// TODO: the f.StructName in the message should be replaced with f.StructRef.StructName
					return errors.New(fmt.Sprintf("the field %s refers to undefined struct %s", f.Name, f.StructName))
				}
			} else {
				f.StructRef.Item = item
			}
		}
	}

	return nil
}

/*
 * Field Methods
 */
func (f Field) IsStructType() bool {
	return f.Typ == AttributeTypeStruct || f.Typ == AttributeTypeRefStruct
}

func (f Field) IsCollectionType() bool {
	return f.Typ == AttributeTypeMap || f.Typ == AttributeTypeArray
}

func (f Field) IsValueType() bool {
	return !f.IsStructType() && !f.IsCollectionType()
}

func (f *Field) GetTagsAsListOfTag(forceFieldNameIfMissing bool, forceOmitIfEmpty bool, tagsOfInterest []string) []Tag {
	var tags []Tag
	if len(f.Tags) > 0 {
		tags = make([]Tag, 0, len(f.Tags)/2)
		for i := 0; i < len(f.Tags); i += 2 {
			nt := NewTag(f.Tags[i], f.Tags[i+1])
			if nt.Value == "" && forceFieldNameIfMissing {
				nt.Value = f.Name
			}

			if strings.HasPrefix(f.Tags[i+1], "-") {
				if strings.Contains(f.Tags[i+1], ",omitempty") {
					f.Tags[i+1] = strings.ReplaceAll(f.Tags[i+1], ",omitempty", "")
				}
			} else {
				if !strings.Contains(f.Tags[i+1], "omitempty") && (forceOmitIfEmpty || strings.HasPrefix(f.Tags[i+1], "_id")) {
					nt.Options = append(nt.Options, "omitempty")
				}
			}

			tags = append(tags, nt)
		}
	}

	if len(tagsOfInterest) > 0 && (forceOmitIfEmpty || forceFieldNameIfMissing) {
		for _, s := range tagsOfInterest {
			found := false
			if len(tags) > 0 {
				for _, t := range tags {
					if t.Name == s {
						found = true
					}
				}
			}

			if !found {
				t := Tag{Name: s}
				if forceFieldNameIfMissing {
					t.Value = f.Name
				}
				if forceOmitIfEmpty {
					t.Options = []string{"omitempty"}
				}

				tags = append(tags, t)
			}
		}
	}

	return tags
}

func (f *Field) GetTag(tagName string) Tag {

	if len(f.Tags) > 0 {
		for ndx, t := range f.Tags {
			if strings.EqualFold(t, tagName) {
				return NewTag(tagName, f.Tags[ndx+1])
			}
		}
	}

	return Tag{Name: tagName}
}

func (f *Field) GetTagNameValue(tagName string) string {

	if len(f.Tags) > 0 {
		for ndx, t := range f.Tags {
			if strings.EqualFold(t, tagName) {
				tv := f.Tags[ndx+1]
				arr := strings.Split(tv, ",")
				return arr[0]
			}
		}
	}

	return f.Name
}

/*
 *
 */
func (f *Field) FindAttributes(infoCollector InfoCollectorVisitor) InfoCollectorVisitor {

	for i := range f.Attributes {
		infoCollector.Attributes = append(infoCollector.Attributes, &f.Attributes[i])
		infoCollector = f.Attributes[i].FindAttributes(infoCollector)
	}

	if f.Item != nil {
		infoCollector = f.Item.FindAttributes(infoCollector)
	}

	return infoCollector
}

func (f *Field) visit(v Visitor) {
	v.startVisit(f)

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

	v.endVisit(f)

}

func (f Field) String() string {
	return fmt.Sprintf("%s - %s - %s", f.Name, f.Typ, f.StructName)
}

/*
 * Read and Validate Function
 */
func ReadCollectionDefinition(reader io.Reader) (*Collection, error) {

	log.Debug().Msg("start reading collection definition")

	schema := &Collection{}

	/*  Alternate: not able to force errors on unknown fields
	if e := json.Unmarshal([]byte(def), schema); e != nil {
		return nil, &CollectionDefError{msg: e.Error()}
	}
	*/
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&schema); err != nil {
		return nil, &CollectionDefError{msg: err.Error()}
	}

	/* Since I just deserialized stuff, just try to check that pretty much che stuff is correct or at least
	 * not massively wrong.
	 */
	if e := validateCollectionDef(schema); e != nil {
		return nil, e
	}

	/* Now traverse the tree and get the list of attrs. yes, it could have been done in the validation phase.
	 * May be later on will get into that processing.
	 */
	fields := schema.findAttributes()
	schema.AllAttributes = fields

	/*
	 * Now it's time to wire the references to structs in order to calculate all the paths or heirarchies to the leaves.
	 * In principle the wiring should detect loops in the config. Loops at the moment might cause a stack issue, but nevertheless
	 * should be handled as pointer to struct in the generation.
	 */
	if err := schema.wireReference2Structs(fields); err != nil {
		return nil, err
	}

	lf := PathFinderVisitor{}
	schema.visit(&lf)

	return schema, nil
}

func validateCollectionDef(c *Collection) error {

	log.Debug().Str("name", c.Name).Msg("start validating collection definition")

	/*
	 * Pre-validation of collection name to assign the root path..
	 */
	pPath := strings.TrimSpace(c.Name)
	if len(pPath) == 0 {
		pPath = "<undefined>"
	}

	if !util.NameWellFormed(c.Name) {
		ce := &CollectionDefError{ctx: pPath, msg: "collection name is not well formed"}
		return ce
	}

	if e := validateProperties(c); e != nil {
		return e
	}

	return validateAttributes(c.Attributes, pPath, nil)
}

func validateProperties(c *Collection) error {

	cp := &c.Properties
	// Doing a XOR...
	if (cp.PackageName != "") == (cp.FolderPath != "") {
		return &CollectionDefError{msg: fmt.Sprintf("one (and only one) of packageName (%s) and folderPath (%s) have to be specified", cp.PackageName, cp.FolderPath)}
	}

	if cp.StructName == "" {
		cp.StructName = util.ToCapitalCase(c.Name)
		log.Debug().Msg("no name provided for struct name...using " + cp.StructName)
	}

	return nil
}

func validateAttributes(props []Field, pPath string, parentField *Field) error {

	log.Debug().Str("path", pPath).Msg("start validating attributes")

	if len(props) == 0 {
		return &CollectionDefError{ctx: pPath, msg: "attributes missing from schema"}
	}

	for i := range props {
		// Use the indexing property to address the underlay-ing object and not a copy.
		if e := validateField(&props[i], pPath, parentField); e != nil {
			return e
		}
	}

	return nil
}

func validateField(f *Field, pPath string, parentField *Field) error {

	arrayItemDefinition := false

	if parentField != nil && (parentField.Typ == AttributeTypeArray || parentField.Typ == AttributeTypeMap) {
		arrayItemDefinition = true
	}

	if !arrayItemDefinition {
		log.Debug().Str("name", f.Name).Str("path", pPath).Msg("start validating field")
	} else {
		log.Debug().Str("path", pPath).Msg("start validating array item")
	}

	if !arrayItemDefinition && !util.FieldNameWellFormed(f.Name) {
		return &CollectionDefError{ctx: pPath, msg: "field name is not well formed"}
	}

	if arrayItemDefinition {
		if len(f.Name) != 0 && f.Name != "[]" && f.Name != "%s" {
			return &CollectionDefError{ctx: pPath, msg: "field name is provided but is not required"}
		}
	} else {
		pPath = strings.Join([]string{pPath, f.Name}, ".")
	}

	if len(f.Tags) > 0 && len(f.Tags)%2 != 0 {
		return &CollectionDefError{ctx: pPath, msg: "tags array has to have an event number of entries"}
	}

	var vErr error = nil

	switch strings.ToLower(f.Typ) {
	case AttributeTypeObjectId:
	case AttributeTypeString:
	case AttributeTypeInt:
	case AttributeTypeLong:
	case AttributeTypeBool:
	case AttributeTypeDate:
	case AttributeTypeDocument:
	case AttributeTypeRefStruct:
		sn := f.StructRef.StructName
		if sn == "" {
			sn = f.StructName
			f.StructRef.StructName = sn
		}
		if sn == "" {
			vErr = &CollectionDefError{ctx: pPath, msg: "struct name required for " + f.Typ}
		}

		if f.StructRef.IsExternal && f.StructRef.Package == "" {
			log.Warn().Msg("struct reference declared external but missing package info " + f.StructRef.StructName)
		}
	case AttributeTypeStruct:
		if f.IsKey {
			vErr = &CollectionDefError{ctx: pPath, msg: "is-key not supported on " + f.Typ}
		} else {
			vErr = validateAttributes(f.Attributes, pPath, f)
		}

		if f.StructName == "" {
			if arrayItemDefinition {
				f.StructName = util.ToCapitalCase(parentField.Name) + "Struct"
			} else {
				f.StructName = util.ToCapitalCase(f.Name) + "Struct"
			}
			log.Debug().Str("path", pPath).Msg("no name provided for struct name...using " + f.StructName)
		}

	case AttributeTypeArray:
		if f.IsKey {
			vErr = &CollectionDefError{ctx: pPath, msg: "is-key not supported on " + f.Typ}
		} else {
			vErr = validateItem(f.Item, pPath, f)
		}
	case AttributeTypeMap:
		if f.IsKey {
			vErr = &CollectionDefError{ctx: pPath, msg: "is-key not supported on " + f.Typ}
		} else {
			vErr = validateItem(f.Item, pPath, f)
		}
	default:
		vErr = &CollectionDefError{ctx: pPath, msg: "unsupported type: " + f.Typ}
	}

	return vErr
}

func validateItem(f *Field, pPath string, containerField *Field) error {

	log.Debug().Str("path", pPath).Msg("start validating array item definition")

	if f == nil {
		return &CollectionDefError{ctx: pPath, msg: "no item provided for type array"}
	}

	/*
	 * The item gets assigned a special name as field. Sort of indexer.
	 */
	if containerField.Typ == "array" {
		f.Name = "[]"
		pPath = strings.Join([]string{pPath, "[i]"}, ".")
	} else {
		f.Name = "%s"
		pPath = strings.Join([]string{pPath, "%s"}, ".")
	}

	return validateField(f, pPath, containerField)
}
