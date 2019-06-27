package main

import (
	"flag"
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
	execmode      = "service"
	urnname       = ""
)

func init() {

	if port == "" {
		os.Setenv("PORTA", "8000")
		port = os.Getenv("PORTA")
	}

	flag.StringVar(&execmode, "execmode", "service", "Executar em modo service (production) ou genurn (generate urn/handlers)")
	flag.StringVar(&urnname, "urn", "", "Se modo genurn deve ser informado o nome do urn que se quer gerar")

}

func main() {

	flag.Parse()

	if execmode == "genurn" {

		if urnname == "" {
			log.Println("URN não gerada, necessário informar o parametro -urn=<nome_urn>")
			os.Exit(3)
		}

		utils.GenerateHandler(urnname)
		os.Exit(2)
	}

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
