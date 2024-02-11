package gomongodb_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/gen/gomongodb"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schema"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

var ex6BasePath = "../../examples/example6"
var ex7BasePath = "../../examples/example7/schema"

func TestGenerate(t *testing.T) {
	const FormatCode = true
	const schemaFileName = "schema.yml"
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	b, err := os.ReadFile(filepath.Join(ex7BasePath, schemaFileName))
	require.NoError(t, err)

	sch, err := schema.ReadSchemaDefinitionFromBuffer(schema.FormatOfFile(schemaFileName), b, schema.IncludeResolver(schema.NewPathResolver(ex7BasePath)))
	if err != nil {
		t.Error(err)
	}

	const targetFolder = "../.."
	cfg := gomongodb.GeneratorConfig{
		Schema:       sch,
		TargetFolder: targetFolder,
		EntityName:   "address",
		Version:      "tpm-morphia@v0.0.0",
		FormatCode:   FormatCode,
	}
	err = gomongodb.GenerateEntity(&cfg)
	require.NoError(t, err)

	cfg = gomongodb.GeneratorConfig{
		Schema:       sch,
		TargetFolder: targetFolder,
		EntityName:   "book",
		Version:      "tpm-morphia@v0.0.0",
		FormatCode:   FormatCode,
	}
	err = gomongodb.GenerateEntity(&cfg)
	require.NoError(t, err)

	cfg = gomongodb.GeneratorConfig{
		Schema:       sch,
		TargetFolder: targetFolder,
		EntityName:   "author",
		Version:      "tpm-morphia@v0.0.0",
		FormatCode:   FormatCode,
	}
	err = gomongodb.GenerateEntity(&cfg)
	require.NoError(t, err)

}
