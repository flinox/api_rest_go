package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/flinox/api_rest_go/routes"
	"github.com/flinox/api_rest_go/utils"
)

var (
	port          = os.Getenv("PORTA")
	servicelog, _ = strconv.ParseBool(os.Getenv("LOG"))
)

func init() {

	// Auxiliar para gerar novos arquivos de API
	utils.WriteHandlers("newhandler")

	if port == "" {
		os.Setenv("PORTA", "8000")
		port = os.Getenv("PORTA")
	}

}

func main() {

	log.Println("[START] Executando o serviço na porta", port)
	log.Println(" [LOG]", servicelog)

	var channel = make(chan os.Signal)
	signal.Notify(channel, syscall.SIGTERM)
	signal.Notify(channel, syscall.SIGINT)
	go func() {
		sig := <-channel
		log.Println(" sig:", sig)
		log.Println(" Aguarde enquanto o serviço está sendo finalizado...")
		time.Sleep(2 * time.Second)
		log.Println("[STOP] Encerrando o serviço na porta", port)
		os.Exit(0)
	}()

	log.Fatal(http.ListenAndServe(":"+port, routes.GetUserRoutes()))

}
