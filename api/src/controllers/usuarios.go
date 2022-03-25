package controllers

import "net/http"

func CriaUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuario!"))
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
