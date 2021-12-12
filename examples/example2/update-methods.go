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

// GetUpdateDocument
// Convenience method to create an Update Document from the values of the top fields of the object. The convenience is in the handling
// the unset because if I pass an empty struct to the update it generates an empty object anyway in the db. Handling the unset eliminates
// the issue and delete an existing value without creating an empty struct.
type UnsetMode int64

const (
	UnSpecified     UnsetMode = 0
	KeepCurrent               = 1
	UnsetData                 = 2
	SetData2Default           = 3
)

type UnsetOption func(uopt *UnsetOptions)

type UnsetOptions struct {
	DefaultMode  UnsetMode
	OId          UnsetMode
	FirstName    UnsetMode
	LastName     UnsetMode
	Age          UnsetMode
	Address      UnsetMode
	ShipAddress  UnsetMode
	Books        UnsetMode
	BusinessRels UnsetMode
}

func WithDefaultUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.DefaultMode = m
	}
}
func WithOIdUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.OId = m
	}
}
func WithFirstNameUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.FirstName = m
	}
}
func WithLastNameUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.LastName = m
	}
}
func WithAgeUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Age = m
	}
}
func WithAddressUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Address = m
	}
}
func WithShipAddressUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.ShipAddress = m
	}
}
func WithBooksUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Books = m
	}
}
func WithBusinessRelsUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.BusinessRels = m
	}
}

func GetUpdateDocument(obj *Author, opts ...UnsetOption) UpdateDocument {

	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}

	ud := UpdateDocument{}
	if obj.FirstName != "" {
		ud.SetFirstName(obj.FirstName)
	} else {
		um := uo.FirstName
		if um == UnSpecified {
			um = uo.DefaultMode
		}
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetFirstName()
		case SetData2Default:
			ud.UnsetFirstName()
		}
	}
	if obj.LastName != "" {
		ud.SetLastName(obj.LastName)
	} else {
		um := uo.LastName
		if um == UnSpecified {
			um = uo.DefaultMode
		}
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetLastName()
		case SetData2Default:
			ud.UnsetLastName()
		}
	}
	if obj.Age != 0 {
		ud.SetAge(obj.Age)
	} else {
		um := uo.Age
		if um == UnSpecified {
			um = uo.DefaultMode
		}
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAge()
		case SetData2Default:
			ud.SetAge(0)
		}
	}
	if !obj.Address.IsZero() {
		ud.SetAddress(obj.Address)
	} else {
		um := uo.Address
		if um == UnSpecified {
			um = uo.DefaultMode
		}
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAddress()
		case SetData2Default:
			ud.UnsetAddress()
		}
	}
	if !obj.ShipAddress.IsZero() {
		ud.SetShipAddress(obj.ShipAddress)
	} else {
		um := uo.ShipAddress
		if um == UnSpecified {
			um = uo.DefaultMode
		}
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetShipAddress()
		case SetData2Default:
			ud.UnsetShipAddress()
		}
	}
	if len(obj.Books) > 0 {
		ud.SetBooks(obj.Books)
	} else {
		um := uo.Books
		if um == UnSpecified {
			um = uo.DefaultMode
		}
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetBooks()
		case SetData2Default:
			ud.UnsetBooks()
		}
	}
	if len(obj.BusinessRels) > 0 {
		ud.SetBusinessRels(obj.BusinessRels)
	} else {
		um := uo.BusinessRels
		if um == UnSpecified {
			um = uo.DefaultMode
		}
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetBusinessRels()
		case SetData2Default:
			ud.UnsetBusinessRels()
		}
	}

	return ud
}

//----- oId - object-id -  [oId]

