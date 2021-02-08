package examples

import (
	"bytes"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/config"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/gen/mongodb"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schema"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system"
	"io/ioutil"
	"testing"
)

func TestExample(t *testing.T) {

	genCfg := config.DefaultConfig
	genCfg.TargetDirectory = "."
	genCfg.ResourceDirectory = ".."
	genCfg.FormatCode = true
	genCfg.CollectionDefScanPath = "."

	colls, err := genCfg.FindCollectionToProcess()
	if err != nil {
		t.Fatal(err)
	}

	for _, collDefFile := range colls {

		if schemaFile, err := ioutil.ReadFile(collDefFile); err != nil {
			t.Fatal(err)
		} else {

			r := bytes.NewReader([]byte(schemaFile))
			schema, e := schema.ReadCollectionDefinition(system.GetLogger(), r)
			if e != nil {
				t.Error(e)
			}

			genDriver, e := mongodb.NewCodeGenCollection(schema)
			if e != nil {
				t.Error(e)
			} else {
				e = mongodb.Generate(system.GetLogger(), &genCfg, genDriver)
				if e != nil {
					t.Error(e)
				}
			}
		}

	}

}
