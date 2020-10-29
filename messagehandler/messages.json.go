package messagehandler

// messagesJson message eror
var messagesJSON = `{
	"messages":[
		{
			"code":"E000001",
			"statusCode":500,
			"message":"Internal Server Error"
		},

		{
			"code": "W000001",
			"statusCode": 400,
			"message": "Insert %s failed."
		},
		{
			"code": "W000002",
			"statusCode": 400,
			"message": "Update %s failed."
		},
		{
			"code": "W000003",
			"statusCode": 400,
			"message": "Delete %s failed."
		},
		{
			"code": "W000004",
			"statusCode": 400,
			"message": "Get %s failed."
		},
		{
			"code": "W000005",
			"statusCode": 400,
			"message": "Binding the request failed."
		},
		{
			"code": "W000006",
			"statusCode": 400,
			"message": "Input %s invalid."
		},
		{
			"code": "W000007",
			"statusCode": 400,
			"message": "Missing required field: %s"
		},
		{
			"code": "W000008",
			"statusCode": 404,
			"message": "%s not found"
		},
		{
			"code": "W000009",
			"statusCode": 409,
			"message": "%s duplicate"
		},

		{
			"code": "I000008",
			"statusCode": 200,
			"message": "Successful."
		  }
	]
	}`