// SetOId No Remarks
func (ud *UpdateDocument) SetOId(p primitive.ObjectID) *UpdateDocument {
	mName := fmt.Sprintf(OID)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetOId No Remarks
func (ud *UpdateDocument) UnsetOId() *UpdateDocument {
	mName := fmt.Sprintf(OID)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

//----- firstName - string -  [firstName]

// SetFirstName No Remarks
func (ud *UpdateDocument) SetFirstName(p string) *UpdateDocument {
	mName := fmt.Sprintf(FIRSTNAME)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetFirstName No Remarks
func (ud *UpdateDocument) UnsetFirstName() *UpdateDocument {
	mName := fmt.Sprintf(FIRSTNAME)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

//----- lastName - string -  [lastName]

// SetLastName No Remarks
func (ud *UpdateDocument) SetLastName(p string) *UpdateDocument {
	mName := fmt.Sprintf(LASTNAME)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetLastName No Remarks
func (ud *UpdateDocument) UnsetLastName() *UpdateDocument {
	mName := fmt.Sprintf(LASTNAME)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

//----- age - int -  [age]

// SetAge No Remarks
func (ud *UpdateDocument) SetAge(p int32) *UpdateDocument {
	mName := fmt.Sprintf(AGE)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetAge No Remarks
func (ud *UpdateDocument) UnsetAge() *UpdateDocument {
	mName := fmt.Sprintf(AGE)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// IncAge No Remarks
func (ud *UpdateDocument) IncAge(p int32) *UpdateDocument {
	mName := fmt.Sprintf(AGE)
	ud.Inc().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

//----- address - struct - Address [address]
// SetAddress No Remarks
func (ud *UpdateDocument) SetAddress(p Address) *UpdateDocument {
	mName := fmt.Sprintf(ADDRESS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetAddress No Remarks
func (ud *UpdateDocument) UnsetAddress() *UpdateDocument {
	mName := fmt.Sprintf(ADDRESS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

//----- city - string -  [address.city shipAddress.city]

// SetAddressCity No Remarks
func (ud *UpdateDocument) SetAddressCity(p string) *UpdateDocument {
	mName := fmt.Sprintf(ADDRESS_CITY)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetAddressCity No Remarks
func (ud *UpdateDocument) UnsetAddressCity() *UpdateDocument {
	mName := fmt.Sprintf(ADDRESS_CITY)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// SetShipAddressCity No Remarks
func (ud *UpdateDocument) SetShipAddressCity(p string) *UpdateDocument {
	mName := fmt.Sprintf(SHIPADDRESS_CITY)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetShipAddressCity No Remarks
func (ud *UpdateDocument) UnsetShipAddressCity() *UpdateDocument {
	mName := fmt.Sprintf(SHIPADDRESS_CITY)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

//----- strt - string -  [address.strt shipAddress.strt]

// SetAddressStrt No Remarks
func (ud *UpdateDocument) SetAddressStrt(p string) *UpdateDocument {
	mName := fmt.Sprintf(ADDRESS_STRT)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetAddressStrt No Remarks
func (ud *UpdateDocument) UnsetAddressStrt() *UpdateDocument {
	mName := fmt.Sprintf(ADDRESS_STRT)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// SetShipAddressStrt No Remarks
func (ud *UpdateDocument) SetShipAddressStrt(p string) *UpdateDocument {
	mName := fmt.Sprintf(SHIPADDRESS_STRT)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetShipAddressStrt No Remarks
func (ud *UpdateDocument) UnsetShipAddressStrt() *UpdateDocument {
	mName := fmt.Sprintf(SHIPADDRESS_STRT)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

//----- shipAddress - ref-struct -  [shipAddress]
// SetShipAddress No Remarks
func (ud *UpdateDocument) SetShipAddress(p Address) *UpdateDocument {
	mName := fmt.Sprintf(SHIPADDRESS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetShipAddress No Remarks
func (ud *UpdateDocument) UnsetShipAddress() *UpdateDocument {
	mName := fmt.Sprintf(SHIPADDRESS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

//----- books - array -  [books]
// SetBooks No Remarks
func (ud *UpdateDocument) SetBooks(p []Book) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetBooks No Remarks
func (ud *UpdateDocument) UnsetBooks() *UpdateDocument {
	mName := fmt.Sprintf(BOOKS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

//----- [] - struct - Book [books.[]]
// SetBooksI No Remarks
func (ud *UpdateDocument) SetBooksI(ndxI int, p Book) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS_I, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetBooksI No Remarks
func (ud *UpdateDocument) UnsetBooksI(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS_I, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

//----- title - string -  [books.[].title books.title]

// SetBooksITitle No Remarks
func (ud *UpdateDocument) SetBooksITitle(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS_I_TITLE, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetBooksITitle No Remarks
func (ud *UpdateDocument) UnsetBooksITitle(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS_I_TITLE, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

//----- isbn - string -  [books.[].isbn books.isbn]

// SetBooksIIsbn No Remarks
func (ud *UpdateDocument) SetBooksIIsbn(ndxI int, p string) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS_I_ISBN, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetBooksIIsbn No Remarks
func (ud *UpdateDocument) UnsetBooksIIsbn(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS_I_ISBN, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

//----- coAuthors - array -  [books.[].coAuthors books.coAuthors]
// SetBooksICoAuthors No Remarks
func (ud *UpdateDocument) SetBooksICoAuthors(ndxI int, p []string) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS_I_COAUTHORS, ndxI)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetBooksICoAuthors No Remarks
func (ud *UpdateDocument) UnsetBooksICoAuthors(ndxI int) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS_I_COAUTHORS, ndxI)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

//----- [] - string -  [books.[].coAuthors.[]]
// SetBooksICoAuthorsJ No Remarks
func (ud *UpdateDocument) SetBooksICoAuthorsJ(ndxI int, ndxJ int, p string) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS_I_COAUTHORS_J, ndxI, ndxJ)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetBooksICoAuthorsJ No Remarks
func (ud *UpdateDocument) UnsetBooksICoAuthorsJ(ndxI int, ndxJ int) *UpdateDocument {
	mName := fmt.Sprintf(BOOKS_I_COAUTHORS_J, ndxI, ndxJ)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

//----- businessRels - map -  [businessRels]
// SetBusinessRels No Remarks
func (ud *UpdateDocument) SetBusinessRels(p map[string]BusinessRel) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetBusinessRels No Remarks
func (ud *UpdateDocument) UnsetBusinessRels() *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

//----- %s - struct - BusinessRel [businessRels.%s]
// SetBusinessRelsS No Remarks
func (ud *UpdateDocument) SetBusinessRelsS(keyS string, p BusinessRel) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S, keyS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetBusinessRelsS No Remarks
func (ud *UpdateDocument) UnsetBusinessRelsS(keyS string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S, keyS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

//----- publisherId - string -  [businessRels.%s.publisherId]

// SetBusinessRelsSPublisherId No Remarks
func (ud *UpdateDocument) SetBusinessRelsSPublisherId(keyS string, p string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_PUBLISHERID, keyS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetBusinessRelsSPublisherId No Remarks
func (ud *UpdateDocument) UnsetBusinessRelsSPublisherId(keyS string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_PUBLISHERID, keyS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

//----- publisherName - string -  [businessRels.%s.publisherName]

// SetBusinessRelsSPublisherName No Remarks
func (ud *UpdateDocument) SetBusinessRelsSPublisherName(keyS string, p string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_PUBLISHERNAME, keyS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetBusinessRelsSPublisherName No Remarks
func (ud *UpdateDocument) UnsetBusinessRelsSPublisherName(keyS string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_PUBLISHERNAME, keyS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

//----- contracts - map -  [businessRels.%s.contracts]
// SetBusinessRelsSContracts No Remarks
func (ud *UpdateDocument) SetBusinessRelsSContracts(keyS string, p map[string]Contract) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS, keyS)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetBusinessRelsSContracts No Remarks
func (ud *UpdateDocument) UnsetBusinessRelsSContracts(keyS string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS, keyS)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

//----- %s - struct - Contract [businessRels.%s.contracts.%s]
// SetBusinessRelsSContractsT No Remarks
func (ud *UpdateDocument) SetBusinessRelsSContractsT(keyS string, keyT string, p Contract) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T, keyS, keyT)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetBusinessRelsSContractsT No Remarks
func (ud *UpdateDocument) UnsetBusinessRelsSContractsT(keyS string, keyT string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T, keyS, keyT)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

//----- contractId - string -  [businessRels.%s.contracts.%s.contractId]

// SetBusinessRelsSContractsTContractId No Remarks
func (ud *UpdateDocument) SetBusinessRelsSContractsTContractId(keyS string, keyT string, p string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T_CONTRACTID, keyS, keyT)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetBusinessRelsSContractsTContractId No Remarks
func (ud *UpdateDocument) UnsetBusinessRelsSContractsTContractId(keyS string, keyT string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T_CONTRACTID, keyS, keyT)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

//----- contractDescr - string -  [businessRels.%s.contracts.%s.contractDescr]

// SetBusinessRelsSContractsTContractDescr No Remarks
func (ud *UpdateDocument) SetBusinessRelsSContractsTContractDescr(keyS string, keyT string, p string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T_CONTRACTDESCR, keyS, keyT)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetBusinessRelsSContractsTContractDescr No Remarks
func (ud *UpdateDocument) UnsetBusinessRelsSContractsTContractDescr(keyS string, keyT string) *UpdateDocument {
	mName := fmt.Sprintf(BUSINESSRELS_S_CONTRACTS_T_CONTRACTDESCR, keyS, keyT)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}
