### Schema session

#### Collection schema definition.

    {
	  "name": "session",
	  "properties": {
	    "folder-path": "./example5",
	    "struct-name": "Session"
	  },
	  "attributes": [
	    {
	      "name": "sid",
	      "type": "string",
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
	      "name": "nickname",
	      "type": "string",
	      "queryable": true,
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "remoteaddr",
	      "type": "string",
	      "queryable": true,
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "flags",
	      "type": "string",
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    }
	  ]
	}

