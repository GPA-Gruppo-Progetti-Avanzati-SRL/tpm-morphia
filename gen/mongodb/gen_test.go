package mongodb

import (
	"bytes"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/config"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schema"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system"
	"testing"
)

/*
   {
     "name": "extReference",
     "type": "ref-struct",
     "struct-ref": {
       "struct-name": "ExternalStruct",
       "is-external": true,
       "package": "lib/external/extpkg"
     },
     "queryable": false
   },
*/
var example = `{
	"name": "commons",
	"properties": {
		"folder-path": "./common"
	,"struct-name": "Commons"
	}
,"attributes": [
	{ "name": "indirizzo", "type": "struct", "struct-name": "Indirizzo", "attributes": [
	{ "name": "tipo"                   , "type": "string"}
	,{ "name": "sottoTipo"              , "type": "string"}
	,{ "name": "intestazione"           , "type": "string"}
	,{ "name": "tipoSpedizione"         , "type": "string"}
	,{ "name": "tipoSpedizioneLegale"   , "type": "string"}
	,{ "name": "civico"                 , "type": "string"}
	,{ "name": "indirizzo"              , "type": "string"}
	,{ "name": "cap"                    , "type": "string"}
	,{ "name": "localita"               , "type": "string"}
	,{ "name": "provincia"              , "type": "string"}
	,{ "name": "nazione"                , "type": "string"}
	,{ "name": "loco"                   , "type": "string"}
	,{ "name": "normalizzato"           , "type": "string"}
	,{ "name": "manuale"                , "type": "string"}
	,{ "name": "infoTecniche"           , "type": "ref-struct", "struct-name": "InfoTecniche"}
	]}

,{ "name": "sysInfo", "type": "struct", "struct-name": "SysInfo", "options": "cust-upd-handling", "attributes": [
 { "name": "dbLastUpdate"            , "type": "date" }
,{ "name": "dbUpsertedCount"         , "type": "int"}
		]}

,{ "name": "infoTecniche", "type": "struct", "struct-name": "InfoTecniche", "attributes": [
	 { "name": "dataOraAlimentazione", "type": "string"}
	,{ "name": "dataOra"             , "type": "string", "queryable":  true}
	,{ "name": "operatore"           , "type": "string"}
	,{ "name": "terminale"           , "type": "string"}
	,{ "name": "filiale"             , "type": "string"}
	,{ "name": "ufficio"             , "type": "string"}
	,{ "name": "dataAggiornamento"   , "type": "string"}
	,{ "name": "tipoOperazione"      , "type": "string"}
	,{ "name": "sistema"             , "type": "string"}
	,{ "name": "processo"            , "type": "string"}
	,{ "name": "certificato"         , "type": "string"}
	,{ "name": "dataOraCertificazione", "type": "string"}
	,{ "name": "dataValA"             , "type": "string"}
	,{ "name": "dataValDa"            , "type": "string", "queryable":  true}
	,{ "name": "ultimaOperazioneBusiness", "type": "string"}
	,{ "name": "job"                     , "type": "string"}
	,{ "name": "program"                 , "type": "string"}
		]}

,{ "name": "cliente", "type": "struct", "struct-name": "RifCliente", "attributes": [
	{ "name": "ndg"                 , "type": "string", "queryable":  true }
	,{ "name": "cognome"             , "type": "string" }
	,{ "name": "nome"                , "type": "string" }
	,{ "name": "codiceFiscale"         , "type": "string" }
	,{ "name": "partitaIVA"            , "type": "string" }
	,{ "name": "natura"                , "type": "string" }
	,{ "name": "tipo"                  , "type": "string" }
	,{ "name": "stato"                 , "type": "string" }
	,{ "name": "intestazioneRidotta"   , "type": "string" }
	,{ "name": "intestazionePostale"   , "type": "string" }
	,{ "name": "ragioneSociale"        , "type": "string" }
	]}

,{ "name": "legame", "type": "struct", "struct-name": "RifLegame", "attributes": [
 	{ "name": "legante", "type": "ref-struct", "struct-name": "RifCliente"}
	,{ "name": "legato", "type": "ref-struct", "struct-name": "RifCliente"}
	,{ "name": "tipoLegame"             , "type": "string", "queryable":  true }
	,{ "name": "stato"                  , "type": "string" }
	]}

,{ "name": "rapporto", "type": "struct", "struct-name": "RifRapporto", "attributes": [
	{ "name": "servizio"               , "type": "string"}
	,{ "name": "numero"                 , "type": "string"}
	,{ "name": "iban"                 , "type": "string"}
	]}
]
}
`
var example2 = `{
  "name": "author",
  "properties": {
    "folder-path": "./author",
    "struct-name": "Author"
  },
  "attributes": [
    {
      "name": "oId",
      "type": "object-id",
      "tags": [
        "json",
        "-",
        "bson",
        "_id"
      ],
      "queryable": true
    },
    {
      "name": "firstName",
      "type": "string",
      "tags": [
        "json",
        "fn",
        "bson",
        "fn"
      ],
      "queryable": true
    },
    {
      "name": "lastUpdate",
      "type": "date",
      "tags": [
        "json",
        "dt",
        "bson",
        "dt"
      ],
      "queryable": true
    },
    {
      "name": "lastName",
      "type": "string",
      "tags": [
        "json",
        "ln",
        "bson",
        "ln"
      ],
      "queryable": true
    },
    {
      "name": "age",
      "type": "int",
      "queryable": true
    },
    {
      "name": "address",
      "type": "struct",
      "tags": [
        "json",
        "addr",
        "bson",
        "addr"
      ],
      "struct-name": "Address",
      "attributes": [
        {
          "name": "city",
          "type": "string",
          "queryable": true
        },
        {
          "name": "strt",
          "type": "string",
          "queryable": true
        },
    {
      "name": "lastUpdate2",
      "type": "date",
      "tags": [
        "json",
        "dt2",
        "bson",
        "dt2"
      ],
      "queryable": true
    }
      ]
    },
    {
      "name": "shipAddress",
      "type": "ref-struct",
      "tags": [
        "json",
        "shipaddr",
        "bson",
        "shipaddr"
      ],
      "struct-ref": {
        "struct-name": "Address"
      },
      "queryable": true
    },

    {
      "name": "books",
      "type": "array",
      "item": {
        "type": "struct",
        "struct-name": "Book",
        "attributes": [
          {
            "name": "title",
            "type": "string",
            "queryable": true
          },
          {
            "name": "isbn",
            "type": "string",
            "queryable": true
          },
          {
            "name": "coAuthors",
            "type": "array",
            "item": {
              "type": "string"
            },
            "queryable": true
          }
        ]
      }
    },
    {
      "name": "businessRels",
      "type": "map",
      "item": {
        "type": "struct",
        "struct-name": "BusinessRel",
        "attributes": [
          {
            "name": "publisherId",
            "type": "string",
            "queryable": true
          },
          {
            "name": "publisherName",
            "type": "string",
            "queryable": true
          },
          {
            "name": "contracts",
            "type": "map",
            "item": {
              "type": "struct",
              "struct-name": "Contract",
              "attributes": [
                {
                  "name": "contractId",
                  "type": "string",
                  "queryable": true
                },
                {
                  "name": "contractDescr",
                  "type": "string",
                  "queryable": true
                }
              ]
            }
          }
        ]
      }
    }
  ]
}`

func TestGeneration(t *testing.T) {

	r := bytes.NewReader([]byte(example))
	schema, e := schema.ReadCollectionDefinition(system.GetLogger(), r)
	if e != nil {
		t.Fatal(e)
	}

	genCfg := config.Config{Version: "v2", TargetDirectory: ".", ResourceDirectory: "../..", FormatCode: true}
	genDriver, e := NewCodeGenCollection(schema)
	if e != nil {
		t.Error(e)
	} else {
		e = Generate(system.GetLogger(), &genCfg, genDriver)
		if e != nil {
			t.Error(e)
		}
	}
}

func TestCodeGenTree(t *testing.T) {

	r := bytes.NewReader([]byte(example))
	schema, e := schema.ReadCollectionDefinition(system.GetLogger(), r)
	if e != nil {
		t.Error(e)
	}

	genDriver, e := NewCodeGenCollection(schema)
	if e != nil {
		t.Error(e)
	}

	fmt.Println(genDriver)
}
