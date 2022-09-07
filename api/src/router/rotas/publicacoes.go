package rotas

import (
	"devbook-api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		Uri:                "/publicacoes",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriaPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscaPublicacoes,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscaPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizaPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletaPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}/publicacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscaPublicacoesPorUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/publicacoes/{publicacaoId}/curtit",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CurtirPublicacao,
		RequerAutenticacao: true,
	},
}
