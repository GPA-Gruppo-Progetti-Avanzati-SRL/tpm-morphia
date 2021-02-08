### author

#### Collection schema definition.

    {
	  "name": "author",
	  "properties": {
	    "folderPath": "./author",
	    "struct-name": "Author"
	  },
	  "attributes": [
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
	          "queryable": true
	        },
	        {
	          "name": "strt",
	          "type": "string",
	          "queryable": true
	        }
	      ]
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
	            "queryable": true
	          },
	          {
	            "name": "isbn",
	            "type": "string",
	            "queryable": true
	          },
	          {
	            "name": "posts",
	            "type": "array",
	            "item": {
	              "name": "[]",
	              "type": "string"
	            },
	            "queryable": true
	          }
	        ]
	      }
	    }
	  ]
	}

