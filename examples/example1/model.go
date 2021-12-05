package example1

import "go.mongodb.org/mongo-driver/bson/primitive"
import "go.mongodb.org/mongo-driver/bson"

const (
	OID            = "_id"
	FIRSTNAME      = "fn"
	LASTNAME       = "ln"
	AGE            = "age"
	DOC            = "doc"
	ADDRESS        = "addr"
	ADDRESS_CITY   = "addr.city"
	ADDRESS_STREET = "addr.street"
)

type Author struct {
	OId       primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	FirstName string             `json:"fn,omitempty" bson:"fn,omitempty"`
	LastName  string             `json:"ln,omitempty" bson:"ln,omitempty"`
	Age       int32              `json:"age,omitempty" bson:"age,omitempty"`
	Doc       bson.M             `json:"doc,omitempty" bson:"doc,omitempty"`
	Address   Address            `json:"addr,omitempty" bson:"addr,omitempty"`
}

func (s Author) IsZero() bool {
	if s.OId != primitive.NilObjectID {
		return false
	}
	if s.FirstName != "" {
		return false
	}
	if s.LastName != "" {
		return false
	}
	if s.Age != 0 {
		return false
	}
	if len(s.Doc) != 0 {
		return false
	}
	if !s.Address.IsZero() {
		return false
	}
	return true
}

type Address struct {
	City   string `json:"city,omitempty" bson:"city,omitempty"`
	Street string `json:"street,omitempty" bson:"street,omitempty"`
}

func (s Address) IsZero() bool {
	if s.City != "" {
		return false
	}
	if s.Street != "" {
		return false
	}
	return true
}
