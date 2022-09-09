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

type UnsetMode int64

const (
	UnSpecified     UnsetMode = 0
	KeepCurrent               = 1
	UnsetData                 = 2
	SetData2Default           = 3
)

type UnsetOption func(uopt *UnsetOptions)

type UnsetOptions struct {
	DefaultMode UnsetMode
	OId         UnsetMode
	FirstName   UnsetMode
	LastName    UnsetMode
	Age         UnsetMode
	Doc         UnsetMode
	Address     UnsetMode
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
func WithDocUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Doc = m
	}
}
func WithAddressUnsetMode(m UnsetMode) UnsetOption {
	return func(uopt *UnsetOptions) {
		uopt.Address = m
	}
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
	// if len(obj.Doc) > 0 {
	//   ud.SetDoc ( obj.Doc)
	// } else {
	ud.setOrUnsetDoc(obj.Doc, uo.ResolveUnsetMode(uo.Doc))
	// }
	// if !obj.Address.IsZero() {
	//   ud.SetAddress ( obj.Address)
	// } else {
	ud.setOrUnsetAddress(obj.Address, uo.ResolveUnsetMode(uo.Address))
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

//----- doc - document -  [doc]

// SetDoc No Remarks
func (ud *UpdateDocument) SetDoc(p bson.M) *UpdateDocument {
	mName := fmt.Sprintf(DOC)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetDoc No Remarks
func (ud *UpdateDocument) UnsetDoc() *UpdateDocument {
	mName := fmt.Sprintf(DOC)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetDoc No Remarks
func (ud *UpdateDocument) setOrUnsetDoc(p bson.M, um UnsetMode) {
	if len(p) > 0 {
		ud.SetDoc(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetDoc()
		case SetData2Default:
			ud.UnsetDoc()
		}
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

//----- city - string -  [address.city]

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

//----- street - string -  [address.street]

// SetAddressStreet No Remarks
func (ud *UpdateDocument) SetAddressStreet(p string) *UpdateDocument {
	mName := fmt.Sprintf(ADDRESS_STREET)
	ud.Set().Add(func() bson.E {
		return bson.E{Key: mName, Value: p}
	})
	return ud
}

// UnsetAddressStreet No Remarks
func (ud *UpdateDocument) UnsetAddressStreet() *UpdateDocument {
	mName := fmt.Sprintf(ADDRESS_STREET)
	ud.Unset().Add(func() bson.E {
		return bson.E{Key: mName, Value: ""}
	})
	return ud
}

// setOrUnsetAddressStreet No Remarks
func (ud *UpdateDocument) setOrUnsetAddressStreet(p string, um UnsetMode) {
	if p != "" {
		ud.SetAddressStreet(p)
	} else {
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetAddressStreet()
		case SetData2Default:
			ud.UnsetAddressStreet()
		}
	}
}
