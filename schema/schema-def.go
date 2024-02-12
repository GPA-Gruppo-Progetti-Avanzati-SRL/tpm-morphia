package schema

import (
	"errors"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strings"
)

type IncludeResolver interface {
	Resolve(fn string) string
	Load(fn string) ([]byte, error)
	ChangePath(fn string) IncludeResolver
}

type IncludePathResolver struct {
	basePath string
}

func (f IncludePathResolver) Load(fn string) ([]byte, error) {
	p := filepath.Join(f.basePath, fn)
	return os.ReadFile(p)
}

func (f IncludePathResolver) Resolve(fn string) string {
	p := filepath.Join(f.basePath, fn)
	return p
}

func (f IncludePathResolver) ChangePath(fn string) IncludeResolver {
	p := filepath.Join(f.basePath, fn)
	return &IncludePathResolver{basePath: p}
}

func NewPathResolver(fn string) IncludeResolver {
	return &IncludePathResolver{basePath: fn}
}

type StructResolver interface {
	IsDefined(structName string) bool
	GetStructByName(structName string) *StructDef
}

type EntityRef struct {
	Filename    string `json:"file-name,omitempty" yaml:"file-name,omitempty"`
	IsSubSchema bool   `json:"sub-schema,omitempty" yaml:"sub-schema,omitempty"`
	IsDocument  bool   `json:"document,omitempty" yaml:"document,omitempty"`
	Package     string `json:"package,omitempty" yaml:"package,omitempty"`
}

type Schema struct {
	ModuleName string `json:"module,omitempty" yaml:"module,omitempty"`
	// FolderPath string       `json:"folder-path,omitempty" yaml:"folder-path,omitempty"`
	Package    string       `json:"package,omitempty" yaml:"package,omitempty"`
	Imports    []string     `json:"imports,omitempty" yaml:"imports,omitempty"`
	EntityRefs []EntityRef  `json:"entities,omitempty" yaml:"entities,omitempty"`
	Structs    []*StructDef `json:"-" yaml:"-"`
}

func (c *Schema) IsDefined(n string) bool {
	for _, s := range c.Structs {
		if strings.ToLower(s.Name) == strings.ToLower(n) {
			return true
		}
	}

	return false
}

func (c *Schema) VisitStruct(n string, v Visitor) {
	n = strings.ToLower(n)
	for _, s := range c.Structs {
		if strings.ToLower(s.Name) == n {
			for i := range s.Attributes {
				v.visit(s.Attributes[i])
				s.Attributes[i].visit(v)
			}
		}
	}
}

func (c *Schema) GetStructByName(n string) *StructDef {
	for _, s := range c.Structs {
		if strings.ToLower(s.Name) == strings.ToLower(n) {
			return s
		}
	}

	return nil
}

func (c *Schema) finalize() error {
	const semLogContext = "tpm-morphia::validate-schema-def"
	log.Trace().Msg(semLogContext)

	// Doing a XOR...
	if c.Package == "" {
		return errors.New("please specify packageName")
	}

	if len(c.Structs) == 0 {
		return fmt.Errorf("no struct has been defined")
	}

	for _, s := range c.Structs {
		err := s.finalize(c)
		if err != nil {
			return err
		}
	}

	return nil
}

type StructDef struct {
	Name       string   `json:"name,omitempty" yaml:"name,omitempty"`
	Attributes []*Field `json:"attributes,omitempty" yaml:"attributes,omitempty"`
	Package    string   `json:"-" yaml:"-"`
	IsDocument bool     `json:"-" yaml:"-"`
	LoadedFrom string   `json:"-" yaml:"-"`
}

type StructReference struct {
	Name         string     `json:"name,omitempty" yaml:"name,omitempty"`
	IsExternal   bool       `json:"is-external,omitempty" yaml:"is-external,omitempty"`
	Package      string     `json:"package,omitempty" yaml:"package,omitempty"`
	StructDefRef *StructDef `json:"-" yaml:"-"`
}

