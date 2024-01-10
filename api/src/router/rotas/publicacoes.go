package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI:    "/publicacoes",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarPublicacao,
		RequerAtenticacao: true,
	},
	{
		URI:    "/publicacoes",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarPublicacoes,
		RequerAtenticacao: true,
	},
	{
		URI:    "/publicacoes/{publicacaoId}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarPublicacao,
		RequerAtenticacao: true,
	},
	{
		URI:    "/publicacoes/{publicacaoId}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarPublicacao,
		RequerAtenticacao: true,
	},
	{
		URI:    "/publicacoes/{publicacaoId}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarPublicacao,
		RequerAtenticacao: true,
	},
}