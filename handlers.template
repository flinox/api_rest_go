package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/flinox/api_rest_go/models"
	"github.com/flinox/api_rest_go/utils"
	"github.com/gorilla/mux"
)

var (
	users []models.User
)

// GetAllUsers Get all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	defer utils.TimeTrack(time.Now(), "GetAllUsers")
	json.NewEncoder(w).Encode(users)
}

// GetUser Get specific user by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	defer utils.TimeTrack(time.Now(), "GetUser")
	params := mux.Vars(r)
	for _, user := range users {
		if user.ID == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.User{})
}

// CreateUser Create a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer utils.TimeTrack(time.Now(), "CreateUser")
	params := mux.Vars(r)
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = params["id"]
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

// DeleteUser Delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	defer utils.TimeTrack(time.Now(), "DeleteUser")
	params := mux.Vars(r)
	for index, user := range users {
		if user.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			w.WriteHeader(200)
			break
		}
	}
	w.WriteHeader(404)
}
