package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:               "/usuarios",
		Metodo:            http.MethodPost,
		Funcao:            controllers.CriarUsuario,
		RequerAtenticacao: false,
	},
	{
		URI:               "/usuarios",
		Metodo:            http.MethodGet,
		Funcao:            controllers.BuscarUsuarios,
		RequerAtenticacao: true,
	},
	{
		URI:               "/usuarios/{usuarioId}",
		Metodo:            http.MethodGet,
		Funcao:            controllers.BuscarUsuario,
		RequerAtenticacao: true,
	},
	{
		URI:               "/usuarios/{usuarioId}",
		Metodo:            http.MethodPut,
		Funcao:            controllers.AtualizarUsuario,
		RequerAtenticacao: true,
	},
	{
		URI:               "/usuarios/{usuarioId}",
		Metodo:            http.MethodDelete,
		Funcao:            controllers.DeletarUsuario,
		RequerAtenticacao: true,
	},
	{
		URI:               "/usuarios/{usuarioId}/seguir",
		Metodo:            http.MethodPost,
		Funcao:            controllers.SeguirUsuario,
		RequerAtenticacao: true,
	},
	{
		URI:               "/usuarios/{usuarioId}/parar-de-seguir",
		Metodo:            http.MethodPost,
		Funcao:            controllers.PararDeSeguir,
		RequerAtenticacao: true,
	},
	{
		URI:               "/usuarios/{usuarioId}/seguidores",
		Metodo:            http.MethodGet,
		Funcao:            controllers.BuscarSeguidores,
		RequerAtenticacao: true,
	},
	{
		URI:               "/usuarios/{usuarioId}/seguindo",
		Metodo:            http.MethodGet,
		Funcao:            controllers.BuscarSeguindo,
		RequerAtenticacao: true,
	},
	{
		URI:               "/usuarios/{usuarioId}/atualizar-senha",
		Metodo:            http.MethodPost,
		Funcao:            controllers.AtualizarSenha,
		RequerAtenticacao: true,
	},

}