func (s *StructDef) finalize(structResolver StructResolver) error {
	const semLogContext = "schema-struct::validate"
	log.Trace().Msg("start validating attributes")

	if len(s.Attributes) == 0 {
		return fmt.Errorf("attributes missing from schema")
	}

	for i := range s.Attributes {
		// Use the indexing property to address the underlay-ing object and not a copy.
		if e := s.Attributes[i].finalize(structResolver); e != nil {
			return e
		}
	}

	return nil
}

const (
	AttributeTypeArray    = "array"
	AttributeTypeMap      = "map"
	AttributeTypeStruct   = "struct"
	AttributeTypeString   = "string"
	AttributeTypeInt      = "int"
	AttributeTypeLong     = "long"
	AttributeTypeBool     = "bool"
	AttributeTypeDate     = "date"
	AttributeTypeObjectId = "object-id"
	AttributeTypeDocument = "document"
)

type Field struct {
	Name      string           `json:"name,omitempty" yaml:"name,omitempty"`
	Typ       string           `json:"type,omitempty" yaml:"type,omitempty"`
	StructRef *StructReference `json:"struct-ref,omitempty" yaml:"struct-ref,omitempty"`
	IsKey     bool             `json:"is-key,omitempty" yaml:"is-key,omitempty"`
	Tags      []Tag            `json:"tags,omitempty" yaml:"tags,omitempty"`
	Item      *Field           `json:"item,omitempty" yaml:"item,omitempty"`
	Options   string           `json:"options,omitempty" yaml:"options,omitempty"`
	Paths     []string         `json:"-" yaml:"-"`
	BsonPaths []string         `json:"-" yaml:"-"`
}

/*
* Field Methods
 */

func (f *Field) HasOption(o string) bool {
	return strings.Contains(f.Options, "with-filter")
}

func (f *Field) IsStructType() bool {
	return f.Typ == AttributeTypeStruct
}

func (f *Field) IsCollectionType() bool {
	return f.Typ == AttributeTypeMap || f.Typ == AttributeTypeArray
}

func (f *Field) IsValueType() bool {
	return !f.IsStructType() && !f.IsCollectionType()
}

func (f *Field) GetTagByName(tagName string) Tag {

	if len(f.Tags) > 0 {
		for _, t := range f.Tags {
			if strings.EqualFold(t.Name, tagName) {
				return t
			}
		}
	}

	return Tag{Name: tagName}
}

func (f *Field) GetTagValueByName(tagName string) (string, string) {

	if len(f.Tags) > 0 {
		for _, t := range f.Tags {
			if strings.EqualFold(t.Name, tagName) {
				arr := strings.Split(t.Value, ",")
				if len(arr) > 1 {
					return arr[0], arr[1]
				}
				return arr[0], ""
			}
		}
	}

	// Default to field name optional
	return f.Name, "omitempty"
}

