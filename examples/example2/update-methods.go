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

func (uo *UnsetOptions) ResolveUnsetMode(um UnsetMode) UnsetMode {
	if um == UnSpecified {
		um = uo.DefaultMode
	}

	return um
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

type UpdateOption func(ud *UpdateDocument)
type UpdateOptions []UpdateOption

// GetUpdateDocument convenience method to create an update document from single updates instead of a whole object
func GetUpdateDocumentFromOptions(opts ...UpdateOption) UpdateDocument {
	ud := UpdateDocument{}
	for _, o := range opts {
		o(&ud)
	}

	return ud
}

// GetUpdateDocument
// Convenience method to create an Update Document from the values of the top fields of the object. The convenience is in the handling
// the unset because if I pass an empty struct to the update it generates an empty object anyway in the db. Handling the unset eliminates
// the issue and delete an existing value without creating an empty struct.
func GetUpdateDocument(obj *Author, opts ...UnsetOption) UpdateDocument {

	uo := &UnsetOptions{DefaultMode: KeepCurrent}
	for _, o := range opts {
		o(uo)
	}

	ud := UpdateDocument{}
	ud.setOrUnsetFirstName(obj.FirstName, uo.ResolveUnsetMode(uo.FirstName))
	ud.setOrUnsetLastName(obj.LastName, uo.ResolveUnsetMode(uo.LastName))
	ud.setOrUnsetAge(obj.Age, uo.ResolveUnsetMode(uo.Age))
	// if !obj.Address.IsZero() {
	//   ud.SetAddress ( obj.Address)
	// } else {
	ud.setOrUnsetAddress(obj.Address, uo.ResolveUnsetMode(uo.Address))
	// }
	// if !obj.ShipAddress.IsZero() {
	//   ud.SetShipAddress ( obj.ShipAddress)
	// } else {
	ud.setOrUnsetShipAddress(obj.ShipAddress, uo.ResolveUnsetMode(uo.ShipAddress))
	// }
	// if len(obj.Books) > 0 {
	//   ud.SetBooks ( obj.Books)
	// } else {
	ud.setOrUnsetBooks(obj.Books, uo.ResolveUnsetMode(uo.Books))
	// }
	// if len(obj.BusinessRels) > 0 {
	//   ud.SetBusinessRels ( obj.BusinessRels)
	// } else {
	ud.setOrUnsetBusinessRels(obj.BusinessRels, uo.ResolveUnsetMode(uo.BusinessRels))
	// }

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

// setOrUnsetOId No Remarks
func (ud *UpdateDocument) setOrUnsetOId(p primitive.ObjectID, um UnsetMode) {
	if !p.IsZero() {
		ud.SetOId(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetOId()
		case SetData2Default:
			ud.UnsetOId()
		}
	}
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

// setOrUnsetFirstName No Remarks
func (ud *UpdateDocument) setOrUnsetFirstName(p string, um UnsetMode) {
	if p != "" {
		ud.SetFirstName(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetFirstName()
		case SetData2Default:
			ud.UnsetFirstName()
		}
	}
}

func UpdateWithFirstName(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetFirstName(p)
		} else {
			ud.UnsetFirstName()
		}
	}
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

// setOrUnsetLastName No Remarks
func (ud *UpdateDocument) setOrUnsetLastName(p string, um UnsetMode) {
	if p != "" {
		ud.SetLastName(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetLastName()
		case SetData2Default:
			ud.UnsetLastName()
		}
	}
}

func UpdateWithLastName(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetLastName(p)
		} else {
			ud.UnsetLastName()
		}
	}
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

// setOrUnsetAge No Remarks
func (ud *UpdateDocument) setOrUnsetAge(p int32, um UnsetMode) {
	if p != 0 {
		ud.SetAge(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAge()
		case SetData2Default:
			ud.UnsetAge()
		}
	}
}

// IncAge No Remarks
func (ud *UpdateDocument) IncAge(p int32) *UpdateDocument {
	mName := fmt.Sprintf(AGE)
	ud.Inc().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

func UpdateWithAge(p int32) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != 0 {
			ud.SetAge(p)
		} else {
			ud.UnsetAge()
		}
	}
}

