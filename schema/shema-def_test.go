package schema_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schema"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"io/fs"
	"os"
	"path/filepath"
	"sigs.k8s.io/yaml"
	"strings"
	"testing"
)

const ex6BasePath = "../examples/example6"
const ex7BasePath = "../examples/example7/schema"

type fmtVisitor struct {
}

const semLogFieldName = "field-name"

func (v *fmtVisitor) startVisit(f *schema.Field) {
	const semLogContext = "fmt-visitor::start-visit"
	log.Info().Str(semLogFieldName, f.Name).Msg(semLogContext)
}

func (v *fmtVisitor) visit(f *schema.Field) {
	const semLogContext = "fmt-visitor::visit"
	log.Info().Str(semLogFieldName, f.Name).Msg(semLogContext)
}

func (v *fmtVisitor) endVisit(f *schema.Field) {
	const semLogContext = "fmt-visitor::end-visit"
	log.Info().Str(semLogFieldName, f.Name).Msg(semLogContext)
}

func TestReadSchema(t *testing.T) {
	const basePath = ex7BasePath
	const schemaName = "schema.yml"
	const schemaFormat = schema.YAMLFormat

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	b, err := os.ReadFile(filepath.Join(basePath, schemaName))
	require.NoError(t, err)

	sch, err := schema.ReadSchemaDefinitionFromBuffer(schemaFormat, b, schema.IncludeResolver(schema.NewPathResolver(basePath)))
	require.NoError(t, err)

	for _, a := range sch.Structs {
		t.Log(a.Name, a.LoadedFrom)
	}
}

func TestReadSchemaAndShowPaths(t *testing.T) {
	const basePath = ex7BasePath
	const schemaName = "schema.yml"
	const schemaFormat = schema.YAMLFormat

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	b, err := os.ReadFile(filepath.Join(basePath, schemaName))
	require.NoError(t, err)

	sch, err := schema.ReadSchemaDefinitionFromBuffer(schemaFormat, b, schema.IncludeResolver(schema.NewPathResolver(basePath)))
	require.NoError(t, err)

	//yamlSchema, err := yaml.Marshal(sch)
	//require.NoError(t, err)

	// fmt.Println(string(yamlSchema))

	ShowPaths(t, sch, "author")
}

func TestConvertSchema(t *testing.T) {
	const basePath = ex6BasePath
	const schemaName = "schema.json"
	const schemaFormat = schema.JSONFormat

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	b, err := os.ReadFile(filepath.Join(basePath, schemaName))
	require.NoError(t, err)

	sch, err := schema.ReadSchemaDefinitionFromBuffer(schemaFormat, b, schema.IncludeResolver(schema.NewPathResolver(basePath)))
	require.NoError(t, err)

	for i := range sch.EntityRefs {
		sch.EntityRefs[i].Filename = strings.TrimSuffix(sch.EntityRefs[i].Filename, ".json") + ".yml"
	}
	yamlSchema, err := yaml.Marshal(sch)
	require.NoError(t, err)

	yamlFile := filepath.Join(basePath, "schema.yml")
	err = os.WriteFile(yamlFile, yamlSchema, fs.ModePerm)
	require.NoError(t, err)

	for _, s := range sch.Structs {
		yamlStruct, err := yaml.Marshal(s)
		require.NoError(t, err)

		yamlFile = filepath.Join(basePath, s.LoadedFrom)
		yamlFile = strings.TrimSuffix(yamlFile, ".json") + ".yml"
		err = os.WriteFile(yamlFile, yamlStruct, fs.ModePerm)
		require.NoError(t, err)

		t.Logf("structure %s loaded from %s", s.Name, s.LoadedFrom)
	}
}

func ShowPaths(t *testing.T, sch *schema.Schema, entityName string) {
	pv := schema.PathFinderVisitor{}
	sch.VisitStruct(entityName, &pv)
	for _, a := range pv.Attributes {
		t.Logf("[%s] - %s / %s", a.Name, strings.Join(a.Paths, ";"), strings.Join(a.BsonPaths, ";"))
	}
}