func (f *Field) finalize(structResolver StructResolver) error {
	const semLogContext = "schema-field::validate"
	const isKeyNotSupportedErrorFormat = "is-key not supported on %s"
	const noItemDeclaredOnType = "no item declared on type %s"

	log.Trace().Str("name", f.Name).Msg(semLogContext + " start validating field")

	if !util.FieldNameWellFormed(f.Name) {
		return fmt.Errorf("field name is not well formed")
	}

	// setup the expected tags if missing
	f.FinalizeTags()

	var vErr error = nil

	switch strings.ToLower(f.Typ) {
	case AttributeTypeObjectId:
	case AttributeTypeString:
	case AttributeTypeInt:
	case AttributeTypeLong:
	case AttributeTypeBool:
	case AttributeTypeDate:
	case AttributeTypeDocument:
	case AttributeTypeStruct:
		if f.IsKey {
			vErr = fmt.Errorf(isKeyNotSupportedErrorFormat, f.Typ)
			return vErr
		}

		sn := f.StructRef.Name
		if sn == "" {
			vErr = fmt.Errorf("struct name required for %s", f.Typ)
			return vErr
		}

		if f.StructRef.IsExternal && f.StructRef.Package == "" {
			log.Warn().Msg("struct reference declared external but missing package info " + f.StructRef.Name)
			return vErr
		}

		if !f.StructRef.IsExternal {
			structRef := structResolver.GetStructByName(sn)
			if structRef == nil {
				err := fmt.Errorf("cannot resolve struct reference %s", f.StructRef.Name)
				log.Error().Err(err).Msg(semLogContext)
				return err
			}
			f.StructRef.StructDefRef = structRef
		}
	case AttributeTypeArray:
		if f.IsKey {
			vErr = fmt.Errorf(isKeyNotSupportedErrorFormat, f.Typ)
			return vErr
		}

		if f.Item == nil {
			vErr = fmt.Errorf(noItemDeclaredOnType, f.Typ)
			return vErr
		}

		vErr = f.finalizeItem(strings.ToLower(f.Typ), structResolver)
	case AttributeTypeMap:
		if f.IsKey {
			vErr = fmt.Errorf(isKeyNotSupportedErrorFormat, f.Typ)
			return vErr
		}

		if f.Item == nil {
			vErr = fmt.Errorf(noItemDeclaredOnType, f.Typ)
			return vErr
		}

		vErr = f.finalizeItem(strings.ToLower(f.Typ), structResolver)
	default:
		vErr = fmt.Errorf("unsupported type: %s", f.Typ)
	}

	return vErr
}

func (f *Field) FinalizeTags() {

	const omitempty = ",omitempty"
	const jsonTag = "json"
	const yamlTg = "yaml"
	const bsonTag = "bson"

	tagFieldName := f.Name
	if len(f.Tags) == 0 {
		f.Tags = []Tag{
			NewTag(jsonTag, f.Name+omitempty),
			NewTag(bsonTag, tagFieldName+omitempty),
			NewTag(yamlTg, tagFieldName+omitempty),
		}
	}

	if t := FindTag(f.Tags, jsonTag); t.Name == "" {
		f.Tags = append(f.Tags, NewTag(jsonTag, tagFieldName+omitempty))
	}

	if t := FindTag(f.Tags, yamlTg); t.Name == "" {
		f.Tags = append(f.Tags, NewTag(yamlTg, tagFieldName+omitempty))
	}

	if t := FindTag(f.Tags, bsonTag); t.Name == "" {
		f.Tags = append(f.Tags, NewTag(bsonTag, tagFieldName+omitempty))
	}
}

func (f *Field) visit(v Visitor) {
	v.startVisit(f)

	// Se il riferimento e' esterno allora non lo 'seguo'. L'elemento e' nil.
	switch f.Typ {
	case AttributeTypeStruct:
		if f.StructRef.StructDefRef != nil {
			for i := range f.StructRef.StructDefRef.Attributes {
				v.visit(f.StructRef.StructDefRef.Attributes[i])
				f.StructRef.StructDefRef.Attributes[i].visit(v)
			}
		}
	case AttributeTypeArray:
		fallthrough
	case AttributeTypeMap:
		v.visit(f.Item)
		f.Item.visit(v)
	}

	v.endVisit(f)
}

