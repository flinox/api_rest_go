package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/flinox/api_rest_go/models"
	"github.com/gorilla/mux"
)

var (
	users []models.User
)

// GetUserRoutes Get all users routes
func GetUserRoutes() *mux.Router {

	users = append(users, models.User{ID: "1", Login: "Flinox", Password: "123456", People: &models.People{ID: "1", Cpf: "22222222222", Name: "Fernando Lino Di Tomazzo Silva"}})

	router := mux.NewRouter()

	router.HandleFunc("/v1/user", GetAllUsers).Methods("GET")
	router.HandleFunc("/v1/user/{id}", GetUser).Methods("GET")
	router.HandleFunc("/v1/user/{id}", CreateUser).Methods("POST")
	router.HandleFunc("/v1/user/{id}", DeleteUser).Methods("DELETE")

	return router

}

// GetAllUsers Get all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	defer utils.timeTrack(time.Now(), "GetAllUsers")
	json.NewEncoder(w).Encode(users)
}

// GetUser Get specific user by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	defer utils.timeTrack(time.Now(), "GetUser")
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
	defer utils.timeTrack(time.Now(), "CreateUser")
	params := mux.Vars(r)
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = params["id"]
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

// DeleteUser Delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	defer utils.timeTrack(time.Now(), "DeleteUser")
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
