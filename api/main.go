package main

import (
	"devbook-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando api...")

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":5000", r))
}
