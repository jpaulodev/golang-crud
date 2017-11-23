package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id      bson.ObjectId `bson:"_id"`
	Name    string        `bson:"name"`
	Address Address       `bson:"address"`
}

type Address struct {
	Street string `bson:"street"`
	Apt    string `bson:"apt"`
	City   string `bson:"city"`
	State  string `bson:"state"`
	Zip    string `bson:"zip"`
}
