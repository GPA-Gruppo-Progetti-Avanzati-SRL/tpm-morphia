package mongodb

import (
	"bytes"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/config"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schema"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system"
	"testing"
)

/*
   {
     "name": "extReference",
     "type": "ref-struct",
     "struct-ref": {
       "struct-name": "ExternalStruct",
       "is-external": true,
       "package": "lib/external/extpkg"
     },
     "queryable": false
   },
*/
var example = `{
  "name": "author",
  "properties": {
    "folder-path": "./author",
    "struct-name": "Author"
  },
  "attributes": [
    {
      "name": "oId",
      "type": "object-id",
      "tags": [
        "json",
        "-",
        "bson",
        "_id"
      ],
      "queryable": true
    },
    {
      "name": "firstName",
      "type": "string",
      "tags": [
        "json",
        "fn",
        "bson",
        "fn"
      ],
      "queryable": true
    },
    {
      "name": "lastUpdate",
      "type": "date",
      "tags": [
        "json",
        "dt",
        "bson",
        "dt"
      ],
      "queryable": true
    },
    {
      "name": "lastName",
      "type": "string",
      "tags": [
        "json",
        "ln",
        "bson",
        "ln"
      ],
      "queryable": true
    },
    {
      "name": "age",
      "type": "int",
      "queryable": true
    },
    {
      "name": "address",
      "type": "struct",
      "tags": [
        "json",
        "addr",
        "bson",
        "addr"
      ],
      "struct-name": "Address",
      "attributes": [
        {
          "name": "city",
          "type": "string",
          "queryable": true
        },
        {
          "name": "strt",
          "type": "string",
          "queryable": true
        }
      ]
    },
    {
      "name": "shipAddress",
      "type": "ref-struct",
      "tags": [
        "json",
        "shipaddr",
        "bson",
        "shipaddr"
      ],
      "struct-ref": {
        "struct-name": "Address"
      },
      "queryable": true
    },

    {
      "name": "books",
      "type": "array",
      "item": {
        "type": "struct",
        "struct-name": "Book",
        "attributes": [
          {
            "name": "title",
            "type": "string",
            "queryable": true
          },
          {
            "name": "isbn",
            "type": "string",
            "queryable": true
          },
          {
            "name": "coAuthors",
            "type": "array",
            "item": {
              "type": "string"
            },
            "queryable": true
          }
        ]
      }
    },
    {
      "name": "businessRels",
      "type": "map",
      "item": {
        "type": "struct",
        "struct-name": "BusinessRel",
        "attributes": [
          {
            "name": "publisherId",
            "type": "string",
            "queryable": true
          },
          {
            "name": "publisherName",
            "type": "string",
            "queryable": true
          },
          {
            "name": "contracts",
            "type": "map",
            "item": {
              "type": "struct",
              "struct-name": "Contract",
              "attributes": [
                {
                  "name": "contractId",
                  "type": "string",
                  "queryable": true
                },
                {
                  "name": "contractDescr",
                  "type": "string",
                  "queryable": true
                }
              ]
            }
          }
        ]
      }
    }
  ]
}`

func TestGeneration(t *testing.T) {

	r := bytes.NewReader([]byte(example))
	schema, e := schema.ReadCollectionDefinition(system.GetLogger(), r)
	if e != nil {
		t.Fatal(e)
	}

	genCfg := config.Config{TargetDirectory: ".", ResourceDirectory: "../..", FormatCode: false}
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
