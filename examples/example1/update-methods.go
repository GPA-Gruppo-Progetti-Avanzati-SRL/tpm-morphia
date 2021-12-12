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
	DefaultMode UnsetMode
	OId         UnsetMode
	FirstName   UnsetMode
	LastName    UnsetMode
	Age         UnsetMode
	Doc         UnsetMode
	Address     UnsetMode
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
	if len(obj.Doc) > 0 {
		ud.SetDoc(obj.Doc)
	} else {
		um := uo.Doc
		if um == UnSpecified {
			um = uo.DefaultMode
		}
		switch um {
		case KeepCurrent:
		case UnsetData:
			ud.UnsetDoc()
		case SetData2Default:
			ud.UnsetDoc()
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
