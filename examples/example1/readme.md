### author

This example is the example0 version but with the find and update operations carried out with the generated
code resulting from the definition of the collection.

#### Collection schema definition.

    {
	  "name": "author",
	  "properties": {
	    "folder-path": "./example1",
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
	          "name": "street",
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
	  ]
	}

