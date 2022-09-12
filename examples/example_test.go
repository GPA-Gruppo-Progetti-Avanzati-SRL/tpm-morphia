package examples_test

import (
	"bytes"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/config"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/gen/mongodb"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schema"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestExamples(t *testing.T) {

	genCfg := config.DefaultConfig
	genCfg.TargetDirectory = "."
	genCfg.ResourceDirectory = ".."
	genCfg.FormatCode = true
	genCfg.CollectionDefScanPath = "."

	colls, err := genCfg.FindCollectionToProcess()
	require.NoError(t, err)

	for _, collDefFile := range colls {
		schemaFile, err := ioutil.ReadFile(collDefFile)
		require.NoError(t, err)

		r := bytes.NewReader([]byte(schemaFile))
		schema, err := schema.ReadCollectionDefinition(r)
		require.NoError(t, err)

		genDriver, err := mongodb.NewCodeGenCollection(schema)
		require.NoError(t, err)
		err = mongodb.Generate(&genCfg, genDriver)
		require.NoError(t, err)
	}
}
