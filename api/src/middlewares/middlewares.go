package middlewares

import (
	"devbook-api/src/autenticacao"
	"devbook-api/src/respostas"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := autenticacao.ValidarToken(r)
		if err != nil {
			respostas.RespondeComErro(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
