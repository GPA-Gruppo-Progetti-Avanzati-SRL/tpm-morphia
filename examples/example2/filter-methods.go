package example2

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
//----- city - string -  [address.city shipAddress.city]

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

func (ca *Criteria) AndAddressCityIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(ADDRESS_CITY)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

// AndShipAddressCityEqTo No Remarks
func (ca *Criteria) AndShipAddressCityEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(SHIPADDRESS_CITY)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndShipAddressCityIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(SHIPADDRESS_CITY)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

//----- strt of type string
//----- strt - string -  [address.strt shipAddress.strt]

// AndAddressStrtEqTo No Remarks
func (ca *Criteria) AndAddressStrtEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(ADDRESS_STRT)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndAddressStrtIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(ADDRESS_STRT)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

// AndShipAddressStrtEqTo No Remarks
func (ca *Criteria) AndShipAddressStrtEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(SHIPADDRESS_STRT)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndShipAddressStrtIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(SHIPADDRESS_STRT)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

//----- title of type string
//----- title - string -  [books.[].title books.title]

// AndBooksTitleEqTo No Remarks
func (ca *Criteria) AndBooksTitleEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(BOOKS_TITLE)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndBooksTitleIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(BOOKS_TITLE)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

//----- isbn of type string
//----- isbn - string -  [books.[].isbn books.isbn]

// AndBooksIsbnEqTo No Remarks
func (ca *Criteria) AndBooksIsbnEqTo(p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(BOOKS_ISBN)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndBooksIsbnIn(p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(BOOKS_ISBN)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

//----- publisherId of type string
//----- publisherId - string -  [businessRels.%s.publisherId]

// AndBusinessRelsSPublisherIdEqTo No Remarks
func (ca *Criteria) AndBusinessRelsSPublisherIdEqTo(keyS string, p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(BUSINESSRELS_S_PUBLISHERID, keyS)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndBusinessRelsSPublisherIdIn(keyS string, p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(BUSINESSRELS_S_PUBLISHERID, keyS)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

//----- publisherName of type string
//----- publisherName - string -  [businessRels.%s.publisherName]

// AndBusinessRelsSPublisherNameEqTo No Remarks
func (ca *Criteria) AndBusinessRelsSPublisherNameEqTo(keyS string, p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(BUSINESSRELS_S_PUBLISHERNAME, keyS)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndBusinessRelsSPublisherNameIn(keyS string, p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(BUSINESSRELS_S_PUBLISHERNAME, keyS)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

//----- contractId of type string
//----- contractId - string -  [businessRels.%s.contracts.%s.contractId]

// AndBusinessRelsSContractsTContractIdEqTo No Remarks
func (ca *Criteria) AndBusinessRelsSContractsTContractIdEqTo(keyS string, keyT string, p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T_CONTRACTID, keyS, keyT)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndBusinessRelsSContractsTContractIdIn(keyS string, keyT string, p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T_CONTRACTID, keyS, keyT)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

//----- contractDescr of type string
//----- contractDescr - string -  [businessRels.%s.contracts.%s.contractDescr]

// AndBusinessRelsSContractsTContractDescrEqTo No Remarks
func (ca *Criteria) AndBusinessRelsSContractsTContractDescrEqTo(keyS string, keyT string, p string) *Criteria {

	if p == "" {
		return ca
	}

	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T_CONTRACTDESCR, keyS, keyT)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndBusinessRelsSContractsTContractDescrIn(keyS string, keyT string, p []string) *Criteria {

	if len(p) == 0 {
		return ca
	}

	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T_CONTRACTDESCR, keyS, keyT)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}
