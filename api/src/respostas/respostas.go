package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondeComJson(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(dados)
	if err != nil {
		log.Fatal(err)
	}
}

func RespondeComErro(w http.ResponseWriter, statusCode int, erro error) {
	type err struct {
		Erro string `json:"erro"`
	}

	var mensagemErro err
	mensagemErro.Erro = erro.Error()

	RespondeComJson(w, statusCode, mensagemErro)
}
