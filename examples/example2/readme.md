### Schema author

#### Collection schema definition.

    {
	  "name": "author",
	  "properties": {
	    "folder-path": "./example2",
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
	      "queryable": true,
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
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
	      "queryable": true,
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
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
	      "queryable": true,
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "age",
	      "type": "int",
	      "queryable": true,
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "address",
	      "struct-name": "Address",
	      "type": "struct",
	      "tags": [
	        "json",
	        "addr",
	        "bson",
	        "addr"
	      ],
	      "attributes": [
	        {
	          "name": "city",
	          "type": "string",
	          "queryable": true,
	          "struct-ref": {
	            "Package": "",
	            "Item": null
	          }
	        },
	        {
	          "name": "strt",
	          "type": "string",
	          "queryable": true,
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
	    {
	      "name": "shipAddress",
	      "type": "ref-struct",
	      "tags": [
	        "json",
	        "shipaddr",
	        "bson",
	        "shipaddr"
	      ],
	      "queryable": true,
	      "struct-ref": {
	        "struct-name": "Address",
	        "Package": "",
	        "Item": {
	          "name": "address",
	          "struct-name": "Address",
	          "type": "struct",
	          "tags": [
	            "json",
	            "addr",
	            "bson",
	            "addr"
	          ],
	          "attributes": [
	            {
	              "name": "city",
	              "type": "string",
	              "queryable": true,
	              "struct-ref": {
	                "Package": "",
	                "Item": null
	              }
	            },
	            {
	              "name": "strt",
	              "type": "string",
	              "queryable": true,
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
	        }
	      }
	    },
	    {
	      "name": "books",
	      "type": "array",
	      "item": {
	        "name": "[]",
	        "struct-name": "Book",
	        "type": "struct",
	        "attributes": [
	          {
	            "name": "title",
	            "type": "string",
	            "queryable": true,
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "isbn",
	            "type": "string",
	            "queryable": true,
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "coAuthors",
	            "type": "array",
	            "item": {
	              "name": "[]",
	              "type": "string",
	              "struct-ref": {
	                "Package": "",
	                "Item": null
	              }
	            },
	            "queryable": true,
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
	      "name": "businessRels",
	      "type": "map",
	      "item": {
	        "name": "%s",
	        "struct-name": "BusinessRel",
	        "type": "struct",
	        "attributes": [
	          {
	            "name": "publisherId",
	            "type": "string",
	            "queryable": true,
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "publisherName",
	            "type": "string",
	            "queryable": true,
	            "struct-ref": {
	              "Package": "",
	              "Item": null
	            }
	          },
	          {
	            "name": "contracts",
	            "type": "map",
	            "item": {
	              "name": "%s",
	              "struct-name": "Contract",
	              "type": "struct",
	              "attributes": [
	                {
	                  "name": "contractId",
	                  "type": "string",
	                  "queryable": true,
	                  "struct-ref": {
	                    "Package": "",
	                    "Item": null
	                  }
	                },
	                {
	                  "name": "contractDescr",
	                  "type": "string",
	                  "queryable": true,
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
	    }
	  ]
	}

