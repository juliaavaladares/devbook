package controllers

import (
	"devbook-api/src/banco"
	"devbook-api/src/modelos"
	"devbook-api/src/repositorio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CriaUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var usuario modelos.Usuario
	err = json.Unmarshal(corpoRequest, &usuario)
	if err != nil {
		log.Fatal(err)
	}

	db, err := banco.Conectar()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repositorio := repositorio.NovoRepositorioUsuarios(db)
	usuarioId, err := repositorio.CriaUsuario(usuario)
	if err != nil {
		log.Fatal(err)
	}

	mensagem := fmt.Sprintf("Ultimo id inserido: %d", usuarioId)
	w.Write([]byte(mensagem))
}

func BuscaUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários!"))
}

func BuscaUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Busca usuario especifico!"))
}

func AtualizaUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualiza usuario especificoo!"))
}
func DeletaUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleta usuarios usuario!"))
}
