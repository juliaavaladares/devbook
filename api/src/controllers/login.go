package controllers

import (
	"devbook-api/src/autenticacao"
	"devbook-api/src/banco"
	"devbook-api/src/modelos"
	"devbook-api/src/repositorio"
	"devbook-api/src/respostas"
	"devbook-api/src/seguranca"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusUnprocessableEntity, nil)
		return
	}

	var usuario modelos.Usuario
	err = json.Unmarshal(corpoRequisicao, &usuario)
	if err != nil {
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
	usuarioSalvoNoBanco, err := repositorio.BuscaUsuarioPorEmail(usuario.Email)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}
	err = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := autenticacao.CriaToken(usuarioSalvoNoBanco.ID)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(token))

}
