# GoLab5 - sort of CRUD

**Instalation**
---
```
$ git clone git@github.com:jpauloeti/golang-crud.git
$ cd golab5
$ go run .\main.go
```


**Usage**
---
 **Insert new user:**

`POST http://localhost:5000/`

Body example:

```
{
	"name": "João",
	"address": {
		"street": "Rua do João",
		"apt": "510",
		"city": "São José do Rio Preto",
		"state": "SP",
		"zip": "15050-000"
	}
}
```

Response:

* String

`User successfully inserted. ObjectId: ObjectIdHex("58b098d56ab3283218b78fb9")`

##

**Get user by ID**

`GET http://localhost:5000/{id}`

Response:

* JSON

```
{
	"Id": "58b098d56ab3283218b78fb9",
	"Name": "João",
	"Address": {
		"Street": "Rua do João",
		"Apt": "510",
		"City": "São José do Rio Preto",
		"State": "SP",
		"Zip": "15050-000"
	}
}
```

##

**Delete user by ID**

`DELETE http://localhost:5000/{id}`

Response:

* String

`UserID 58b098d56ab3283218b78fb9 deleted successfully`
