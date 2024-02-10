package schema

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"path/filepath"
	yaml "sigs.k8s.io/yaml/goyaml.v3"
)

type Format string

const (
	NotImplementedFormat Format = "not-implemented"
	JSONFormat           Format = "json"
	YAMLFormat           Format = "yaml"
)

var FileExtensionFormatMap = map[string]Format{
	".yml":  YAMLFormat,
	".yaml": YAMLFormat,
	".json": JSONFormat,
}

func FormatOfFile(fn string) Format {
	ext := filepath.Ext(fn)
	f, ok := FileExtensionFormatMap[ext]
	if !ok {
		f = NotImplementedFormat
	}

	return f
}

func ReadSchemaDefinition(f Format, def []byte, includeResolver IncludeResolver) (*Schema, error) {
	const semLogContext = "tpm-morphia::read-schema-def"
	var err error

	log.Debug().Msg(semLogContext)

	schema, err := UnmarshalSchemaFromBuffer(includeResolver, f, def)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return nil, err
	}

	refs, err := ResolveImports(includeResolver, schema)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return nil, err
	}

	schema.EntityRefs = append(schema.EntityRefs, refs...)
	for _, i := range schema.EntityRefs {

		s, err := unmarshalStructDefFromFile(includeResolver, i.Filename)
		if err != nil {
			log.Error().Err(err).Msg(semLogContext)
			return nil, err
		}

		s.IsDocument = i.IsDocument
		s.Package = schema.Package
		s.LoadedFrom = i.Filename
		if i.Package != "" {
			s.Package = i.Package
		}
		schema.Structs = append(schema.Structs, s)
	}

	/*  Alternate: not able to force errors on unknown fields
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&schema); err != nil {
		return nil, &DefinitionError{msg: err.Error()}
	}
	*/

	/* Since I just deserialized stuff, just try to check that pretty much che stuff is correct or at least
	 * not massively wrong.
	 */
	if e := schema.finalize(); e != nil {
		return nil, e
	}

	///* Now traverse the tree and get the list of attrs. yes, it could have been done in the validation phase.
	// * May be later on will get into that processing.
	// */
	//fields := schema.findAttributes()
	//schema.AllAttributes = fields
	//
	///*
	// * Now it's time to wire the references to structs in order to calculate all the paths or hierarchies to the leaves.
	// * In principle the wiring should detect loops in the config. Loops at the moment might cause a stack issue, but nevertheless
	// * should be handled as pointer to struct in the generation.
	// */
	//if err := schema.wireReference2Structs(fields); err != nil {
	//	return nil, err
	//}
	//
	//lf := PathFinderVisitor{}
	//schema.visit(&lf)

	return schema, nil
}

func UnmarshalSchemaFromBuffer(resolver IncludeResolver, f Format, data []byte) (*Schema, error) {
	const semLogContext = "tpm-morphia::unmarshal-schema-def"
	var err error

	schema := &Schema{}

	switch f {
	case JSONFormat:
		err = json.Unmarshal(data, schema)
	case YAMLFormat:
		err = yaml.Unmarshal(data, schema)
	default:
		err = fmt.Errorf("unsupported file format %s", f)
	}

	for i := range schema.EntityRefs {
		if schema.EntityRefs[i].Package == "" {
			schema.EntityRefs[i].Package = schema.Package
		}
	}

	return schema, err
}

func UnmarshalSchemaFromFile(resolver IncludeResolver, fn string) (*Schema, error) {
	const semLogContext = "tpm-morphia::unmarshal-schema-def"
	var err error

	b, err := resolver.Load(fn)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return nil, err
	}

	return UnmarshalSchemaFromBuffer(resolver, FormatOfFile(fn), b)
}

func ResolveImports(resolver IncludeResolver, sch *Schema) ([]EntityRef, error) {
	const semLogContext = "tpm-morphia::resolve-imports"

	var refs []EntityRef
	/*	for _, e := range sch.EntityRefs {
		e.Filename = resolver.Resolve(e.Filename)
		refs = append(refs, e)
	}*/

	for _, i := range sch.Imports {
		sch1, err := UnmarshalSchemaFromFile(resolver, i)
		if err != nil {
			return nil, err
		}

		nestedResolver := resolver.ChangePath(filepath.Dir(i))
		for _, e := range sch1.EntityRefs {
			newPath := filepath.Join(filepath.Dir(i), e.Filename)
			e.Filename = newPath //nestedResolver.Resolve(e.Filename)
			refs = append(refs, e)
		}

		subRefs, err := ResolveImports(nestedResolver, sch1)
		if err != nil {
			return nil, err
		}

		refs = append(refs, subRefs...)
	}

	return refs, nil
}

func unmarshalStructDefFromFile(resolver IncludeResolver, fn string) (*StructDef, error) {
	const semLogContext = "tpm-morphia::unmarshal-struct-def"
	var err error

	b, err := resolver.Load(fn)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
	}

	structDef := &StructDef{}
	switch FormatOfFile(fn) {
	case JSONFormat:
		err = json.Unmarshal(b, structDef)
	case YAMLFormat:
		err = yaml.Unmarshal(b, structDef)
	default:
		err = fmt.Errorf("unsupported file format %s", FormatOfFile(fn))
	}

	return structDef, err
}
