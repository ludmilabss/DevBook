package models

// DadosAutenticacao contem o id e o token do usuário autenticado
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
