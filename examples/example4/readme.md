### cliente

#### Collection schema definition.

    {
	  "name": "cliente",
	  "properties": {
	    "folder-path": "./example4",
	    "struct-name": "Cliente"
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
	      "queryable": true,
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "ndg",
	      "type": "string",
	      "queryable": true,
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "codiceFiscale",
	      "type": "string",
	      "queryable": true,
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "partitaIVA",
	      "type": "string",
	      "queryable": true,
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "natura",
	      "type": "string",
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "stato",
	      "type": "string",
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "indirizzi",
	      "type": "map",
	      "item": {
	        "name": "%s",
	        "struct-name": "Indirizzo",
	        "type": "struct",
	        "attributes": [
	          {
	            "name": "indirizzo",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "cap",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "localita",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "provincia",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "nazione",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          }
	        ],
	        "struct-ref": {
	          "Package": "",
	          "Item": null
	        }
	      },
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "legati",
	      "type": "array",
	      "item": {
	        "name": "[]",
	        "struct-name": "Legame",
	        "type": "struct",
	        "attributes": [
	          {
	            "name": "ndg",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "cognome",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "nome",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "codiceFiscale",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "partitaIVA",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "natura",
	            "type": "string",
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          }
	        ],
	        "struct-ref": {
	          "Package": "",
	          "Item": null
	        }
	      },
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "leganti",
	      "type": "array",
	      "item": {
	        "name": "[]",
	        "type": "ref-struct",
	        "struct-ref": {
	          "struct-name": "Legame",
	          "Package": "",
	          "Item": null
	        }
	      },
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    }
	  ]
	}

