### Schema

Below is reported a slightly more elaborate sample than the one presented in the top readme. 
This example collection is used in the example2 package under [examples](../examples) 
With the help of this example the table below should be easier to interpret.

```json
{
  "name": "author",
  "properties": {
    "folder-path": "./example1",
    "struct-name": "Author"
  },
  "attributes": [
    { "name": "oId", "type": "object-id", "tags": [ "json", "-",  "bson", "_id" ], "queryable": true },
    { "name": "firstName", "type": "string", "tags": [ "json", "fn", "bson", "fn" ], "queryable": true },
    { "name": "lastName", "type": "string", "tags": [ "json", "ln", "bson", "ln" ], "queryable": true },
    { "name": "age", "type": "int", "queryable": true },
    {
      "name": "address", "type": "struct", 
      "tags": [ "json", "addr", "bson", "addr"  ], 
      "struct-name": "Address",
      "attributes": [
        { "name": "city", "type": "string", "queryable": true },
        { "name": "strt", "type": "string", "queryable": true }
      ]
    },
    { "name": "shipAddress", "type": "ref-struct", "struct-ref": { "struct-name": "Address" }, "tags": ["json", "shipaddr", "bson", "shipaddr" ], "queryable": true },
    { "name": "books", "type": "array", 
      "item": { "type": "struct", "struct-name": "Book", "attributes": [
          { "name": "title", "type": "string", "queryable": true },
          { "name": "isbn",  "type": "string", "queryable": true },
          { "name": "coAuthors", "type": "array", "item": { "type": "string" }, "queryable": true }
      ]}
    },
    {
      "name": "businessRels", "type": "map",
      "item": { "type": "struct", "struct-name": "BusinessRel", "attributes": [
          { "name": "publisherId", "type": "string", "queryable": true },
          { "name": "publisherName", "type": "string", "queryable": true },
          { "name": "contracts", "type": "map", 
              "item": {"type": "struct", "struct-name": "Contract", "attributes": [
                { "name": "contractId", "type": "string", "queryable": true },
                { "name": "contractDescr", "type": "string", "queryable": true }
              ]}
          }
        ]
      }
    }
  ]
}
```

Basically a collection is defined by a name, some properties and a list of attributes.

#### Properties
Properties customize the generation process.

| Name        | Note                                                                                                                                                          |
| ----------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| folder-path | Defines the package of the artefacts and the folder under which artifacts are created. The exact directory depends on the -out-dir param used on the cmd line |
| prefix      | If specified the collection specific files are named with the prefix before the regular name. This can help when for some reason two collections have to be generated in the same folder |
| packageName | Deprecated. Use folder-path |
| struct-name | Is the name of top struct that is generated. For the naming the normal go-way is applied. In general capital case... |

#### Attributes
Attributes types can be of

- simple types: strings, int, date (the other types are yet to be implemented)
- structs types: struct and ref-struct (address and shipAddress)
- array type (see books in the definition above)
- map type (see businessRels and contracts in the definition above)

An attribute has

- a name
- a type 
- an item definition if it's a map or an array; the item definition can be of any of the supported types 
  (note: array of array is unexplored from a behaviour perspective, so you should assume is not supported).
- a list of attributes for property of type struct and collection types where the item is a struct
- a reference to an elsewhere defined struct for properties of types ref-struct  
- an optional set of tags 
- optional attributes to control the generation output.
  
| Name        | Note                                                                                                                                                          |
| ----------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| name        | The name of the attribute. First letter gets translated to uppercase in the generated code |
| type        | The type of attribute: string, int, date, object-id, struct, ref-struct, array, map (...to be completed)       |
| tags        | An array with an even number of strings specifying the tag (index 0, 2, ...) and the value of the tag (index 1, 3, ...). If not specified json and bson tags get the value of attribute name and the omitempty flag |
| queryable   | This flag enables the generation of filter methods for the field. If false no method is generated in order to keep the artifact free of unused code. This flag applies only to simple types. In the future it will be moved to the options propperty |
| options     | CSV list of options: at the moment the list of the available values consist of one item: cust-upd-handling used for a specific use case.... (TODO: write a section on this) |
| struct-ref | Used for ref-struct attributes and allow to specify the name of the referenced struct and if the struct is external or not |
| struct-ref.struct-name | Name of the referenced struct |
| struct-ref.is-external | Declares if the struct is defined in the current collection or is defined elsewhere. |
| struct-ref.package     | full name of the package where the referenced struct is declared: use to generate the proper import and references |
| item                   | Declares the typo of the items in an array or a map. The declaration is similar to the other declarations |

The definition above will generate the following structures in a file named [model.go] (../examples/example2) (note: zeroable interface are not showed in here)

```go
type Author struct {
  OId          primitive.ObjectID     `json:"-" bson:"_id,omitempty"`
  FirstName    string                 `json:"fn,omitempty" bson:"fn,omitempty"`
  LastName     string                 `json:"ln,omitempty" bson:"ln,omitempty"`
  Age          int32                  `json:"age,omitempty" bson:"age,omitempty"`
  Address      Address                `json:"addr,omitempty" bson:"addr,omitempty"`
  ShipAddress  Address                `json:"shipaddr,omitempty" bson:"shipaddr,omitempty"`
  Books        []Book                 `json:"books,omitempty" bson:"books,omitempty"`
  BusinessRels map[string]BusinessRel `json:"businessRels,omitempty" bson:"businessRels,omitempty"`
}

type Address struct {
  City string `json:"city,omitempty" bson:"city,omitempty"`
  Strt string `json:"strt,omitempty" bson:"strt,omitempty"`
}

type Book struct {
  Title     string   `json:"title,omitempty" bson:"title,omitempty"`
  Isbn      string   `json:"isbn,omitempty" bson:"isbn,omitempty"`
  CoAuthors []string `json:"coAuthors,omitempty" bson:"coAuthors,omitempty"`
}

type BusinessRel struct {
  PublisherId   string              `json:"publisherId,omitempty" bson:"publisherId,omitempty"`
  PublisherName string              `json:"publisherName,omitempty" bson:"publisherName,omitempty"`
  Contracts     map[string]Contract `json:"contracts,omitempty" bson:"contracts,omitempty"`
}

type Contract struct {
  ContractId    string `json:"contractId,omitempty" bson:"contractId,omitempty"`
  ContractDescr string `json:"contractDescr,omitempty" bson:"contractDescr,omitempty"`
}
```






