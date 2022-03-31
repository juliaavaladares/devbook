package rotas

import (
	"devbook-api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		Uri:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriaUsuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscaUsuarios,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscaUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizaUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletaUsuario,
		RequerAutenticacao: true,
	},
}