func UpdateWithAgeIncrement(p int32) UpdateOption {
	return func(ud *UpdateDocument) {
		ud.IncAge(p)
	}
}

// ----- address - struct - Address [address]
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

// setOrUnsetAddress No Remarks - here2
func (ud *UpdateDocument) setOrUnsetAddress(p Address, um UnsetMode) {

	//----- struct\n

	if !p.IsZero() {
		ud.SetAddress(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAddress()
		case SetData2Default:
			ud.UnsetAddress()
		}
	}
}

func UpdateWithAddress(p Address) UpdateOption {
	return func(ud *UpdateDocument) {

		if !p.IsZero() {
			ud.SetAddress(p)
		} else {
			ud.UnsetAddress()
		}
	}
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

// setOrUnsetAddressCity No Remarks
func (ud *UpdateDocument) setOrUnsetAddressCity(p string, um UnsetMode) {
	if p != "" {
		ud.SetAddressCity(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAddressCity()
		case SetData2Default:
			ud.UnsetAddressCity()
		}
	}
}

func UpdateWithAddressCity(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetAddressCity(p)
		} else {
			ud.UnsetAddressCity()
		}
	}
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

// setOrUnsetShipAddressCity No Remarks
func (ud *UpdateDocument) setOrUnsetShipAddressCity(p string, um UnsetMode) {
	if p != "" {
		ud.SetShipAddressCity(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetShipAddressCity()
		case SetData2Default:
			ud.UnsetShipAddressCity()
		}
	}
}

func UpdateWithShipAddressCity(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetShipAddressCity(p)
		} else {
			ud.UnsetShipAddressCity()
		}
	}
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

// setOrUnsetAddressStrt No Remarks
func (ud *UpdateDocument) setOrUnsetAddressStrt(p string, um UnsetMode) {
	if p != "" {
		ud.SetAddressStrt(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAddressStrt()
		case SetData2Default:
			ud.UnsetAddressStrt()
		}
	}
}

func UpdateWithAddressStrt(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetAddressStrt(p)
		} else {
			ud.UnsetAddressStrt()
		}
	}
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

// setOrUnsetShipAddressStrt No Remarks
func (ud *UpdateDocument) setOrUnsetShipAddressStrt(p string, um UnsetMode) {
	if p != "" {
		ud.SetShipAddressStrt(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetShipAddressStrt()
		case SetData2Default:
			ud.UnsetShipAddressStrt()
		}
	}
}

func UpdateWithShipAddressStrt(p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetShipAddressStrt(p)
		} else {
			ud.UnsetShipAddressStrt()
		}
	}
}

// ----- shipAddress - ref-struct -  [shipAddress]
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

// setOrUnsetShipAddress No Remarks - here2
func (ud *UpdateDocument) setOrUnsetShipAddress(p Address, um UnsetMode) {

	//----- ref-struct\n

	if !p.IsZero() {
		ud.SetShipAddress(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetShipAddress()
		case SetData2Default:
			ud.UnsetShipAddress()
		}
	}
}

func UpdateWithShipAddress(p Address) UpdateOption {
	return func(ud *UpdateDocument) {

		if !p.IsZero() {
			ud.SetShipAddress(p)
		} else {
			ud.UnsetShipAddress()
		}
	}
}

// ----- books - array -  [books]
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

// setOrUnsetBooks No Remarks - here2
func (ud *UpdateDocument) setOrUnsetBooks(p []Book, um UnsetMode) {

	//----- array\n

	if len(p) > 0 {
		ud.SetBooks(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetBooks()
		case SetData2Default:
			ud.UnsetBooks()
		}
	}
}

func UpdateWithBooks(p []Book) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetBooks(p)
		} else {
			ud.UnsetBooks()
		}
	}
}

// ----- [] - struct - Book [books.[]]
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

