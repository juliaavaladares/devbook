package controllers

import (
	"devbook-api/src/autenticacao"
	"devbook-api/src/banco"
	"devbook-api/src/modelos"
	"devbook-api/src/repositorio"
	"devbook-api/src/respostas"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	var publicacoes []modelos.Publicacao

	usuarioId, err := autenticacao.ExtraiUsuarioId(r)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusUnauthorized, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	respositorio := repositorio.NovoRepositorioPublicacoes(db)
	publicacoes, err = respositorio.BuscaPublicacoes(usuarioId)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.RespondeComJson(w, http.StatusCreated, publicacoes)
}
func BuscaPublicacao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicacaoId, err := strconv.ParseInt(params["publicacaoId"], 10, 64)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorio.NovoRepositorioPublicacoes(db)
	usuario, err := repositorio.BuscaPublicacao(publicacaoId)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.RespondeComJson(w, http.StatusOK, usuario)

}
func AtualizaPublicacao(w http.ResponseWriter, r *http.Request) {
	var publicacao modelos.Publicacao

	usuarioId, err := autenticacao.ExtraiUsuarioId(r)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	publicacaoId, err := strconv.ParseInt(params["publicacaoId"], 10, 64)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorio.NovoRepositorioPublicacoes(db)
	publicacaoSalvaNoBanco, err := repositorio.BuscaPublicacao(publicacaoId)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}

	if publicacaoSalvaNoBanco.AutorId != usuarioId {
		respostas.RespondeComErro(w, http.StatusForbidden, errors.New("nao é possível atualizar publicacao que nao seja sua"))
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
	err = publicacao.Preparar()
	if err != nil {
		respostas.RespondeComErro(w, http.StatusBadRequest, err)
		return
	}

	err = repositorio.AtualizaPublicacao(publicacaoId, publicacao)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.RespondeComJson(w, http.StatusNoContent, nil)

}
func DeletaPublicacao(w http.ResponseWriter, r *http.Request) {

}
