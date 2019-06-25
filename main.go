package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/flinox/api_rest_go/handlers"
)

var (
	port       = os.Getenv("PORTA")
	servicelog = true
)

func init() {

	if port == "" {
		os.Setenv("PORTA", "8000")
		port = os.Getenv("PORTA")
	}

	if os.Getenv("LOG") == "" {
		os.Setenv("LOG", strconv.FormatBool(servicelog))
	}

	// servicelog, err := strconv.ParseBool(os.Getenv("LOG"))

	// if err == nil {
	// 	/** displayg the type of the b variable */
	// 	fmt.Printf("Type: %T \n", servicelog)

	// 	/** displaying the string variable into the console */
	// 	fmt.Println("Value:", servicelog)
	// }

}

func main() {
	log.Println("[START] Executando o serviço na porta", port)
	log.Println("Log está ativo?", servicelog)
	log.Fatal(http.ListenAndServe(":"+port, handlers.GetUserRoutes()))
	log.Println("[STOP] Encerrand o serviço na porta", port)
}
