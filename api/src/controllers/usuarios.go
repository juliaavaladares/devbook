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
	"strings"

	"github.com/gorilla/mux"
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

	if err = usuario.Preparar("cadastro"); err != nil {
		respostas.RespondeComErro(w, http.StatusBadRequest, err)
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
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, err := banco.Conectar()
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorio.NovoRepositorioUsuarios(db)
	usuarios, err := repositorio.BuscaUsuario(nomeOuNick)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.RespondeComJson(w, http.StatusOK, usuarios)

}

func BuscaUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usuarioId, err := strconv.ParseInt(params["usuarioId"], 10, 64)
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

	repositorio := repositorio.NovoRepositorioUsuarios(db)
	usuario, err := repositorio.BuscaUsuarioPorId(usuarioId)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.RespondeComJson(w, http.StatusOK, usuario)
}

func AtualizaUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusUnprocessableEntity, err)
		return
	}

	usuarioId, err := strconv.ParseInt(params["usuarioId"], 10, 64)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusBadRequest, err)
		return
	}

	usuarioIdNoToken, err := autenticacao.ExtraiUsuarioId(r)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioId != usuarioIdNoToken {
		respostas.RespondeComErro(w, http.StatusForbidden, errors.New("nao é possível atualizar usuario que nao seja o seu"))
		return
	}

	var usuario modelos.Usuario
	err = json.Unmarshal(corpoRequest, &usuario)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusBadRequest, err)
		return
	}

	if err = usuario.Preparar("atualizacao"); err != nil {
		respostas.RespondeComErro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorio.NovoRepositorioUsuarios(db)
	err = repositorio.AtualizaUsuario(usuarioId, usuario)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.RespondeComJson(w, http.StatusNoContent, nil)
}
func DeletaUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usuarioId, err := strconv.ParseInt(params["usuarioId"], 10, 64)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusBadRequest, nil)
		return
	}

	usuarioIdNoToken, err := autenticacao.ExtraiUsuarioId(r)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusUnauthorized, err)
		return
	}

	if usuarioId != usuarioIdNoToken {
		respostas.RespondeComErro(w, http.StatusForbidden, errors.New("nao é possível deletar usuario que nao seja o seu"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorio.NovoRepositorioUsuarios(db)
	err = repositorio.DeletaUsuario(usuarioId)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.RespondeComJson(w, http.StatusNoContent, nil)

}

func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorId, err := autenticacao.ExtraiUsuarioId(r)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusUnauthorized, err)
		return
	}

	parametros := mux.Vars(r)
	usuarioId, err := strconv.ParseInt(parametros["usuarioId"], 10, 64)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusBadRequest, err)
		return
	}

	if usuarioId == seguidorId {
		erro := errors.New("nao é possível seguir voce mesmo")
		respostas.RespondeComErro(w, http.StatusForbidden, erro)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorio.NovoRepositorioUsuarios(db)
	err = repositorio.SeguirUsuario(usuarioId, seguidorId)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.RespondeComJson(w, http.StatusNoContent, nil)

}

func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorId, err := autenticacao.ExtraiUsuarioId(r)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusUnauthorized, err)
		return
	}

	parametros := mux.Vars(r)
	usuarioId, err := strconv.ParseInt(parametros["usuarioId"], 10, 64)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusBadRequest, err)
		return
	}

	if usuarioId == seguidorId {
		erro := errors.New("nao é possível deixar de seguir voce mesmo")
		respostas.RespondeComErro(w, http.StatusForbidden, erro)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositorio.NovoRepositorioUsuarios(db)
	err = repositorio.PararDeSeguirUsuario(usuarioId, seguidorId)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.RespondeComJson(w, http.StatusNoContent, nil)

}

func BuscaSeguidores(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, err := strconv.ParseInt(parametros["usuarioId"], 10, 64)
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

	repositorio := repositorio.NovoRepositorioUsuarios(db)
	seguidores, err := repositorio.BuscaSeguidores(usuarioId)
	if err != nil {
		respostas.RespondeComErro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.RespondeComJson(w, http.StatusOK, seguidores)

}
