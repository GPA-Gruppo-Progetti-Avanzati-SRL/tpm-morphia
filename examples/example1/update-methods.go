package example1

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateMethodsGoInfo() string {
	i := fmt.Sprintf("tpm_morphia query filter support generated for %s package on %s", "author", time.Now().String())
	return i
}

// oId - object-id -  [oId]
func (upds *Updates) SetOId(p primitive.ObjectID) *Updates {
	mName := fmt.Sprintf(OID)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

// firstName - string -  [firstName]
func (upds *Updates) SetFirstName(p string) *Updates {
	mName := fmt.Sprintf(FIRSTNAME)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

// lastName - string -  [lastName]
func (upds *Updates) SetLastName(p string) *Updates {
	mName := fmt.Sprintf(LASTNAME)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

// age - int -  [age]
func (upds *Updates) SetAge(p int32) *Updates {
	mName := fmt.Sprintf(AGE)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

// address - struct - Address [address]
func (upds *Updates) SetAddress(p Address) *Updates {
	mName := fmt.Sprintf(ADDRESS)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

// city - string -  [address.city shipAddress.city]
func (upds *Updates) SetAddressCity(p string) *Updates {
	mName := fmt.Sprintf(ADDRESS_CITY)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}
func (upds *Updates) SetShipAddressCity(p string) *Updates {
	mName := fmt.Sprintf(SHIPADDRESS_CITY)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

// strt - string -  [address.strt shipAddress.strt]
func (upds *Updates) SetAddressStrt(p string) *Updates {
	mName := fmt.Sprintf(ADDRESS_STRT)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}
func (upds *Updates) SetShipAddressStrt(p string) *Updates {
	mName := fmt.Sprintf(SHIPADDRESS_STRT)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

// shipAddress - ref-struct -  [shipAddress]
func (upds *Updates) SetShipAddress(p Address) *Updates {
	mName := fmt.Sprintf(SHIPADDRESS)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

// books - array -  [books]
func (upds *Updates) SetBooks(p []Book) *Updates {
	mName := fmt.Sprintf(BOOKS)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

// [] - struct - Book [books.[]]
func (upds *Updates) SetBooksI(ndxI int, p Book) *Updates {
	mName := fmt.Sprintf(BOOKS_I, ndxI)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

// title - string -  [books.[].title books.title]
func (upds *Updates) SetBooksITitle(ndxI int, p string) *Updates {
	mName := fmt.Sprintf(BOOKS_I_TITLE, ndxI)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

// isbn - string -  [books.[].isbn books.isbn]
func (upds *Updates) SetBooksIIsbn(ndxI int, p string) *Updates {
	mName := fmt.Sprintf(BOOKS_I_ISBN, ndxI)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

// coAuthors - array -  [books.[].coAuthors books.coAuthors]
func (upds *Updates) SetBooksICoAuthors(ndxI int, p []string) *Updates {
	mName := fmt.Sprintf(BOOKS_I_COAUTHORS, ndxI)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

// [] - string -  [books.[].coAuthors.[]]
func (upds *Updates) SetBooksICoAuthorsJ(ndxI int, ndxJ int, p string) *Updates {
	mName := fmt.Sprintf(BOOKS_I_COAUTHORS_J, ndxI, ndxJ)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

// businessRels - map -  [businessRels]
func (upds *Updates) SetBusinessRels(p map[string]BusinessRel) *Updates {
	mName := fmt.Sprintf(BUSINESSRELS)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

// %s - struct - BusinessRel [businessRels.%s]
func (upds *Updates) SetBusinessRelsS(keyS string, p BusinessRel) *Updates {
	mName := fmt.Sprintf(BUSINESSRELS_S, keyS)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

// publisherId - string -  [businessRels.%s.publisherId]
func (upds *Updates) SetBusinessRelsSPublisherId(keyS string, p string) *Updates {
	mName := fmt.Sprintf(BUSINESSRELS_S_PUBLISHERID, keyS)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

// publisherName - string -  [businessRels.%s.publisherName]
func (upds *Updates) SetBusinessRelsSPublisherName(keyS string, p string) *Updates {
	mName := fmt.Sprintf(BUSINESSRELS_S_PUBLISHERNAME, keyS)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

// contracts - map -  [businessRels.%s.contracts]
func (upds *Updates) SetBusinessRelsSContracts(keyS string, p map[string]Contract) *Updates {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS, keyS)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

// %s - struct - Contract [businessRels.%s.contracts.%s]
func (upds *Updates) SetBusinessRelsSContractsT(keyS string, keyT string, p Contract) *Updates {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T, keyS, keyT)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

// contractId - string -  [businessRels.%s.contracts.%s.contractId]
func (upds *Updates) SetBusinessRelsSContractsTContractId(keyS string, keyT string, p string) *Updates {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T_CONTRACTID, keyS, keyT)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

// contractDescr - string -  [businessRels.%s.contracts.%s.contractDescr]
func (upds *Updates) SetBusinessRelsSContractsTContractDescr(keyS string, keyT string, p string) *Updates {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T_CONTRACTDESCR, keyS, keyT)
	upds.Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return upds
}

/*
func (upds *Updates) SetAddress(a Address) *Updates {

	upds.Add(func() bson.E {
		return bson.E{ Key: ADDRESS, Value: a}
	})

    return upds
}

func (upds *Updates) SetShippingAddress(a Address) *Updates {

	upds.Add(func() bson.E {
		return bson.E{ Key: SHIPADDRESS, Value: a}
	})

	return upds
}

func (ca *Updates) SetBusinessRelsS(keyS string, p BusinessRel) *Updates {
	mName := fmt.Sprintf(BUSINESSRELS_S, keyS)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	ca.Add(c)
	return ca
}

func (ca *Updates) SetBusinessRelsSContractsT(keyS string, keyT string, p Contract) *Updates {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T, keyS, keyT)
	c := func() bson.E { return bson.E{Key: mName, Value: p} }
	ca.Add(c)
	return ca
}
*/
