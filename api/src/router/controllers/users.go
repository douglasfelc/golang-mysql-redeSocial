package controllers

import "net/http"

// CreateUser insert a user...
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário!"))
}

// GetUsers ...
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos usuários!"))
}

// GetUser ...
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário!"))
}

// UpdateUser ...
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário!"))
}

// DeleteUser ...
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário!"))
}
