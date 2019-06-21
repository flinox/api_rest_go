package main

import (
	"log"
	"net/http"
	"os"

	"github.com/flinox/api_rest_go/handlers"
)

var (
	port = os.Getenv("PORTA")
)

func init() {

	if port == "" {
		os.Setenv("PORTA", "8000")
		port = os.Getenv("PORTA")
	}

}

func main() {
	log.Println("Executando o serviço na porta", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.GetUserRoutes()))
	log.Println("Encerrand o serviço...")
}
