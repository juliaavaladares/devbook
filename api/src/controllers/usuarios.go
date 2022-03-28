package controllers

import (
	"devbook-api/src/banco"
	"devbook-api/src/modelos"
	"devbook-api/src/repositorio"
	"devbook-api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CriaUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario modelos.Usuario
	err = json.Unmarshal(corpoRequest, &usuario)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusBadRequest, err)
		return
	}

	if err = usuario.Preparar(); err != nil {
		respostas.RespondeComErro(w, http.StatusBadRequest, nil)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorio.NovoRepositorioUsuarios(db)
	usuario.ID, err = repositorio.CriaUsuario(usuario)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.RespondeComJson(w, http.StatusCreated, usuario)
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
