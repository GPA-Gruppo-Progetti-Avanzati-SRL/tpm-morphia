package schema

import (
	"encoding/json"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system/util"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"io"
	"strings"
)

const (
	AttributeTypeArray    = "array"
	AttributeTypeStruct   = "struct"
	AttributeTypeString   = "string"
	AttributeTypeInt      = "int"
	AttributeTypeObjectId = "object-id"
)

type CollectionDefError struct {
	ctx string
	msg string
}

func (ce *CollectionDefError) Error() string {
	return fmt.Sprintf("[%s] - %s", ce.ctx, ce.msg)
}

type CollectionProps struct {
	FolderPath  string `json:"folderPath,omitempty"`
	PackageName string `json:"packageName,omitempty"`
	StructName  string `json:"struct-name,omitempty"`
}

type Collection struct {
	Name       string          `json:"name,omitempty"`
	Properties CollectionProps `json:"properties,omitempty"`
	Attributes []Field         `json:"attributes,omitempty"`
}

type Field struct {
	Name       string   `json:"name,omitempty"`
	StructName string   `json:"struct-name,omitempty"`
	Typ        string   `json:"type,omitempty"`
	IsKey      bool     `json:"is-key,omitempty"`
	Tags       []string `json:"tags,omitempty"`
	Attributes []Field  `json:"attributes,omitempty"`
	Item       *Field   `json:"item,omitempty"`
	Queryable  bool     `json:"queryable,omitempty"`
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
func (sch *Collection) ToJsonString(prefix string, indent string) string {

	if s, e := json.MarshalIndent(sch, prefix, indent); e != nil {
		return e.Error()
	} else {
		return string(s)
	}
}

/*
 * Field Methods
 */
func (f *Field) GetTagsAsListOfTag(forceFieldNameIfMissing bool, forceOmitIfEmpty bool, tagsOfInterest []string) []Tag {
	var tags []Tag
	if len(f.Tags) > 0 {
		tags = make([]Tag, 0, len(f.Tags)/2)
		for i := 0; i < len(f.Tags); i += 2 {
			nt := NewTag(f.Tags[i], f.Tags[i+1])
			if nt.Value == "" && forceFieldNameIfMissing {
				nt.Value = f.Name
			}

			if !strings.Contains(f.Tags[i+1], "omitempty") && forceOmitIfEmpty {
				nt.Options = append(nt.Options, "omitempty")
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
 * Read and Validate Function
 */
func ReadCollectionDefinition(logger log.Logger, reader io.Reader) (*Collection, error) {

	_ = level.Debug(logger).Log("msg", "start reading collection definition")

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

	if e := validateCollectionDef(logger, schema); e != nil {
		return nil, e
	}

	return schema, nil
}

func validateCollectionDef(logger log.Logger, c *Collection) error {

	_ = level.Debug(logger).Log("msg", "start validating collection definition", "name", c.Name)

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

	if e := validateProperties(logger, c); e != nil {
		return e
	}

	return validateAttributes(logger, c.Attributes, pPath, nil)
}

func validateProperties(logger log.Logger, c *Collection) error {

	cp := &c.Properties
	// Doing a XOR...
	if (cp.PackageName != "") == (cp.FolderPath != "") {
		return &CollectionDefError{msg: fmt.Sprintf("one (and only one) of packageName (%s) and folderPath (%s) have to be specified", cp.PackageName, cp.FolderPath)}
	}

	if cp.StructName == "" {
		cp.StructName = util.ToCapitalCase(c.Name)
		_ = level.Debug(logger).Log("msg", "no name provided for struct name...using "+cp.StructName)
	}

	return nil
}

func validateAttributes(logger log.Logger, props []Field, pPath string, parentField *Field) error {

	_ = level.Debug(logger).Log("msg", "start validating properties", "path", pPath)

	if len(props) == 0 {
		return &CollectionDefError{ctx: pPath, msg: "attributes missing from schema"}
	}

	for i := range props {
		// Use the indexing property to address the underlay-ing object and not a copy.
		if e := validateField(logger, &props[i], pPath, parentField); e != nil {
			return e
		}
	}

	return nil
}

func validateField(logger log.Logger, f *Field, pPath string, parentField *Field) error {

	arrayItemDefinition := false

	if parentField != nil && parentField.Typ == AttributeTypeArray {
		arrayItemDefinition = true
	}

	if !arrayItemDefinition {
		_ = level.Debug(logger).Log("msg", "start validating field", "name", f.Name, "path", pPath)
	} else {
		_ = level.Debug(logger).Log("msg", "start validating array item", "path", pPath)
	}

	if !arrayItemDefinition && !util.FieldNameWellFormed(f.Name) {
		return &CollectionDefError{ctx: pPath, msg: "field name is not well formed"}
	}

	if arrayItemDefinition {
		if len(f.Name) != 0 && f.Name != "[]" {
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
	case AttributeTypeStruct:
		if f.IsKey {
			vErr = &CollectionDefError{ctx: pPath, msg: "is-key not supported on " + f.Typ}
		} else {
			vErr = validateAttributes(logger, f.Attributes, pPath, f)
		}

		if f.StructName == "" {
			if arrayItemDefinition {
				f.StructName = util.ToCapitalCase(parentField.Name) + "Struct"
			} else {
				f.StructName = util.ToCapitalCase(f.Name) + "Struct"
			}
			_ = level.Debug(logger).Log("msg", "no name provided for struct name...using "+f.StructName, "path", pPath)
		}

	case AttributeTypeArray:
		if f.IsKey {
			vErr = &CollectionDefError{ctx: pPath, msg: "is-key not supported on " + f.Typ}
		} else {
			vErr = validateItem(logger, f.Item, pPath, f)
		}
	default:
		vErr = &CollectionDefError{ctx: pPath, msg: "unsupported type: " + f.Typ}
	}

	return vErr
}

func validateItem(logger log.Logger, f *Field, pPath string, arrayField *Field) error {

	_ = level.Debug(logger).Log("msg", "start validating array item definition", "path", pPath)

	if f == nil {
		return &CollectionDefError{ctx: pPath, msg: "no item provided for type array"}
	}

	/*
	 * The item gets assigned a special name as field. Sort of indexer.
	 */
	f.Name = "[]"

	pPath = strings.Join([]string{pPath, "[i]"}, ".")
	return validateField(logger, f, pPath, arrayField)
}
