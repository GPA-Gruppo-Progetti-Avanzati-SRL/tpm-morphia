## Generation Details

Generated artifacts include:

- model.go: contains the types and the implementation of the Zeroable interface, it also defines constants for each of the fields in the collection
- filter.go: contains boilerplate code and types to support filter methods. The same code is in every generated group of artifacts, sort of constant.
  It is not factored in a common package because is fairly limited and by including in the package can be extended by code in the same package.
- filter-methods.go: methods specific for the collection for filtering
- update.go: contains boilerplate code and types to support update methods. The same code is in every generated group of artifacts, sort of constant.
  It is not factored in a common package because is fairly limited and by including in the package can be extended by code in the same package.
- update-methods.go: methods specific for the collection for updating

Let's briefly see the structure of this generated stuff.

---
### model.go
The file on top defines the constants for the properties of the collection. From example2 we can have something similar 
(the list shown only a small portion of the constants of the example2):

```go
const (
  OID                                      = "_id"
  FIRSTNAME                                = "fn"
  ADDRESS                                  = "addr"
  ADDRESS_CITY                             = "addr.city"
  SHIPADDRESS_CITY                         = "shipaddr.city"
  BOOKS_I_TITLE                            = "books.%d.title"
  BOOKS_I_COAUTHORS_J                      = "books.%d.coAuthors.%d"
  BUSINESSRELS                             = "businessRels"
  BUSINESSRELS_S                           = "businessRels.%s"
  BUSINESSRELS_S_PUBLISHERNAME             = "businessRels.%s.publisherName"
  BUSINESSRELS_S_CONTRACTS                 = "businessRels.%s.contracts"
  BUSINESSRELS_S_CONTRACTS_T               = "businessRels.%s.contracts.%s"
)
```

- The constants are named after the go names of the attributes and mapped to the persisted name (i.e. FIRSTANAME is mapped to "fn" that is the mongoDb name)
- Structs and reference to structs (in case the reference is internal) use the dotted notation and consider the various path to the attributes
  (i.e. both ADDRESS_CITY and SHIPADDRESS_CITY are generated). In case the reference is external this is not possible and some manual work is required.
- Arrays define place constants with and without index placeholders. Placeholders are mapped to a fmt.Sprintf notation with the %d in the position of the required index.
  The name of the constants put the chars 'I', 'J', ... in the place holder positions.
- Maps do pretty much the same but with the notable difference that place holders are strings. The constants in this case use 'S', 'T', ... as
  specifiers in the name
  
After the const section the file contains a list of the generated structs with the IsZero method for the ZeroableInterface.

---
### filter-methods.go
The generated model allows the easily creation of filters of the following structure:

> (predicate1 and predicate2 and ...) or (predicate3 and predicate4 and ...)

that pretty much cover a reasonable number of use cases. For more specifice queries not covered by this model is always possible to write some custom
code as described below.

Methods are created for each simple type in the hierarchy that has been declared queryable (just not to generate pile of unused stuff you have no intention to query on)
Different methods with different operators are created for different types.
At the momento the support of types and methods is limited but is a matter of simple extension to include more. 
Something I have not the time to do.

| Data type | Note | 
| --------- | ------- | 
| string    | Filter for equality and in clause (i.e. func (ca *Criteria) AndAddressCityIn(p []string) *Criteria) |
| int       | Filter for equality and greater than (i.e. func (ca *Criteria) AndAgeGt(p int, nullValue ...int) *Criteria) |
| map       | Filter dependent on the simple type of the property but indexed to address a specific key of the map (i.e. func (ca *Criteria) AndBusinessRelsSPublisherIdEqTo(keyS string, p string) *Criteria) |

Methods to query for a specific item at a position in the array is not generated at this time and considered very low priority.
Types have to be extended.

Some methods have a nullValue optional list of params. I didn't find a better code solution to indicate the value that should be considered as a null value 
(for strings the empty string can be considered the null value but for int this might be tricky). The use of this parameter is clear if you look at a generated method.
Basically at the beginning there is a check to see if the argument should be considered null. In that case the filter is not applied. This should 
easier form base type of queries where you can have optional fields and is tedious to write an if to check for nulliness.
(note: the code for validation if required should handle anomalies well before getting to the actual filter method)

```go
func (ca *Criteria) AndAgeGt(p int, nullValue ...int) *Criteria {
  if len(nullValue) > 0 && p == nullValue[0] {
     return ca
  }

  mName := fmt.Sprintf(AGE)
  c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$gt", p}}} }
  *ca = append(*ca, c)
  return ca
}
```
---
### update-methods.go
The file contains the update methods for different attributes. The list below describe a list of sample signatures from example2

- func (ud *UpdateDocument) SetFirstName(p string) *UpdateDocument
- func (ud *UpdateDocument) UnsetAge() *UpdateDocument
- func (ud *UpdateDocument) IncAge(p int32) *UpdateDocument 
- func (ud *UpdateDocument) SetAddress(p Address) *UpdateDocument
- func (ud *UpdateDocument) SetBooks(p []Book) *UpdateDocument
- func (ud *UpdateDocument) SetBooksI(ndxI int, p Book) *UpdateDocument
- func (ud *UpdateDocument) SetBusinessRelsS(keyS string, p BusinessRel) *UpdateDocument
- func (ud *UpdateDocument) SetBusinessRelsSPublisherId(keyS string, p string) *UpdateDocument

Basically the $set, $unset, $inc operators are supported together with the indexed form for arrays and maps.

---
### extension files
The use of extension files is pretty obvious. It's a matter of manually creating a file in the same package and add methods that are not generated or that require a more elaborate
structure to address different patterns. I tend to name thos files like filter-methods-ext.go or update-methods-ext.go. They are not overwritten by generation
and can contain the code you deem useful. I have been using these files for external ref-struct by copy-ing and pasting generated in the common collection.