// setOrUnsetBooksI No Remarks
func (ud *UpdateDocument) setOrUnsetBooksI(ndxI int, p Book, um UnsetMode) {
	if !p.IsZero() {
		ud.SetBooksI(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetBooksI(ndxI)
		case SetData2Default:
			ud.UnsetBooksI(ndxI)
		}
	}
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

// setOrUnsetBooksITitle No Remarks
func (ud *UpdateDocument) setOrUnsetBooksITitle(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetBooksITitle(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetBooksITitle(ndxI)
		case SetData2Default:
			ud.UnsetBooksITitle(ndxI)
		}
	}
}

func UpdateWithBooksITitle(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetBooksITitle(ndxI, p)
		} else {
			ud.UnsetBooksITitle(ndxI)
		}
	}
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

// setOrUnsetBooksIIsbn No Remarks
func (ud *UpdateDocument) setOrUnsetBooksIIsbn(ndxI int, p string, um UnsetMode) {
	if p != "" {
		ud.SetBooksIIsbn(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetBooksIIsbn(ndxI)
		case SetData2Default:
			ud.UnsetBooksIIsbn(ndxI)
		}
	}
}

func UpdateWithBooksIIsbn(ndxI int, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetBooksIIsbn(ndxI, p)
		} else {
			ud.UnsetBooksIIsbn(ndxI)
		}
	}
}

// ----- coAuthors - array -  [books.[].coAuthors books.coAuthors]
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

// setOrUnsetBooksICoAuthors No Remarks - here2
func (ud *UpdateDocument) setOrUnsetBooksICoAuthors(ndxI int, p []string, um UnsetMode) {

	//----- array\n

	if len(p) > 0 {
		ud.SetBooksICoAuthors(ndxI, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetBooksICoAuthors(ndxI)
		case SetData2Default:
			ud.UnsetBooksICoAuthors(ndxI)
		}
	}
}

func UpdateWithBooksICoAuthors(ndxI int, p []string) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetBooksICoAuthors(ndxI, p)
		} else {
			ud.UnsetBooksICoAuthors(ndxI)
		}
	}
}

// ----- [] - string -  [books.[].coAuthors.[]]
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

// setOrUnsetBooksICoAuthorsJ No Remarks
func (ud *UpdateDocument) setOrUnsetBooksICoAuthorsJ(ndxI int, ndxJ int, p string, um UnsetMode) {
	// Warining.... should not get here
	if p != "" {
		ud.SetBooksICoAuthorsJ(ndxI, ndxJ, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetBooksICoAuthorsJ(ndxI, ndxJ)
		case SetData2Default:
			ud.UnsetBooksICoAuthorsJ(ndxI, ndxJ)
		}
	}
}

// ----- businessRels - map -  [businessRels]
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

// setOrUnsetBusinessRels No Remarks - here2
func (ud *UpdateDocument) setOrUnsetBusinessRels(p map[string]BusinessRel, um UnsetMode) {

	//----- map\n

	if len(p) > 0 {
		ud.SetBusinessRels(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetBusinessRels()
		case SetData2Default:
			ud.UnsetBusinessRels()
		}
	}
}

func UpdateWithBusinessRels(p map[string]BusinessRel) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetBusinessRels(p)
		} else {
			ud.UnsetBusinessRels()
		}
	}
}

// ----- %s - struct - BusinessRel [businessRels.%s]
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

// setOrUnsetBusinessRelsS No Remarks
func (ud *UpdateDocument) setOrUnsetBusinessRelsS(keyS string, p BusinessRel, um UnsetMode) {
	if !p.IsZero() {
		ud.SetBusinessRelsS(keyS, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetBusinessRelsS(keyS)
		case SetData2Default:
			ud.UnsetBusinessRelsS(keyS)
		}
	}
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

// setOrUnsetBusinessRelsSPublisherId No Remarks
func (ud *UpdateDocument) setOrUnsetBusinessRelsSPublisherId(keyS string, p string, um UnsetMode) {
	if p != "" {
		ud.SetBusinessRelsSPublisherId(keyS, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetBusinessRelsSPublisherId(keyS)
		case SetData2Default:
			ud.UnsetBusinessRelsSPublisherId(keyS)
		}
	}
}

func UpdateWithBusinessRelsSPublisherId(keyS string, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetBusinessRelsSPublisherId(keyS, p)
		} else {
			ud.UnsetBusinessRelsSPublisherId(keyS)
		}
	}
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

