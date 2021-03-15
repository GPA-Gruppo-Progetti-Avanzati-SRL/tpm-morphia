## TPM-Morphia

tpm-morphia is simple stuff. It is simple in the name and in substance. 
It's a generation tool to support easier coding in golang with mongo.
The following notes try to clarify the way it works and the rationale. May be out there you can find tons of code that pretty much overlap.
I didn't stumble across anything like that. May be because there is no need for doing that or simply because that is not the golang way of
doing it.

---
### Toy example
Let's use a fictitious example to introduce the stuff.
Suppose you have a collection with the following profile: you have few fields and and address sub-structure.

```json
{
  "_id": {
    "$oid": "604f32e7ba935d557df60b9a"
  },
  "addr": {
    "city": "Atlanta",
    "street": "Marietta St."
  },
  "age": 30,
  "fn": "John",
  "ln": "Smith"
}
```

In mongo style you might want to code something like this. The observant reader sure have noticed that the structure Address is embedded
in the Author struct and not referenced through a pointer. This of course has some subtle implications in the way stuff gets persisted in the db.

```go
type Author struct {
  OId       primitive.ObjectID `json:"-" bson:"_id,omitempty"`
  FirstName string             `json:"fn,omitempty" bson:"fn,omitempty"`
  LastName  string             `json:"ln,omitempty" bson:"ln,omitempty"`
  Age       int32              `json:"age,omitempty" bson:"age,omitempty"`
  Address   Address            `json:"addr,omitempty" bson:"addr,omitempty"`
}

type Address struct {
  City   string `json:"city,omitempty" bson:"city,omitempty"`
  Street string `json:"strt,omitempty" bson:"strt,omitempty"`
}
```

The obvious operation is to create a few records and add to a collection. In our case the code could be something like (from example0):

```go
a := example0.Author{
  FirstName: fn,
  LastName:  ln,
  Age:       30,
  Address:   example0.Address{City: city, Street: strt},
}

r, err := aCollection.InsertOne(ctx, a)
```

The resulting object in the collection is pretty much what we expect to be, with just one caveat: if we provide an empty city and street the 
addr property is set to an empty structure. This migh be counter intuitive considering that we explicitly specified the omitempty
option in struct bson tag. Actually the omitempty to be honored in this context requires that the structure implements the Zeroable interface.
In that case the empty struct gets not persisted and the property is unset. This is not required in the Address was a pointer to structure since 
the marshaller in that case esily understands what to do with it. May be you decide to implement that interface in a way similar to what is reported below.

```go
func (s Author) IsZero() bool {
   if s.OId != primitive.NilObjectID {
      return false
   }
   if s.FirstName != "" {
   	  return false
   }
   if s.LastName != "" {
   	  return false
   }
   if s.Age != 0 {
   	  return false
   }
   if !s.Address.IsZero() {
   	  return false
   }
   return true
}

func (s Address) IsZero() bool {
   if s.City != "" {
      return false
   }
   
   if s.Street != "" {
      return false
   }
   return true
}
```

Ok, so far so good. Now it's time to query this collection and update a record just to show the syntax... A typical 
snippet is reported below. 

```go
filter := bson.D{
   {"$or", bson.A{bson.D{{"fn", fn}}, bson.D{{"ln", ln}, {"addr.city", cy}}}},
}

cur, err := aCollection.Find(ctx, filter)
```

In here we look for records verifyng this condition in pseudocode:

> firstName == fn or (lastName == ln and address.city == cy)

This snippet, at least to me, has a few problems:

- too many strings which you can mess up with a simple typo, 
- too many curly braces, with errors that might pop-up in situaztions a little more complicated.
- the need to specify the mongo field nammes as opposed to the object field names: this might suggest to keep them in synch but it is not always possible.
- the operator sort of prefix notation is somewhat innatural.

What if we want to do an update? That requires two documents: one to filter the records, one to do the update.

```go
opts := options.Update().SetUpsert(true)

filter := bson.D{
   {"$and", bson.A{bson.D{{"fn", fn}}, bson.D{{"ln", ln}}}},
}

updateDoc := bson.D{{"$set", bson.D{{"addr.city", cy}}}}
if ur, err := aCollection.UpdateOne(ctx, filter, updateDoc, opts); err != nil {
   return err
} else {
   _ = level.Info(logger).Log("msg", "update result", "upsertedCound", ur.UpsertedCount, "modifiedCount", ur.ModifiedCount)
}
```

