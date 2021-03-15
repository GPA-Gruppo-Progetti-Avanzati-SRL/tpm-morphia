package example0

import "go.mongodb.org/mongo-driver/bson/primitive"

type Author struct {
	OId       primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	FirstName string             `json:"fn,omitempty" bson:"fn,omitempty"`
	LastName  string             `json:"ln,omitempty" bson:"ln,omitempty"`
	Age       int32              `json:"age,omitempty" bson:"age,omitempty"`
	Address   Address            `json:"addr,omitempty" bson:"addr,omitempty"`
}

type Address struct {
	City   string `json:"city,omitempty" bson:"city,omitempty"`
	Street string `json:"strt,omitempty" bson:"strt,omitempty"`
}

/*
 * The implementation of this interface is mandatory for the bson omitempty tag to work in case fields are structs
 * and not pointers to structs. In the latter case the nil value do the job...
 */
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
	if !s.Address.IsZero() {
		return false
	}
	return true
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