func (f *Field) finalizeItem(parentAttributeType string, structResolver StructResolver) error {
	const semLogContext = "schema-field::finalize-item"
	log.Trace().Str("path", f.Name).Msg("start validating array item definition")

	f.Item.Name = "[]"
	if parentAttributeType == AttributeTypeMap {
		f.Item.Name = "%s"
	}

	switch f.Item.Typ {
	case AttributeTypeObjectId:
	case AttributeTypeString:
	case AttributeTypeInt:
	case AttributeTypeLong:
	case AttributeTypeBool:
	case AttributeTypeDate:
	case AttributeTypeDocument:
	case AttributeTypeStruct:
		structRef := structResolver.GetStructByName(f.Item.StructRef.Name)
		if structRef == nil {
			err := fmt.Errorf("cannot resolve struct reference %s", f.Item.StructRef.Name)
			log.Error().Err(err).Msg(semLogContext)
			return err
		}
		f.Item.StructRef.StructDefRef = structRef
	case AttributeTypeMap:
		err := f.Item.finalizeItem(AttributeTypeMap, structResolver)
		if err != nil {
			log.Error().Err(err).Msg(semLogContext)
		}
		return err
	case AttributeTypeArray:
		err := f.Item.finalizeItem(AttributeTypeArray, structResolver)
		/*		err := fmt.Errorf("unsupported item type %s", f.Item.Typ)
				log.Error().Err(err).Msg(semLogContext)
		*/
		if err != nil {
			log.Error().Err(err).Msg(semLogContext)
		}
		return err
	default:
		err := fmt.Errorf("item type not recognized %s", f.Item.Typ)
		log.Error().Err(err).Msg(semLogContext)
		return err
	}

	return nil
}

func (f *Field) Package() (string, string) {

	// The actual field to check the package against
	var f1 *Field

	switch f.Typ {
	case AttributeTypeStruct:
		f1 = f
	case AttributeTypeMap:
		fallthrough
	case AttributeTypeArray:
		f1 = f
		for f1.Typ == AttributeTypeArray || f1.Typ == AttributeTypeMap {
			f1 = f1.Item
		}

		if f1.Typ != AttributeTypeStruct {
			f1 = nil
		}
	}

	var pkg string
	var sn string
	if f1 != nil {
		sn = f1.StructRef.Name
		if f1.StructRef.IsExternal {
			pkg = f1.StructRef.Package
		} else {
			pkg = f1.StructRef.StructDefRef.Package
		}
	}

	return sn, pkg
}

/*
func (f *Field) computePackageOfAttribute(currentPackage string) string {

	// The actual field to check the package against
	var f1 *Field

	switch f.Typ {
	case AttributeTypeStruct:
		f1 = f
	case AttributeTypeMap:
		fallthrough
	case AttributeTypeArray:
		if f.Item.Typ == AttributeTypeStruct {
			f1 = f.Item
		}
	}

	var pkg string
	if f1 != nil {
		if f1.StructRef.IsExternal {
			pkg = f1.StructRef.Package
		} else {
			attrPackage := f1.StructRef.StructDefRef.Package
			if attrPackage != currentPackage {
				pkg = attrPackage
			}
		}
	}

	return pkg
}
*/
//

//

//
//type CollectionProps struct {
//	StructName     string `json:"struct-name,omitempty"`
//	MorphiaPackage string `json:"morphia-pkg,omitempty"`
//}
//
//type Collection struct {

//	Properties CollectionProps `json:"properties,omitempty"`
//	Attributes []Field         `json:"attributes,omitempty"`
//
//	AllAttributes []*Field `json:"-"`
//}
//

//
//type Field struct {
//	Name       string          `json:"name,omitempty"`
//	StructName string          `json:"struct-name,omitempty"`
//	Typ        string          `json:"type,omitempty"`
//	IsKey      bool            `json:"is-key,omitempty"`
//	Tags       []string        `json:"tags,omitempty"`
//	Attributes []Field         `json:"attributes,omitempty"`
//	Item       *Field          `json:"item,omitempty"`
//	Queryable  bool            `json:"queryable,omitempty"`
//	StructRef  StructReference `json:"struct-ref,omitempty"`
//	Paths      []string        `json:"-"`
//	BsonPaths  []string        `json:"-"`
//	Options    string          `json:"options,omitempty"`
//}
//
//type InfoCollectorVisitor struct {
//	Attributes []*Field
//}
//
///*type VisitPhase int
//const (
//  StartVisit VisitPhase = iota
//  DoVisit
//  EndVisit
//)
//
//type VisitorState interface {
//
//}
//
//*/
//
//type Visitor interface {
//	startVisit(f *Field)
//	visit(f *Field)
//	endVisit(f *Field)
//}
//

