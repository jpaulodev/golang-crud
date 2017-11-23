package controllers

import (
	"errors"
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"../models"
	"../server"
)

type UserController struct{}

func User() *UserController {
	return &UserController{}
}

func (uc UserController) Insert(receivedUser models.User) bson.ObjectId {
	receivedUser.Id = bson.NewObjectId()

	session := server.GetSession()
	conn := session.DB("golab").C("users")
	err := conn.Insert(receivedUser)

	if err != nil {
		panic(err)
	}

	fmt.Printf("User received: %s", receivedUser)
	return receivedUser.Id
}

func (uc UserController) Get(userId string) (error, models.User) {
	if !bson.IsObjectIdHex(userId) {
		return errors.New("Insert a valid ID"), models.User{}
	}

	session := server.GetSession()
	conn := session.DB("golab").C("users")

	result := models.User{}
	err := conn.Find(bson.M{"_id": bson.ObjectIdHex(userId)}).One(&result)

	if err != nil {
		return errors.New("User not found"), models.User{}
	}

	fmt.Printf("User found: %v", result)
	return nil, result
}

func (uc UserController) Delete(userId string) (error, bool) {
	if !bson.IsObjectIdHex(userId) {
		return errors.New("Insert a valid ID"), false
	}

	session := server.GetSession()
	conn := session.DB("golab").C("users")

	err := conn.Remove(bson.M{"_id": bson.ObjectIdHex(userId)})

	if err != nil {
		return err, false
	}

	return nil, true
}
