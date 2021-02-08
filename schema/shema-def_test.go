package schema

import (
	"bytes"
	"fmt"
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
      ,{ "name": "age", "type": "int" }
      ,{ "name": "address",  "type": "struct", "tags": [ "json", "addr", "bson", "addr" ]
         ,"struct-name": "Address"
         ,"attributes": [
            { "name": "city", "type": "string" }
           ,{ "name": "strt", "type": "string" }
         ]
      }
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
   ] 
}`

func TestParse(t *testing.T) {

	r := bytes.NewReader([]byte(example))
	schema, e := ReadCollectionDefinition(system.GetLogger(), r)
	if e != nil {
		t.Error(e)
	}

	fmt.Println(schema)
}
