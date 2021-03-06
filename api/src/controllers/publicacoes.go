package controllers

import (
	"devbook-api/src/autenticacao"
	"devbook-api/src/banco"
	"devbook-api/src/modelos"
	"devbook-api/src/repositorio"
	"devbook-api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CriaPublicacao(w http.ResponseWriter, r *http.Request) {
	var publicacao modelos.Publicacao

	usuarioId, err := autenticacao.ExtraiUsuarioId(r)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusUnauthorized, err)
		return
	}

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = json.Unmarshal(corpoRequisicao, &publicacao)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}

	publicacao.AutorId = usuarioId
	err = publicacao.Preparar()
	if err != nil {
		if err != nil {
			respostas.RespondeComErro(w, http.StatusBadRequest, err)
			return
		}
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	respositorio := repositorio.NovoRepositorioPublicacoes(db)
	publicacao.Id, err = respositorio.CriaPublicacao(publicacao)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.RespondeComJson(w, http.StatusCreated, publicacao)

}
func BuscaPublicacoes(w http.ResponseWriter, r *http.Request) {

}
func BuscaPublicacao(w http.ResponseWriter, r *http.Request) {

}
func AtualizaPublicacao(w http.ResponseWriter, r *http.Request) {

}
func DeletaPublicacao(w http.ResponseWriter, r *http.Request) {

}
