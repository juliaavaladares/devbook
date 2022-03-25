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
	fmt.Println("Rodando api...")

	r := router.Gerar()
	porta := fmt.Sprintf(":%d", config.Porta)
	log.Fatal(http.ListenAndServe(porta, r))
}
