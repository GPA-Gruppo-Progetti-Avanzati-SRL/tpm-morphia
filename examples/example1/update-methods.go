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
	if len(obj.Doc) > 0 {
		ud.SetDoc(obj.Doc)
	} else {
		ud.UnsetDoc()
	}
	if !obj.Address.IsZero() {
		ud.SetAddress(obj.Address)
	} else {
		ud.UnsetAddress()
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
