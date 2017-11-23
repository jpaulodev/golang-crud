package main

import (
	"fmt"
	"net/http"

	"./controllers"
	"./models"

	"github.com/gotschmarcel/goserv"
)

func main() {
	server := goserv.NewServer()
	server.Post("/", createRecord)
	server.Get("/:record_id", readRecord)
	server.Delete("/:record_id", deleteRecord)

	err := server.Listen(":5000")

	if err != nil {
		fmt.Println("Error in starting rest server")
		panic(err)
	}
}

func createRecord(w http.ResponseWriter, r *http.Request) {
	var receivedUser models.User

	if err := goserv.ReadJSONBody(r, &receivedUser); err != nil {
		goserv.Context(r).Error(err, http.StatusBadRequest)
		return
	}

	userController := controllers.User()
	insertedID := userController.Insert(receivedUser)

	goserv.WriteStringf(w, "User successfully inserted. ObjectId: %s", insertedID)
}

func readRecord(w http.ResponseWriter, r *http.Request) {
	id := goserv.Context(r).Param("record_id")

	userController := controllers.User()
	err, user := userController.Get(id)

	if err != nil {
		goserv.WriteStringf(w, "%s", err)
		return
	}

	goserv.WriteJSON(w, user)
}

func deleteRecord(w http.ResponseWriter, r *http.Request) {
	id := goserv.Context(r).Param("record_id")

	userController := controllers.User()
	err, _ := userController.Delete(id)

	if err != nil {
		goserv.WriteStringf(w, "UserID %s was not deleted. Reason: %s ", id, err)
		return
	}

	goserv.WriteStringf(w, "UserID %s deleted successfully", id)
}
