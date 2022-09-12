package schema

import (
	"os"
	"testing"
)

var example = `{
   "name": "author",
   "properties": {
       "folder-path": "./author"
      ,"prefix": "author"
      ,"struct-name": "Author"
      ,"morphia-pkg": "../morphia"
   },
   "attributes": [
       { "name": "firstName", "type": "string", "tags": [ "json", "fn", "bson", "fn" ], "queryable": true }
      ,{ "name": "lastName", "type": "string", "tags": [ "json", "ln", "bson", "ln" ], "queryable": true }
      ,{ "name": "age", "type": "int" }
      ,{ "name": "lastUpdate", "type": "date" }
      ,{ "name": "address",  "type": "struct", "tags": [ "json", "addr", "bson", "addr" ]
         ,"struct-name": "Address"
         ,"attributes": [
            { "name": "city", "type": "string", "options": "o1,o2" }
           ,{ "name": "strt", "type": "string" }
         ]
      }
      ,{ "name": "amazon",  "type": "ref-struct", "tags": [ "json", "amazon", "bson", "amz" ], "struct-ref": { "struct-name": "Address", "is-external": true, "package": "zucca/pkg" }}
      ,{ 
        "name": "books", "type": "array",
        "item": {
           "type": "struct"
          ,"struct-name": "Book"
          ,"attributes": [
               { "name": "title", "type": "string" }
              ,{ "name": "isbn", "type": "string" }
              ,{ "name": "posts", "type": "array", "item": { "type": "string" }}
           ]
        }
      }
      ,{ 
        "name": "props", "type": "map",
        "item": {
           "type": "struct"
          ,"struct-name": "Property"
          ,"attributes": [
               { "name": "ptitle", "type": "string" }
              ,{ "name": "pisbn", "type": "string" } 
           ]
        }
      } 
   ] 
}`

var tests_datatypes = "../tests/datatypes-tpmm.json"
var tests_cliente = "../tests/cliente-tpmm.json"

func TestParse(t *testing.T) {

	f, err := os.Open(tests_cliente)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	_, e := ReadCollectionDefinition(f)
	if e != nil {
		t.Error(e)
	}

}
