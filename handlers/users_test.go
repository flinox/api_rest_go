package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/flinox/api_rest_go/models"
	"github.com/gorilla/mux"
)

// CreateUserTest Create a user test
func CreateUserTest() {

	// Criando alguns usuarios manualmente para teste
	users = append(users, models.User{ID: "1", Login: "Flinox", Password: "123456", People: &models.People{ID: "1", Cpf: "22222222222", Name: "Fernando Lino Di Tomazzo Silva"}})
}

// TestGetAllUsers Test get all users
func TestGetAllUsers(t *testing.T) {

	CreateUserTest()

	req, err := http.NewRequest("GET", "/v1/user", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllUsers)
	handler.ServeHTTP(resp, req)

	if status := resp.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"id":"1","login":"Flinox","password":"123456","people":{"id":"1","cpf":"22222222222","name":"Fernando Lino Di Tomazzo Silva"}}]`

	if strings.TrimSuffix(resp.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp.Body.String(), expected)
	}
}

// TestGetUser Test get specific user by id
func TestGetUser(t *testing.T) {

	CreateUserTest()

	req, err := http.NewRequest("GET", "/v1/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

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

// TestCreateUser Create a user
func TestCreateUser(t *testing.T) {

	payload := []byte(`{"id":"2","login":"Flinox","password":"123456","people":{"id":"1","cpf":"22222222222","name":"Fernando Lino Di Tomazzo Silva"}}`)
	req, err := http.NewRequest("POST", "/v1/user", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"id": "2",
	})

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(resp, req)

	if status := resp.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":"2","login":"Flinox","password":"123456","people":{"id":"1","cpf":"22222222222","name":"Fernando Lino Di Tomazzo Silva"}}`

	if strings.TrimSuffix(resp.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp.Body.String(), expected)
	}

}

// DeleteUser Delete user
func TestDeleteUser(t *testing.T) {

	//CreateUserTest()

	req, err := http.NewRequest("DELETE", "/v1/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUser)
	handler.ServeHTTP(resp, req)

	if status := resp.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
