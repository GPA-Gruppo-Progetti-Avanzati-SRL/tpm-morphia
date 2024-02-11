package schema_test

import (
	"fmt"
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

var ex6BasePath = "../examples/example6"
var ex7BasePath = "../examples/example7/schema"

type fmtVisitor struct {
}

func (v *fmtVisitor) startVisit(f *schema.Field) {
	const semLogContext = "fmt-visitor::start-visit"
	log.Info().Str("field-name", f.Name).Msg(semLogContext)
}

func (v *fmtVisitor) visit(f *schema.Field) {
	const semLogContext = "fmt-visitor::visit"
	log.Info().Str("field-name", f.Name).Msg(semLogContext)
}

func (v *fmtVisitor) endVisit(f *schema.Field) {
	const semLogContext = "fmt-visitor::end-visit"
	log.Info().Str("field-name", f.Name).Msg(semLogContext)
}

func TestReadSchema(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	b, err := os.ReadFile(filepath.Join(ex6BasePath, "schema.json"))
	require.NoError(t, err)

	sch, err := schema.ReadSchemaDefinitionFromBuffer(schema.JSONFormat, b, schema.IncludeResolver(schema.NewPathResolver(ex6BasePath)))
	require.NoError(t, err)

	yamlSchema, err := yaml.Marshal(sch)
	require.NoError(t, err)

	fmt.Println(string(yamlSchema))
	pv := schema.PathFinderVisitor{}
	sch.VisitStruct("author", &pv)
	for _, a := range pv.Attributes {
		t.Log(a.Name, a.Paths)
	}
}

func TestConvertSchema(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	b, err := os.ReadFile(filepath.Join(ex6BasePath, "schema.json"))
	require.NoError(t, err)

	sch, err := schema.ReadSchemaDefinitionFromBuffer(schema.JSONFormat, b, schema.IncludeResolver(schema.NewPathResolver(ex6BasePath)))
	require.NoError(t, err)

	for i := range sch.EntityRefs {
		sch.EntityRefs[i].Filename = strings.TrimSuffix(sch.EntityRefs[i].Filename, ".json") + ".yml"
	}
	yamlSchema, err := yaml.Marshal(sch)
	require.NoError(t, err)

	yamlFile := filepath.Join(ex6BasePath, "schema.yml")
	err = os.WriteFile(yamlFile, yamlSchema, fs.ModePerm)
	require.NoError(t, err)

	for _, s := range sch.Structs {
		yamlStruct, err := yaml.Marshal(s)
		require.NoError(t, err)

		yamlFile = filepath.Join(ex6BasePath, s.LoadedFrom)
		yamlFile = strings.TrimSuffix(yamlFile, ".json") + ".yml"
		err = os.WriteFile(yamlFile, yamlStruct, fs.ModePerm)
		require.NoError(t, err)

		t.Logf("structure %s loaded from %s", s.Name, s.LoadedFrom)
	}
}

func TestReadSchemaWithImports(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	b, err := os.ReadFile(filepath.Join(ex7BasePath, "schema.yml"))
	require.NoError(t, err)

	sch, err := schema.ReadSchemaDefinitionFromBuffer(schema.YAMLFormat, b, schema.IncludeResolver(schema.NewPathResolver(ex7BasePath)))
	require.NoError(t, err)

	for _, a := range sch.Structs {
		t.Log(a.Name, a.LoadedFrom)
	}
}
