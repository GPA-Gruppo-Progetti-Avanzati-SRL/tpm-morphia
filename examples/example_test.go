package examples_test

import (
	"bytes"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/config"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/gen/mongodb"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schema"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system"
	"github.com/go-kit/kit/log/level"
	"io/ioutil"
	"testing"
)

func TestExamples(t *testing.T) {

	genCfg := config.DefaultConfig
	genCfg.TargetDirectory = "."
	genCfg.ResourceDirectory = ".."
	genCfg.FormatCode = true
	genCfg.CollectionDefScanPath = "."

	logger := system.GetLogger()

	colls, err := genCfg.FindCollectionToProcess(logger)
	if err != nil {
		_ = level.Error(logger).Log("msg", "FindCollectionToProcess error", "err", err.Error())
		return
	}

	for _, collDefFile := range colls {

		if schemaFile, err := ioutil.ReadFile(collDefFile); err != nil {
			_ = level.Error(logger).Log("msg", "ioutil.ReadFile error", "err", err.Error())
			return
		} else {

			r := bytes.NewReader([]byte(schemaFile))
			schema, e := schema.ReadCollectionDefinition(system.GetLogger(), r)
			if e != nil {
				_ = level.Error(logger).Log("msg", "schema.ReadCollectionDefinition error", "err", e.Error())
				return
			}

			genDriver, e := mongodb.NewCodeGenCollection(schema)
			if e != nil {
				_ = level.Error(logger).Log("msg", "mongodb.NewCodeGenCollection error", "err", e.Error())
				return
			} else {
				e = mongodb.Generate(system.GetLogger(), &genCfg, genDriver)
				if e != nil {
					_ = level.Error(logger).Log("msg", "mongodb.Generate error", "err", e.Error())
					return
				}
			}
		}
	}
}
