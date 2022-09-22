package example1

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FilterGoInfo() string {
	i := fmt.Sprintf("tpm_morphia query filter support generated for %s package on %s", "author", time.Now().String())
	return i
}

//----- oId of type object-id
//----- oId - object-id -  [oId]

// AndOIdEqTo No Remarks
func (ca *Criteria) AndOIdEqTo(oId primitive.ObjectID) *Criteria {

	if oId == primitive.NilObjectID {
		return ca
	}

	mName := fmt.Sprintf(OID)
	c := func() bson.E { return bson.E{Key: mName, Value: oId} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndOIdIn(p []primitive.ObjectID) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(OID)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

//----- firstName of type string
//----- firstName - string -  [firstName]

// AndFirstNameEqTo No Remarks
func (ca *Criteria) AndFirstNameEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(FIRSTNAME)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

// AndFirstNameIsNullOrUnset No Remarks
func (ca *Criteria) AndFirstNameIsNullOrUnset() *Criteria {

	mName := fmt.Sprintf(FIRSTNAME)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndFirstNameIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(FIRSTNAME)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

//----- lastName of type string
//----- lastName - string -  [lastName]

// AndLastNameEqTo No Remarks
func (ca *Criteria) AndLastNameEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(LASTNAME)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

// AndLastNameIsNullOrUnset No Remarks
func (ca *Criteria) AndLastNameIsNullOrUnset() *Criteria {

	mName := fmt.Sprintf(LASTNAME)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndLastNameIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(LASTNAME)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

//----- age of type int
//----- age - int -  [age]

// AndAgeEqTo No Remarks
func (ca *Criteria) AndAgeEqTo(p int, nullValue ...int) *Criteria {

	if len(nullValue) > 0 && p == nullValue[0] {
		return ca
	}

	mName := fmt.Sprintf(AGE)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndAgeGt(p int, nullValue ...int) *Criteria {

	if len(nullValue) > 0 && p == nullValue[0] {
		return ca
	}

	mName := fmt.Sprintf(AGE)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$gt", p}}} }
	*ca = append(*ca, c)
	return ca
}

//----- city of type string
//----- city - string -  [address.city]

// AndAddressCityEqTo No Remarks
func (ca *Criteria) AndAddressCityEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(ADDRESS_CITY)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

// AndAddressCityIsNullOrUnset No Remarks
func (ca *Criteria) AndAddressCityIsNullOrUnset() *Criteria {

	mName := fmt.Sprintf(ADDRESS_CITY)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndAddressCityIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(ADDRESS_CITY)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

//----- street of type string
//----- street - string -  [address.street]

// AndAddressStreetEqTo No Remarks
func (ca *Criteria) AndAddressStreetEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(ADDRESS_STREET)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

// AndAddressStreetIsNullOrUnset No Remarks
func (ca *Criteria) AndAddressStreetIsNullOrUnset() *Criteria {

	mName := fmt.Sprintf(ADDRESS_STREET)
	c := func() bson.E { return bson.E{Key: mName, Value: nil} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndAddressStreetIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(ADDRESS_STREET)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}
