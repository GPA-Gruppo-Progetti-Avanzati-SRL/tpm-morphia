package example4

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	OID                    = "_id"
	NDG                    = "ndg"
	CODICEFISCALE          = "codiceFiscale"
	PARTITAIVA             = "partitaIVA"
	NATURA                 = "natura"
	STATO                  = "stato"
	INDIRIZZI              = "indirizzi"
	INDIRIZZI_S            = "indirizzi.%s"
	INDIRIZZI_S_INDIRIZZO  = "indirizzi.%s.indirizzo"
	INDIRIZZI_S_CAP        = "indirizzi.%s.cap"
	INDIRIZZI_S_LOCALITA   = "indirizzi.%s.localita"
	INDIRIZZI_S_PROVINCIA  = "indirizzi.%s.provincia"
	INDIRIZZI_S_NAZIONE    = "indirizzi.%s.nazione"
	LEGATI                 = "legati"
	LEGATI_I               = "legati.%d"
	LEGATI_I_NDG           = "legati.%d.ndg"
	LEGATI_NDG             = "legati.ndg"
	LEGATI_I_COGNOME       = "legati.%d.cognome"
	LEGATI_COGNOME         = "legati.cognome"
	LEGATI_I_NOME          = "legati.%d.nome"
	LEGATI_NOME            = "legati.nome"
	LEGATI_I_CODICEFISCALE = "legati.%d.codiceFiscale"
	LEGATI_CODICEFISCALE   = "legati.codiceFiscale"
	LEGATI_I_PARTITAIVA    = "legati.%d.partitaIVA"
	LEGATI_PARTITAIVA      = "legati.partitaIVA"
	LEGATI_I_NATURA        = "legati.%d.natura"
	LEGATI_NATURA          = "legati.natura"
	LEGANTI                = "leganti"
	LEGANTI_I              = "leganti.%d"
)

type Cliente struct {
	OId           primitive.ObjectID   `json:"-" bson:"_id,omitempty"`
	Ndg           string               `json:"ndg,omitempty" bson:"ndg,omitempty"`
	CodiceFiscale string               `json:"codiceFiscale,omitempty" bson:"codiceFiscale,omitempty"`
	PartitaIVA    string               `json:"partitaIVA,omitempty" bson:"partitaIVA,omitempty"`
	Natura        string               `json:"natura,omitempty" bson:"natura,omitempty"`
	Stato         string               `json:"stato,omitempty" bson:"stato,omitempty"`
	Indirizzi     map[string]Indirizzo `json:"indirizzi,omitempty" bson:"indirizzi,omitempty"`
	Legati        []Legame             `json:"legati,omitempty" bson:"legati,omitempty"`
	Leganti       []Legame             `json:"leganti,omitempty" bson:"leganti,omitempty"`
}

func (s Cliente) IsZero() bool {
	/*
	   if s.OId == primitive.NilObjectID {
	       return false
	   }
	   if s.Ndg == "" {
	       return false
	   }
	   if s.CodiceFiscale == "" {
	       return false
	   }
	   if s.PartitaIVA == "" {
	       return false
	   }
	   if s.Natura == "" {
	       return false
	   }
	   if s.Stato == "" {
	       return false
	   }
	   if len(s.Indirizzi) == 0 {
	       return false
	   }
	   if len(s.Legati) == 0 {
	       return false
	   }
	   if len(s.Leganti) == 0 {
	       return false
	   }
	       return true
	*/

	return s.OId == primitive.NilObjectID && s.Ndg == "" && s.CodiceFiscale == "" && s.PartitaIVA == "" && s.Natura == "" && s.Stato == "" && len(s.Indirizzi) == 0 && len(s.Legati) == 0 && len(s.Leganti) == 0
}

type Indirizzo struct {
	Indirizzo string `json:"indirizzo,omitempty" bson:"indirizzo,omitempty"`
	Cap       string `json:"cap,omitempty" bson:"cap,omitempty"`
	Localita  string `json:"localita,omitempty" bson:"localita,omitempty"`
	Provincia string `json:"provincia,omitempty" bson:"provincia,omitempty"`
	Nazione   string `json:"nazione,omitempty" bson:"nazione,omitempty"`
}

func (s Indirizzo) IsZero() bool {
	/*
	   if s.Indirizzo == "" {
	       return false
	   }
	   if s.Cap == "" {
	       return false
	   }
	   if s.Localita == "" {
	       return false
	   }
	   if s.Provincia == "" {
	       return false
	   }
	   if s.Nazione == "" {
	       return false
	   }
	       return true
	*/
	return s.Indirizzo == "" && s.Cap == "" && s.Localita == "" && s.Provincia == "" && s.Nazione == ""
}

type Legame struct {
	Ndg           string `json:"ndg,omitempty" bson:"ndg,omitempty"`
	Cognome       string `json:"cognome,omitempty" bson:"cognome,omitempty"`
	Nome          string `json:"nome,omitempty" bson:"nome,omitempty"`
	CodiceFiscale string `json:"codiceFiscale,omitempty" bson:"codiceFiscale,omitempty"`
	PartitaIVA    string `json:"partitaIVA,omitempty" bson:"partitaIVA,omitempty"`
	Natura        string `json:"natura,omitempty" bson:"natura,omitempty"`
}

func (s Legame) IsZero() bool {
	/*
	   if s.Ndg == "" {
	       return false
	   }
	   if s.Cognome == "" {
	       return false
	   }
	   if s.Nome == "" {
	       return false
	   }
	   if s.CodiceFiscale == "" {
	       return false
	   }
	   if s.PartitaIVA == "" {
	       return false
	   }
	   if s.Natura == "" {
	       return false
	   }
	       return true
	*/
	return s.Ndg == "" && s.Cognome == "" && s.Nome == "" && s.CodiceFiscale == "" && s.PartitaIVA == "" && s.Natura == ""
}
