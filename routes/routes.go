package routes

import (
	"fmt"
	"net/http"

	"github.com/flinox/api_rest_go/handlers"
	"github.com/gorilla/mux"
)

// GetUserRoutes Get all users routes
func GetUserRoutes() *mux.Router {

	router := mux.NewRouter()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Server is running")
	})

	// Users routes
	router.HandleFunc("/v1/user", handlers.GetAllUsers).Methods("GET")
	router.HandleFunc("/v1/user/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/v1/user/{id}", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/v1/user/{id}", handlers.DeleteUser).Methods("DELETE")

	return router

}
