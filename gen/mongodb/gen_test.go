package mongodb

import (
	"bytes"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/config"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schema"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system"
	"testing"
)

var example = `{
   "name": "author",
   "properties": {
       "folderPath": "./author"
      ,"struct-name": "Author"
   },
   "attributes": [
       { "name": "firstName", "type": "string", "tags": [ "json", "fn", "bson", "fn" ], "queryable": true }
      ,{ "name": "lastName", "type": "string", "tags": [ "json", "ln", "bson", "ln" ], "queryable": true }
      ,{ "name": "age", "type": "int", "queryable": true }
      ,{ "name": "address",  "type": "struct", "tags": [ "json", "addr", "bson", "addr" ]
         ,"struct-name": "Address"
         ,"attributes": [
            { "name": "city", "type": "string", "queryable": true }
           ,{ "name": "strt", "type": "string", "queryable": true }
         ]
      }
      ,{ 
        "name": "books", "type": "array",
        "item": {
           "type": "struct"
          ,"struct-name": "Book"
          ,"attributes": [
               { "name": "title", "type": "string", "queryable": true }
              ,{ "name": "isbn", "type": "string", "queryable": true }
              ,{ "name": "posts", "type": "array", "item": { "type": "string" }, "queryable": true}
           ]
        }
      }
   ] 
}`

func TestGeneration(t *testing.T) {

	r := bytes.NewReader([]byte(example))
	schema, e := schema.ReadCollectionDefinition(system.GetLogger(), r)
	if e != nil {
		t.Error(e)
	}

	genCfg := config.Config{TargetDirectory: ".", ResourceDirectory: "../..", FormatCode: true}
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

	r := bytes.NewReader([]byte(example))
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
