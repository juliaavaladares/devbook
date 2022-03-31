package main

import (
	"devbook-api/src/config"
	"devbook-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	r := router.Gerar()

	porta := fmt.Sprintf(":%d", config.Porta)
	fmt.Printf("Escutando na porta %s", porta)
	log.Fatal(http.ListenAndServe(porta, r))
}
