package example2

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	OID                                      = "_id"
	FIRSTNAME                                = "fn"
	LASTNAME                                 = "ln"
	AGE                                      = "age"
	ADDRESS                                  = "addr"
	ADDRESS_CITY                             = "addr.city"
	SHIPADDRESS_CITY                         = "shipaddr.city"
	ADDRESS_STRT                             = "addr.strt"
	SHIPADDRESS_STRT                         = "shipaddr.strt"
	SHIPADDRESS                              = "shipaddr"
	BOOKS                                    = "books"
	BOOKS_I                                  = "books.%d"
	BOOKS_I_TITLE                            = "books.%d.title"
	BOOKS_TITLE                              = "books.title"
	BOOKS_I_ISBN                             = "books.%d.isbn"
	BOOKS_ISBN                               = "books.isbn"
	BOOKS_I_COAUTHORS                        = "books.%d.coAuthors"
	BOOKS_COAUTHORS                          = "books.coAuthors"
	BOOKS_I_COAUTHORS_J                      = "books.%d.coAuthors.%d"
	BUSINESSRELS                             = "businessRels"
	BUSINESSRELS_S                           = "businessRels.%s"
	BUSINESSRELS_S_PUBLISHERID               = "businessRels.%s.publisherId"
	BUSINESSRELS_S_PUBLISHERNAME             = "businessRels.%s.publisherName"
	BUSINESSRELS_S_CONTRACTS                 = "businessRels.%s.contracts"
	BUSINESSRELS_S_CONTRACTS_T               = "businessRels.%s.contracts.%s"
	BUSINESSRELS_S_CONTRACTS_T_CONTRACTID    = "businessRels.%s.contracts.%s.contractId"
	BUSINESSRELS_S_CONTRACTS_T_CONTRACTDESCR = "businessRels.%s.contracts.%s.contractDescr"
)

type Author struct {
	OId          primitive.ObjectID     `json:"-" bson:"_id,omitempty"`
	FirstName    string                 `json:"fn,omitempty" bson:"fn,omitempty"`
	LastName     string                 `json:"ln,omitempty" bson:"ln,omitempty"`
	Age          int32                  `json:"age,omitempty" bson:"age,omitempty"`
	Address      Address                `json:"addr,omitempty" bson:"addr,omitempty"`
	ShipAddress  Address                `json:"shipaddr,omitempty" bson:"shipaddr,omitempty"`
	Books        []Book                 `json:"books,omitempty" bson:"books,omitempty"`
	BusinessRels map[string]BusinessRel `json:"businessRels,omitempty" bson:"businessRels,omitempty"`
}

func (s Author) IsZero() bool {
	/*
	   if s.OId == primitive.NilObjectID {
	       return false
	   }
	   if s.FirstName == "" {
	       return false
	   }
	   if s.LastName == "" {
	       return false
	   }
	   if s.Age == 0 {
	       return false
	   }
	   if s.Address.IsZero() {
	       return false
	   }
	   if s.ShipAddress.IsZero() {
	       return false
	   }
	   if len(s.Books) == 0 {
	       return false
	   }
	   if len(s.BusinessRels) == 0 {
	       return false
	   }
	       return true
	*/

	return s.OId == primitive.NilObjectID && s.FirstName == "" && s.LastName == "" && s.Age == 0 && s.Address.IsZero() && s.ShipAddress.IsZero() && len(s.Books) == 0 && len(s.BusinessRels) == 0
}

type Address struct {
	City string `json:"city,omitempty" bson:"city,omitempty"`
	Strt string `json:"strt,omitempty" bson:"strt,omitempty"`
}

func (s Address) IsZero() bool {
	/*
	   if s.City == "" {
	       return false
	   }
	   if s.Strt == "" {
	       return false
	   }
	       return true
	*/
	return s.City == "" && s.Strt == ""
}

type Book struct {
	Title     string   `json:"title,omitempty" bson:"title,omitempty"`
	Isbn      string   `json:"isbn,omitempty" bson:"isbn,omitempty"`
	CoAuthors []string `json:"coAuthors,omitempty" bson:"coAuthors,omitempty"`
}

func (s Book) IsZero() bool {
	/*
	   if s.Title == "" {
	       return false
	   }
	   if s.Isbn == "" {
	       return false
	   }
	   if len(s.CoAuthors) == 0 {
	       return false
	   }
	       return true
	*/
	return s.Title == "" && s.Isbn == "" && len(s.CoAuthors) == 0
}

type BusinessRel struct {
	PublisherId   string              `json:"publisherId,omitempty" bson:"publisherId,omitempty"`
	PublisherName string              `json:"publisherName,omitempty" bson:"publisherName,omitempty"`
	Contracts     map[string]Contract `json:"contracts,omitempty" bson:"contracts,omitempty"`
}

func (s BusinessRel) IsZero() bool {
	/*
	   if s.PublisherId == "" {
	       return false
	   }
	   if s.PublisherName == "" {
	       return false
	   }
	   if len(s.Contracts) == 0 {
	       return false
	   }
	       return true
	*/
	return s.PublisherId == "" && s.PublisherName == "" && len(s.Contracts) == 0
}

type Contract struct {
	ContractId    string `json:"contractId,omitempty" bson:"contractId,omitempty"`
	ContractDescr string `json:"contractDescr,omitempty" bson:"contractDescr,omitempty"`
}

func (s Contract) IsZero() bool {
	/*
	   if s.ContractId == "" {
	       return false
	   }
	   if s.ContractDescr == "" {
	       return false
	   }
	       return true
	*/
	return s.ContractId == "" && s.ContractDescr == ""
}
