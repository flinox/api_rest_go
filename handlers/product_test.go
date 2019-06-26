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
func CreateProductTest() {

	// Criando alguns usuarios manualmente para teste
	products = append(products, models.Product{ID: "1", Name: "Produto 1", Category: &models.Category{ID: "1", Name: "Eletronicos"}, User: &models.User{ID: "1", Login: "Flinox", Password: "123456", People: &models.People{ID: "1", Cpf: "22222222222", Name: "Fernando Lino Di Tomazzo Silva"}}})
}

// TestGetAllProducts Test get all products
func TestGetAllProducts(t *testing.T) {

	CreateUserTest()

	req, err := http.NewRequest("GET", "/v1/product", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllProducts)
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

// TestGetProduct Test get specific product by id
func TestGetProduct(t *testing.T) {

	CreateProductTest()

	req, err := http.NewRequest("GET", "/v1/product/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(GetProduct)
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

// TestCreateProduct Create a product
func TestCreateProduct(t *testing.T) {

	payload := []byte(`{"id":"2","login":"Flinox","password":"123456","people":{"id":"1","cpf":"22222222222","name":"Fernando Lino Di Tomazzo Silva"}}`)
	req, err := http.NewRequest("POST", "/v1/product", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"id": "2",
	})

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateProduct)
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

// DeleteProduct Delete product
func TestDeleteProduct(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/v1/product/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteProduct)
	handler.ServeHTTP(resp, req)

	if status := resp.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
