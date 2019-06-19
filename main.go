package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	port = os.Getenv("PORTA")
)

type User struct {
	ID       string  `json:"id,omitempty"`
	Login    string  `json:"login,omitempty"`
	Password string  `json:"password,omitempty"`
	People   *People `json:"people,omitempty"`
}
type People struct {
	ID   string `json:"id,omitempty"`
	Cpf  string `json:"cpf,omitempty"`
	Name string `json:"name,omitempty"`
}

var users []User

func init() {

	if port == "" {
		os.Setenv("PORTA", "8000")
		port = os.Getenv("PORTA")
	}

	// Criando alguns usuarios manualmente para teste
	users = append(users, User{ID: "1", Login: "Flinox", Password: "123456", People: &People{ID: "1", Cpf: "22222222222", Name: "Fernando Lino Di Tomazzo Silva"}})

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/user", GetAllUsers).Methods("GET")
	router.HandleFunc("/user/{id}", GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, user := range users {
		if user.ID == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = params["id"]
	users = append(users, user)
	json.NewEncoder(w).Encode(user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, user := range users {
		if user.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(users)

	}
}
