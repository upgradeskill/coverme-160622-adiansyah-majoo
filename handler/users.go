package handler

import (
	"encoding/json"
	"go-basic-server/db"
	"go-basic-server/models"
	"go-basic-server/utils"
	"net/http"
)

// Get Users handler
func GetUsers(res http.ResponseWriter, req *http.Request) {
	var users []models.User

	for _, user := range db.Userdb {
		users = append(users, user)
	}

	// parse the user data into json format
	userJSON, err := json.Marshal(&users)
	if err != nil {
		// Add the response return message
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Error parsing the user data",
		}`)

		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
		return
	}

	utils.ReturnJsonResponse(res, http.StatusOK, userJSON)
}

func GetUser(res http.ResponseWriter, req *http.Request, id string) {

	user, ok := db.Userdb[id]
	if !ok {
		// Add the response return message
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Requested user not found",
		}`)

		utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
		return
	}

	// parse the user data into json format
	userJSON, err := json.Marshal(&user)
	if err != nil {
		// Add the response return message
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Error parsing the user data",
		}`)

		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
		return
	}

	utils.ReturnJsonResponse(res, http.StatusOK, userJSON)
}

func CreateUser(res http.ResponseWriter, req *http.Request) {

	var user models.User

	payload := req.Body

	defer req.Body.Close()
	// parse the user data into json format
	err := json.NewDecoder(payload).Decode(&user)
	if err != nil {
		// Add the response return message
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Error parsing the user data",
		}`)

		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
		return
	}

	db.Userdb[user.ID] = user
	// Add the response return message
	HandlerMessage := []byte(`{
		"success": true,
		"message": "User was successfully created",
	}`)

	utils.ReturnJsonResponse(res, http.StatusCreated, HandlerMessage)
}

func UpdateUser(res http.ResponseWriter, req *http.Request, id string) {
	user, ok := db.Userdb[id]
	if !ok {
		// Add the response return message
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Requested user not found",
		}`)

		utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
		return
	}

	payload := req.Body

	defer req.Body.Close()
	// parse the user data into json format
	err := json.NewDecoder(payload).Decode(&user)
	if err != nil {
		// Add the response return message
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Error parsing the user data",
		}`)

		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
		return
	}

	db.Userdb[user.ID] = user
	// Add the response return message
	HandlerMessage := []byte(`{
		"success": true,
		"message": "User was successfully updated",
	}`)

	utils.ReturnJsonResponse(res, http.StatusOK, HandlerMessage)
}

func DeleteUser(res http.ResponseWriter, req *http.Request, id string) {
	findUser, ok := db.Userdb[id]
	if !ok {
		// Add the response return message
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Requested user not found",
		}`)

		utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
		return
	}

	var temp = make(map[string]models.User)
	for _, user := range db.Userdb {
		if user != findUser {
			temp[user.ID] = user
		}
	}

	db.Userdb = make(map[string]models.User)
	db.Userdb = temp

	// Add the response return message
	HandlerMessage := []byte(`{
		"success": true,
		"message": "User was successfully deleted",
	}`)
	utils.ReturnJsonResponse(res, http.StatusOK, HandlerMessage)
}