//
///*
// * Collection Methods
// */
//func (c *Collection) ToJsonString(prefix string, indent string) string {
//
//	if s, e := json.MarshalIndent(c, prefix, indent); e != nil {
//		return e.Error()
//	} else {
//		return string(s)
//	}
//}
//
//func (c *Collection) findAttributes() []*Field {
//	infoCollector := InfoCollectorVisitor{Attributes: make([]*Field, 0)}
//	for i := range c.Attributes {
//		infoCollector.Attributes = append(infoCollector.Attributes, &c.Attributes[i])
//		infoCollector = c.Attributes[i].FindAttributes(infoCollector)
//	}
//
//	return infoCollector.Attributes
//}
//
//func (c *Collection) visit(v Visitor) {
//	for i := range c.Attributes {
//		v.visit(&c.Attributes[i])
//		c.Attributes[i].visit(v)
//	}
//}
//
//func (c *Collection) wireReference2Structs(fields []*Field) error {
//	refs := make(map[string]*Field)
//	for _, f := range fields {
//		if f.Typ == AttributeTypeStruct {
//			refs[f.StructName] = f
//		} else if (f.Typ == AttributeTypeMap || f.Typ == AttributeTypeArray) && f.Item.Typ == AttributeTypeStruct {
//			refs[f.Item.StructName] = f.Item
//		}
//	}
//
//	for _, f := range fields {
//		if f.Typ == AttributeTypeRefStruct {
//			if item, ok := refs[f.StructRef.StructName]; !ok {
//				// Unresolved reference. Can be declared externally to current file.
//				// Another approach could be to process two schemas as a bundle but is way to elaborate.
//				// Should probably add a dependency in terms of a package import
//				if !f.StructRef.IsExternal {
//					// TODO: the f.StructName in the message should be replaced with f.StructRef.StructName
//					return errors.New(fmt.Sprintf("the field %s refers to undefined struct %s", f.Name, f.StructName))
//				}
//			} else {
//				f.StructRef.Item = item
//			}
//		}
//	}
//
//	return nil
//}
//

//
///*
// *
// */
//func (f *Field) FindAttributes(infoCollector InfoCollectorVisitor) InfoCollectorVisitor {
//
//	for i := range f.Attributes {
//		infoCollector.Attributes = append(infoCollector.Attributes, &f.Attributes[i])
//		infoCollector = f.Attributes[i].FindAttributes(infoCollector)
//	}
//
//	if f.Item != nil {
//		infoCollector = f.Item.FindAttributes(infoCollector)
//	}
//
//	return infoCollector
//}
//
//func (f *Field) visit(v Visitor) {
//	v.startVisit(f)
//
//	// Se il riferimento e' esterno allora non lo 'seguo'. L'elemento e' nil.
//	if f.Typ == AttributeTypeRefStruct {
//		if f.StructRef.Item != nil {
//			for i := range f.StructRef.Item.Attributes {
//				v.visit(&f.StructRef.Item.Attributes[i])
//				f.StructRef.Item.Attributes[i].visit(v)
//			}
//		}
//	} else {
//		for i := range f.Attributes {
//			v.visit(&f.Attributes[i])
//			f.Attributes[i].visit(v)
//		}
//
//		if f.Item != nil {
//			f.Item.visit(v)
//		}
//	}
//
//	v.endVisit(f)
//
//}
//
//func (f Field) String() string {
//	return fmt.Sprintf("%s - %s - %s", f.Name, f.Typ, f.StructName)
//}
//
///*
// * Read and Validate Function
// */
