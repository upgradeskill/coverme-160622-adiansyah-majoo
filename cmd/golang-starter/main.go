package main

import (
	"go-basic-server/db"
	"go-basic-server/handler"
	"go-basic-server/models"
	"go-basic-server/utils"
	"log"
	"net/http"
	"strings"
)

func route(w http.ResponseWriter, r *http.Request) {
	arrPath := strings.Split(r.URL.Path, "/")

	switch r.Method {
	case "GET":
		if len(arrPath) >= 3 {
			id := arrPath[2]
			handler.GetUser(w, r, id)
		} else {
			handler.GetUsers(w, r)
		}
	case "POST":
		handler.CreateUser(w, r)
	case "PUT":
		id := arrPath[2]
		handler.UpdateUser(w, r, id)
	case "DELETE":
		id := arrPath[2]
		handler.DeleteUser(w, r, id)
	default:
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Method not found",
		}`)

		utils.ReturnJsonResponse(w, http.StatusNotFound, HandlerMessage)
	}
}

func main() {
	// initialize the database
	db.Userdb["001"] = models.User{ID: "001", Name: "Adi", Address: "Malang"}
	db.Userdb["002"] = models.User{ID: "002", Name: "Fulan", Address: "Surabaya"}

	http.HandleFunc("/", route)
	http.HandleFunc("/user", route)

	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Print("The Server Running on localhost port 8080")
}
