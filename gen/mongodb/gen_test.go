package mongodb

import (
	"bytes"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/config"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schema"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system"
	"io/ioutil"
	"os"
	"testing"
)

var tests_datatypes = "../../tests/datatypes-tpmm.json"
var tests_cliente = "../../tests/cliente-tpmm.json"

func TestGeneration(t *testing.T) {

	f, err := os.Open(tests_cliente)
	// b, err := ioutil.ReadFile(tests_datatypes)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	schema, e := schema.ReadCollectionDefinition(system.GetLogger(), f)
	if e != nil {
		t.Fatal(e)
	}

	genCfg := config.Config{Version: "v2", TargetDirectory: "../../tests", ResourceDirectory: "../..", FormatCode: false}
	genDriver, e := NewCodeGenCollection(schema)
	if e != nil {
		t.Error(e)
	} else {
		e = Generate(system.GetLogger(), &genCfg, genDriver)
		if e != nil {
			t.Error(e)
		}
	}
}

func TestCodeGenTree(t *testing.T) {

	b, err := ioutil.ReadFile(tests_datatypes)
	if err != nil {
		t.Fatal(err)
	}

	r := bytes.NewReader(b)
	schema, e := schema.ReadCollectionDefinition(system.GetLogger(), r)
	if e != nil {
		t.Error(e)
	}

	genDriver, e := NewCodeGenCollection(schema)
	if e != nil {
		t.Error(e)
	}

	fmt.Println(genDriver)
}
