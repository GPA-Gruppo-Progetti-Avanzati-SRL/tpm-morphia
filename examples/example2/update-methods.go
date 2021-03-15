package example2

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

/*
 * Convenience method to create an Update Document from the values of the top fields of the object. The convenience is in the handling
 * the unset because if I pass an empty struct to the update it generates an empty object anyway in the db. Handling the unset eliminates
 * the issue and delete an existing value without creating an empty struct..
 */
func GetUpdateDocument(obj *Author) UpdateDocument {
	ud := UpdateDocument{}
	if obj.FirstName != "" {
		ud.SetFirstName(obj.FirstName)
	} else {
		ud.UnsetFirstName()
	}
	if obj.LastName != "" {
		ud.SetLastName(obj.LastName)
	} else {
		ud.UnsetLastName()
	}
	if obj.Age != 0 {
		ud.SetAge(obj.Age)
	} else {
		ud.UnsetAge()
	}
	if !obj.Address.IsZero() {
		ud.SetAddress(obj.Address)
	} else {
		ud.UnsetAddress()
	}
	if !obj.ShipAddress.IsZero() {
		ud.SetShipAddress(obj.ShipAddress)
	} else {
		ud.UnsetShipAddress()
	}
	if len(obj.Books) > 0 {
		ud.SetBooks(obj.Books)
	} else {
		ud.UnsetBooks()
	}
	if len(obj.BusinessRels) > 0 {
		ud.SetBusinessRels(obj.BusinessRels)
	} else {
		ud.UnsetBusinessRels()
	}

	return ud
}

