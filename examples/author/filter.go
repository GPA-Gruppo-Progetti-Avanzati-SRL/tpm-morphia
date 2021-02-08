package author

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func FilterGoInfo() string {
	i := fmt.Sprintf("tpm_morphia query filter support generated for %s package on %s", "author", time.Now().String())
	return i
}

type Criterion func() bson.E
type Criteria []Criterion

type Filter struct {
	listOfCriteria []Criteria
}

func (f *Filter) or() *Criteria {
	ca := make(Criteria, 0)

	/*
		 * Note: the array gets allocated to avoid some nasty problem related to the fact that the underlying array can be moved
		 * when needs to make room. In this case if somebody write something like this:

			f = Filter{}
			ca := f.or().lastNameEqTo("Smith").firstNameEqTo("Colin")
			f.or().firstNameEqTo("John")
			ca.books().titleEq("My Best Seller")

		 * and saves the variable ca to be reused what happens is that the second f.or() causes the listOfCriteria
		 * array to be moved and that slice is dead as to speak.
		 * Anyway, best option would be not to interleave the conditions between or() calls.
	*/
	if len(f.listOfCriteria) == 0 {
		f.listOfCriteria = make([]Criteria, 0, 5)
	}

	// Should check top criteria if empty...
	f.listOfCriteria = append(f.listOfCriteria, ca)
	return &f.listOfCriteria[len(f.listOfCriteria)-1]
}

var emptyFilter = bson.D{}

func (f *Filter) build() bson.D {

	if len(f.listOfCriteria) == 0 {
		return emptyFilter
	}

	docA := bson.A{}
	for _, cas := range f.listOfCriteria {
		doc := bson.D{}
		for _, c := range cas {
			doc = append(doc, c())
		}
		docA = append(docA, doc)
	}

	if len(docA) == 1 {
		return docA[0].(bson.D)
	}

	return bson.D{{"$or", docA}}
}

//----- firstName of type string
func (ca *Criteria) firstNameEqTo(s string) *Criteria {
	c := func() bson.E { return bson.E{Key: FIRSTNAME, Value: s} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) firstNameIn(s []string) *Criteria {
	c := func() bson.E { return bson.E{Key: FIRSTNAME, Value: bson.D{{"$in", s}}} }
	*ca = append(*ca, c)
	return ca
}

//----- lastName of type string
func (ca *Criteria) lastNameEqTo(s string) *Criteria {
	c := func() bson.E { return bson.E{Key: LASTNAME, Value: s} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) lastNameIn(s []string) *Criteria {
	c := func() bson.E { return bson.E{Key: LASTNAME, Value: bson.D{{"$in", s}}} }
	*ca = append(*ca, c)
	return ca
}

//----- age of type int
func (ca *Criteria) ageEqTo(i int) *Criteria {
	c := func() bson.E { return bson.E{Key: AGE, Value: i} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) ageGt(a int) *Criteria {
	c := func() bson.E { return bson.E{Key: AGE, Value: bson.D{{"$gt", a}}} }
	*ca = append(*ca, c)
	return ca
}

// address struct: Address
type AddressCriteria struct {
	criteria *Criteria
}

func (ca *Criteria) address() AddressCriteria {
	return AddressCriteria{criteria: ca}
}

// address.city

//----- city of type string
func (sca AddressCriteria) cityEqTo(s string) AddressCriteria {
	c := func() bson.E { return bson.E{Key: ADDRESS_CITY, Value: s} }
	*sca.criteria = append(*sca.criteria, c)
	return sca
}

func (sca AddressCriteria) cityIn(s []string) AddressCriteria {
	c := func() bson.E { return bson.E{Key: ADDRESS_CITY, Value: bson.D{{"$in", s}}} }
	*sca.criteria = append(*sca.criteria, c)
	return sca
}

// address.strt

//----- strt of type string
func (sca AddressCriteria) strtEqTo(s string) AddressCriteria {
	c := func() bson.E { return bson.E{Key: ADDRESS_STRT, Value: s} }
	*sca.criteria = append(*sca.criteria, c)
	return sca
}

func (sca AddressCriteria) strtIn(s []string) AddressCriteria {
	c := func() bson.E { return bson.E{Key: ADDRESS_STRT, Value: bson.D{{"$in", s}}} }
	*sca.criteria = append(*sca.criteria, c)
	return sca
}

// [] struct: Book
type BookCriteria struct {
	criteria *Criteria
}

func (ca *Criteria) books() BookCriteria {
	return BookCriteria{criteria: ca}
}

// books.[].title

//----- title of type string
func (sca BookCriteria) titleEqTo(s string) BookCriteria {
	c := func() bson.E { return bson.E{Key: BOOKS_TITLE, Value: s} }
	*sca.criteria = append(*sca.criteria, c)
	return sca
}

func (sca BookCriteria) titleIn(s []string) BookCriteria {
	c := func() bson.E { return bson.E{Key: BOOKS_TITLE, Value: bson.D{{"$in", s}}} }
	*sca.criteria = append(*sca.criteria, c)
	return sca
}

// books.[].isbn

//----- isbn of type string
func (sca BookCriteria) isbnEqTo(s string) BookCriteria {
	c := func() bson.E { return bson.E{Key: BOOKS_ISBN, Value: s} }
	*sca.criteria = append(*sca.criteria, c)
	return sca
}

func (sca BookCriteria) isbnIn(s []string) BookCriteria {
	c := func() bson.E { return bson.E{Key: BOOKS_ISBN, Value: bson.D{{"$in", s}}} }
	*sca.criteria = append(*sca.criteria, c)
	return sca
}

// books.[].posts

//----- posts of type array
func (sca BookCriteria) postsEqTo(s string) BookCriteria {
	c := func() bson.E { return bson.E{Key: BOOKS_POSTS, Value: s} }
	*sca.criteria = append(*sca.criteria, c)
	return sca
}

func (sca BookCriteria) postsIn(s []string) BookCriteria {
	c := func() bson.E { return bson.E{Key: BOOKS_POSTS, Value: bson.D{{"$in", s}}} }
	*sca.criteria = append(*sca.criteria, c)
	return sca
}