On top of previous problems, a different one gets uin the way. There is the need to aggregate the operations by operator or 
to repeat an array of bson documents each with the required operator. In this case, at least in a different form 
it pop-up the issue of the empty structures in case you specifically do a $set with an empty one. 

### Toy example reloaded
In this simple case would be good if we could express the find and update operations in a more expressive way.

```go
f := example1.Filter{}
f.Or().AndFirstNameEqTo(fn)
f.Or().AndLastNameEqTo(ln).AndAddressCityEqTo(cy)

cur, err := aCollection.Find(ctx, f.Build())
```

If we could do that we could get rid of the strings, have support from the IDE with auto-completion and express the conditions in terms of the 
struct properties and not in terms of the actual bson names.

In a similar way, it could be good to rephrase the update statement

```go
opts := options.Update().SetUpsert(true)

f := example1.Filter{}
f.Or().AndFirstNameEqTo(fn).AndLastNameEqTo(ln)

updateDoc := example1.UpdateDocument{}
updateDoc.SetAddressCity(cy)

if ur, err := aCollection.UpdateOne(ctx, f.Build(), updateDoc.Build(), opts); err != nil {
   return err
} else {
   _ = level.Info(logger).Log("msg", "update result", "upsertedCound", ur.UpsertedCount, "modifiedCount", ur.ModifiedCount)
}
```

---
### TPM-Morphia
The idea of tpm-morphia is to generate some code to allow the writing of stuff like that. Basically the tool starts with the definition of the collection of interest
and tries to generate the types and the methods to support this coding. 

Collection definition is a simple Json file where properties of the collection have to be specificed together with a few metadata to taylor the generation process.
In the case of our toy example the collection definition (see example1-tpmm.json in the examples directory) it would look like something like that.

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
        { "name": "street", "type": "string", "queryable": true }
      ]
    }
  ]
}
```

Out of this definition the generator creates (or use if it exists) a folder named example1 (relative to the file definition path) with a number of files in there:

- model.go: contains the types and the implementation of the Zeroable interface, it also defines constants for each of the fields in the collection
- filter.go: contains boilerplate code and types to support filter methods
- filter-methods.go: methods specific for the collection for filtering
- update.go: contains boilerplate code and types to support update methods
- update-methods.go: methods specific for the collection for updating

Below, just to give the feeling of what gets generated in the filter-methods of the example discussed so far:

func (ca *Criteria) AndFirstNameEqTo(p string) *Criteria {

```go
func (ca *Criteria) AndFirstNameEqTo(p string) *Criteria {  
  if p == "" {
    return ca  
  }

  mName := fmt.Sprintf(FIRSTNAME)
  c := func () bson.E { return bson.E{Key: mName, Value: p} }
  *ca = append(*ca, c)
  return ca
}	
```

---
### Usage
Next the option that can be specified to trigger the generation.

| Option                       | Note                                                                     | Default  |
| ---------------------------- | ------------------------------------------------------------------------ | -------- |
| -out-dir                     | Mount point of generated content (see. folder-path property)             |          |
| -collection-def-file         | Collection definition file                                               |          |
| -collection-def-scan-path    | Collection scan path (it scans for names like *-tpmm.json)               |          |
| -format-code                 | Boolean value to format the generated code                               | true     |
| -tmpl-ver                    | Version of templates (v2)                                                | v2       |

The following command invokes the main and searches the schema files in ./model and create artifacts in folders under ./model

```go
tpm-morphia -collection-def-scan-path ./model -out-dir ./model
```

---
### Schema definition
For the syntax of the schema files please refer to [collection schema](./schema)

### Generation details
For some details about the generated stuff please see [generation details](./gen/mongodb)


---
### Examples
For examples look into examples directory [examples](./examples)

- examples: contains a number of schema definitions and a test function to generate the artifacts. The git already contains a run of generation ready to be inspected.
- example0: sample code of querying and updating mongodb with vanilla API.
- example1: identical functionality of example0 but using the generated methods to carry out the required ops.