// setOrUnsetBusinessRelsSPublisherName No Remarks
func (ud *UpdateDocument) setOrUnsetBusinessRelsSPublisherName(keyS string, p string, um UnsetMode) {
	if p != "" {
		ud.SetBusinessRelsSPublisherName(keyS, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetBusinessRelsSPublisherName(keyS)
		case SetData2Default:
			ud.UnsetBusinessRelsSPublisherName(keyS)
		}
	}
}

func UpdateWithBusinessRelsSPublisherName(keyS string, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetBusinessRelsSPublisherName(keyS, p)
		} else {
			ud.UnsetBusinessRelsSPublisherName(keyS)
		}
	}
}

// ----- contracts - map -  [businessRels.%s.contracts]
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

// setOrUnsetBusinessRelsSContracts No Remarks - here2
func (ud *UpdateDocument) setOrUnsetBusinessRelsSContracts(keyS string, p map[string]Contract, um UnsetMode) {

	//----- map\n

	if len(p) > 0 {
		ud.SetBusinessRelsSContracts(keyS, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetBusinessRelsSContracts(keyS)
		case SetData2Default:
			ud.UnsetBusinessRelsSContracts(keyS)
		}
	}
}

func UpdateWithBusinessRelsSContracts(keyS string, p map[string]Contract) UpdateOption {
	return func(ud *UpdateDocument) {
		if len(p) > 0 {
			ud.SetBusinessRelsSContracts(keyS, p)
		} else {
			ud.UnsetBusinessRelsSContracts(keyS)
		}
	}
}

// ----- %s - struct - Contract [businessRels.%s.contracts.%s]
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

// setOrUnsetBusinessRelsSContractsT No Remarks
func (ud *UpdateDocument) setOrUnsetBusinessRelsSContractsT(keyS string, keyT string, p Contract, um UnsetMode) {
	if !p.IsZero() {
		ud.SetBusinessRelsSContractsT(keyS, keyT, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetBusinessRelsSContractsT(keyS, keyT)
		case SetData2Default:
			ud.UnsetBusinessRelsSContractsT(keyS, keyT)
		}
	}
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

// setOrUnsetBusinessRelsSContractsTContractId No Remarks
func (ud *UpdateDocument) setOrUnsetBusinessRelsSContractsTContractId(keyS string, keyT string, p string, um UnsetMode) {
	if p != "" {
		ud.SetBusinessRelsSContractsTContractId(keyS, keyT, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetBusinessRelsSContractsTContractId(keyS, keyT)
		case SetData2Default:
			ud.UnsetBusinessRelsSContractsTContractId(keyS, keyT)
		}
	}
}

func UpdateWithBusinessRelsSContractsTContractId(keyS string, keyT string, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetBusinessRelsSContractsTContractId(keyS, keyT, p)
		} else {
			ud.UnsetBusinessRelsSContractsTContractId(keyS, keyT)
		}
	}
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

// setOrUnsetBusinessRelsSContractsTContractDescr No Remarks
func (ud *UpdateDocument) setOrUnsetBusinessRelsSContractsTContractDescr(keyS string, keyT string, p string, um UnsetMode) {
	if p != "" {
		ud.SetBusinessRelsSContractsTContractDescr(keyS, keyT, p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetBusinessRelsSContractsTContractDescr(keyS, keyT)
		case SetData2Default:
			ud.UnsetBusinessRelsSContractsTContractDescr(keyS, keyT)
		}
	}
}

func UpdateWithBusinessRelsSContractsTContractDescr(keyS string, keyT string, p string) UpdateOption {
	return func(ud *UpdateDocument) {
		if p != "" {
			ud.SetBusinessRelsSContractsTContractDescr(keyS, keyT, p)
		} else {
			ud.UnsetBusinessRelsSContractsTContractDescr(keyS, keyT)
		}
	}
}
