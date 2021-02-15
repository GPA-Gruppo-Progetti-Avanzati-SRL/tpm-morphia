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

// oId - object-id -  [oId]
func (ca *Criteria) AndOIdEqTo(oId primitive.ObjectID) *Criteria {
	mName := fmt.Sprintf(OID)
	c := func() bson.E { return bson.E{Key: mName, Value: oId} }
	*ca = append(*ca, c)
	return ca
}

// firstName - string -  [firstName]
func (ca *Criteria) AndFirstNameEqTo(p string) *Criteria {
	mName := fmt.Sprintf(FIRSTNAME)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndFirstNameIn(p []string) *Criteria {
	mName := fmt.Sprintf(FIRSTNAME)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

// lastName - string -  [lastName]
func (ca *Criteria) AndLastNameEqTo(p string) *Criteria {
	mName := fmt.Sprintf(LASTNAME)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndLastNameIn(p []string) *Criteria {
	mName := fmt.Sprintf(LASTNAME)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

//----- age of type int

// age - int -  [age]
func (ca *Criteria) AndAgeEqTo(p int) *Criteria {
	mName := fmt.Sprintf(AGE)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndAgeGt(p int) *Criteria {
	mName := fmt.Sprintf(AGE)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$gt", p}}} }
	*ca = append(*ca, c)
	return ca
}

// city - string -  [address.city shipAddress.city]
func (ca *Criteria) AndAddressCityEqTo(p string) *Criteria {
	mName := fmt.Sprintf(ADDRESS_CITY)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndAddressCityIn(p []string) *Criteria {
	mName := fmt.Sprintf(ADDRESS_CITY)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}
func (ca *Criteria) AndShipAddressCityEqTo(p string) *Criteria {
	mName := fmt.Sprintf(SHIPADDRESS_CITY)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndShipAddressCityIn(p []string) *Criteria {
	mName := fmt.Sprintf(SHIPADDRESS_CITY)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

// strt - string -  [address.strt shipAddress.strt]
func (ca *Criteria) AndAddressStrtEqTo(p string) *Criteria {
	mName := fmt.Sprintf(ADDRESS_STRT)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndAddressStrtIn(p []string) *Criteria {
	mName := fmt.Sprintf(ADDRESS_STRT)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}
func (ca *Criteria) AndShipAddressStrtEqTo(p string) *Criteria {
	mName := fmt.Sprintf(SHIPADDRESS_STRT)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndShipAddressStrtIn(p []string) *Criteria {
	mName := fmt.Sprintf(SHIPADDRESS_STRT)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

// title - string -  [books.[].title books.title]
func (ca *Criteria) AndBooksTitleEqTo(p string) *Criteria {
	mName := fmt.Sprintf(BOOKS_TITLE)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndBooksTitleIn(p []string) *Criteria {
	mName := fmt.Sprintf(BOOKS_TITLE)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

// isbn - string -  [books.[].isbn books.isbn]
func (ca *Criteria) AndBooksIsbnEqTo(p string) *Criteria {
	mName := fmt.Sprintf(BOOKS_ISBN)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndBooksIsbnIn(p []string) *Criteria {
	mName := fmt.Sprintf(BOOKS_ISBN)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

// publisherId - string -  [businessRels.%s.publisherId]
func (ca *Criteria) AndBusinessRelsSPublisherIdEqTo(keyS string, p string) *Criteria {
	mName := fmt.Sprintf(BUSINESSRELS_S_PUBLISHERID, keyS)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndBusinessRelsSPublisherIdIn(keyS string, p []string) *Criteria {
	mName := fmt.Sprintf(BUSINESSRELS_S_PUBLISHERID, keyS)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

// publisherName - string -  [businessRels.%s.publisherName]
func (ca *Criteria) AndBusinessRelsSPublisherNameEqTo(keyS string, p string) *Criteria {
	mName := fmt.Sprintf(BUSINESSRELS_S_PUBLISHERNAME, keyS)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndBusinessRelsSPublisherNameIn(keyS string, p []string) *Criteria {
	mName := fmt.Sprintf(BUSINESSRELS_S_PUBLISHERNAME, keyS)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

// contractId - string -  [businessRels.%s.contracts.%s.contractId]
func (ca *Criteria) AndBusinessRelsSContractsTContractIdEqTo(keyS string, keyT string, p string) *Criteria {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T_CONTRACTID, keyS, keyT)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndBusinessRelsSContractsTContractIdIn(keyS string, keyT string, p []string) *Criteria {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T_CONTRACTID, keyS, keyT)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}

// contractDescr - string -  [businessRels.%s.contracts.%s.contractDescr]
func (ca *Criteria) AndBusinessRelsSContractsTContractDescrEqTo(keyS string, keyT string, p string) *Criteria {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T_CONTRACTDESCR, keyS, keyT)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	*ca = append(*ca, c)
	return ca
}

func (ca *Criteria) AndBusinessRelsSContractsTContractDescrIn(keyS string, keyT string, p []string) *Criteria {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T_CONTRACTDESCR, keyS, keyT)
	c := func() bson.E { return bson.E{Key: mName, Value: bson.D{{"$in", p}}} }
	*ca = append(*ca, c)
	return ca
}