// oId - object-id -  [oId]
func (ud *UpdateDocument) SetOId(p primitive.ObjectID) *UpdateDocument {
	mName := fmt.Sprintf(OID)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetOId() *UpdateDocument {
	mName := fmt.Sprintf(OID)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// firstName - string -  [firstName]
func (ud *UpdateDocument) SetFirstName(p string) *UpdateDocument {
	mName := fmt.Sprintf(FIRSTNAME)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetFirstName() *UpdateDocument {
	mName := fmt.Sprintf(FIRSTNAME)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// lastName - string -  [lastName]
func (ud *UpdateDocument) SetLastName(p string) *UpdateDocument {
	mName := fmt.Sprintf(LASTNAME)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetLastName() *UpdateDocument {
	mName := fmt.Sprintf(LASTNAME)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// age - int -  [age]
func (ud *UpdateDocument) SetAge(p int32) *UpdateDocument {
	mName := fmt.Sprintf(AGE)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetAge() *UpdateDocument {
	mName := fmt.Sprintf(AGE)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

func (ud *UpdateDocument) IncAge(p int32) *UpdateDocument {
	mName := fmt.Sprintf(AGE)
	ud.Inc().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// address - struct - Address [address]
func (ud *UpdateDocument) SetAddress(p Address) *UpdateDocument {
	mName := fmt.Sprintf(ADDRESS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetAddress() *UpdateDocument {
	mName := fmt.Sprintf(ADDRESS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// city - string -  [address.city shipAddress.city]
func (ud *UpdateDocument) SetAddressCity(p string) *UpdateDocument {
	mName := fmt.Sprintf(ADDRESS_CITY)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetAddressCity() *UpdateDocument {
	mName := fmt.Sprintf(ADDRESS_CITY)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}
func (ud *UpdateDocument) SetShipAddressCity(p string) *UpdateDocument {
	mName := fmt.Sprintf(SHIPADDRESS_CITY)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetShipAddressCity() *UpdateDocument {
	mName := fmt.Sprintf(SHIPADDRESS_CITY)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// strt - string -  [address.strt shipAddress.strt]
func (ud *UpdateDocument) SetAddressStrt(p string) *UpdateDocument {
	mName := fmt.Sprintf(ADDRESS_STRT)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetAddressStrt() *UpdateDocument {
	mName := fmt.Sprintf(ADDRESS_STRT)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}
func (ud *UpdateDocument) SetShipAddressStrt(p string) *UpdateDocument {
	mName := fmt.Sprintf(SHIPADDRESS_STRT)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetShipAddressStrt() *UpdateDocument {
	mName := fmt.Sprintf(SHIPADDRESS_STRT)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// shipAddress - ref-struct -  [shipAddress]
func (ud *UpdateDocument) SetShipAddress(p Address) *UpdateDocument {
	mName := fmt.Sprintf(SHIPADDRESS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetShipAddress() *UpdateDocument {
	mName := fmt.Sprintf(SHIPADDRESS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// books - array -  [books]
func (ud *UpdateDocument) SetBooks(p []Book) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetBooks() *UpdateDocument {
	mName := fmt.Sprintf(BOOKS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// [] - struct - Book [books.[]]
func (ud *UpdateDocument) SetBooksI(ndxI int, p Book) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS_I, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetBooksI(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS_I, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// title - string -  [books.[].title books.title]
func (ud *UpdateDocument) SetBooksITitle(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS_I_TITLE, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetBooksITitle(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS_I_TITLE, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// isbn - string -  [books.[].isbn books.isbn]
func (ud *UpdateDocument) SetBooksIIsbn(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS_I_ISBN, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetBooksIIsbn(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS_I_ISBN, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// coAuthors - array -  [books.[].coAuthors books.coAuthors]
func (ud *UpdateDocument) SetBooksICoAuthors(ndxI int, p []string) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS_I_COAUTHORS, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetBooksICoAuthors(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS_I_COAUTHORS, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// [] - string -  [books.[].coAuthors.[]]
func (ud *UpdateDocument) SetBooksICoAuthorsJ(ndxI int, ndxJ int, p string) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS_I_COAUTHORS_J, ndxI, ndxJ)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetBooksICoAuthorsJ(ndxI int, ndxJ int) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS_I_COAUTHORS_J, ndxI, ndxJ)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// businessRels - map -  [businessRels]
func (ud *UpdateDocument) SetBusinessRels(p map[string]BusinessRel) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetBusinessRels() *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// %s - struct - BusinessRel [businessRels.%s]
func (ud *UpdateDocument) SetBusinessRelsS(keyS string, p BusinessRel) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S, keyS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetBusinessRelsS(keyS string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S, keyS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// publisherId - string -  [businessRels.%s.publisherId]
func (ud *UpdateDocument) SetBusinessRelsSPublisherId(keyS string, p string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_PUBLISHERID, keyS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetBusinessRelsSPublisherId(keyS string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_PUBLISHERID, keyS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// publisherName - string -  [businessRels.%s.publisherName]
func (ud *UpdateDocument) SetBusinessRelsSPublisherName(keyS string, p string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_PUBLISHERNAME, keyS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetBusinessRelsSPublisherName(keyS string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_PUBLISHERNAME, keyS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// contracts - map -  [businessRels.%s.contracts]
func (ud *UpdateDocument) SetBusinessRelsSContracts(keyS string, p map[string]Contract) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS, keyS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetBusinessRelsSContracts(keyS string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS, keyS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// %s - struct - Contract [businessRels.%s.contracts.%s]
func (ud *UpdateDocument) SetBusinessRelsSContractsT(keyS string, keyT string, p Contract) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T, keyS, keyT)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetBusinessRelsSContractsT(keyS string, keyT string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T, keyS, keyT)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// contractId - string -  [businessRels.%s.contracts.%s.contractId]
func (ud *UpdateDocument) SetBusinessRelsSContractsTContractId(keyS string, keyT string, p string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T_CONTRACTID, keyS, keyT)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetBusinessRelsSContractsTContractId(keyS string, keyT string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T_CONTRACTID, keyS, keyT)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// contractDescr - string -  [businessRels.%s.contracts.%s.contractDescr]
func (ud *UpdateDocument) SetBusinessRelsSContractsTContractDescr(keyS string, keyT string, p string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T_CONTRACTDESCR, keyS, keyT)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func (ud *UpdateDocument) UnsetBusinessRelsSContractsTContractDescr(keyS string, keyT string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T_CONTRACTDESCR, keyS, keyT)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}
