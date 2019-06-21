package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/flinox/api_rest_go/models"
)

// CreateUserTest test
func CreateUserTest() {
	// Criando alguns usuarios manualmente para teste
	users = append(users, models.User{ID: "1", Login: "Flinox", Password: "123456", People: &models.People{ID: "1", Cpf: "22222222222", Name: "Fernando Lino Di Tomazzo Silva"}})
}

// TestGetAllUsers Get all users
func TestGetAllUsers(t *testing.T) {

	CreateUserTest()

	req, err := http.NewRequest("GET", "/v1/user", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllUsers)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"id":"1","login":"Flinox","password":"123456","people":{"id":"1","cpf":"22222222222","name":"Fernando Lino Di Tomazzo Silva"}}]`

	if strings.TrimSuffix(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// TestGetUser Get specific user by id
func TestGetUser(t *testing.T) {

	CreateUserTest()

	// payload := []byte({"name":"test product","price":11.22})
	// req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(payload))
	// response := executeRequest(req)

	req, err := http.NewRequest("GET", "/v1/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUser)
	handler.ServeHTTP(resp, req)

	if status := resp.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":"1","login":"Flinox","password":"123456","people":{"id":"1","cpf":"22222222222","name":"Fernando Lino Di Tomazzo Silva"}}`

	if strings.TrimSuffix(resp.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp.Body.String(), expected)
	}

}

// // CreateUser Create a user
// func TestCreateUser(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var user models.User
// 	_ = json.NewDecoder(r.Body).Decode(&user)
// 	user.ID = params["id"]
// 	users = append(users, user)
// 	json.NewEncoder(w).Encode(user)

// }

// // DeleteUser Delete user
// func TestDeleteUser(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	for index, user := range users {
// 		if user.ID == params["id"] {
// 			users = append(users[:index], users[index+1:]...)
// 			break
// 		}
// 		json.NewEncoder(w).Encode(users)

// 	}
// }
