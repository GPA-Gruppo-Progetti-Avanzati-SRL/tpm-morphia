package author

/*
type Criterion func() bson.E
type Criteria []Criterion
*/

/*
type AddrCriteria struct {
	criteria *Criteria
}


func (ca *Criteria) addr() AddrCriteria {
	return AddrCriteria{ criteria: ca}
}


func (subca AddrCriteria) cityEq(a string) AddrCriteria {
	c := func() bson.E { return bson.E { Key: ADDRESS_CITY, Value: a } }
	*subca.criteria = append(*subca.criteria, c)
	return subca
}

func (subca BookCriteria) titleEq(a string, titNdx ...interface{}) BookCriteria {
	c := func() bson.E {
		var d bson.E
		switch len(titNdx) {
		case 0:
			d = bson.E { Key: BOOKS_TITLE, Value: a }
		case 1:
			d = bson.E { Key: fmt.Sprintf(BOOKS_i_TITLE, titNdx...), Value: a }
		default:
			panic("Unsupported....")
		}
		return d
	}
	*subca.criteria = append(*subca.criteria, c)
	return subca
}

func (subca AddrCriteria) strtEq(a string) AddrCriteria {
	c := func() bson.E { return bson.E { Key: ADDRESS_STRT, Value: a } }
	*subca.criteria = append(*subca.criteria, c)
	return subca
}

*/
/*
The syntax:
	f := Filter{}
	f.or().firstNameEqTo("John")
	f.or().firstNameEqTo("Colin")

an expression like: x or y
or in term of mongo query something like:

db.authors.find({ $or : [{ "fn": "Colin" }, { "fn": "John" }] })

that is:

{"$or", bson.A{bson.D{{ "fn", "John"}}, bson.D{ {"fn", "Colin"}}}

to express something like: z and (x or y)
db.authors.find({ "ln" : "Smith", $or : [{ "fn": "Colin" }, { "fn": "John" }] })

we might want to duplicate the condition:

	f.or().firstNameEqTo("John").lastNameEqTo("Smith")
	f.or().firstNameEqTo("Colin").lastNameEqTo("Smith")

That is (z and x) or (z and y)
As a matter of fact creating:

db.authors.find({ $or : [ { $and : [{ "fn": "Colin" }, { "ln" : "Smith" }] }, { $and : [{ "fn": "John" }, { "ln" : "Smith" }]}] })

A different approach could be to accept some form of composition of filters...
but is by far more convoluted. The proposed case can be solved if we consider the in operator that is something like an or between
simple values as to speak.

	f := Filter{}
    f.and(f1, f2, ... fn)

    f1.or().lastNameEqTo("Smith")
	f2.or().firstNameEqTo("John")
	f2.or().firstNameEqTo("Colin")

we might also want to do:

filterD := bson.D{
{LASTNAME, "Smith"},
{"$or", bson.A{bson.D{{ "fn", "John"}}, bson.D{ {"fn", "Colin"}}} },
}
*/
