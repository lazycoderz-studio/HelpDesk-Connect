package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type User struct {
}

func (c *User) Register(r chi.Router) {
}

func (c *User) CreateUser(w http.ResponseWriter, r *http.Request) {
}

func (c *User) GetUser(w http.ResponseWriter, r *http.Request) {
}

func (c *User) UpdateUser(w http.ResponseWriter, r *http.Request) {
}

func (c *User) DeleteUser(w http.ResponseWriter, r *http.Request) {
}

func (c *User) GetUsers(w http.ResponseWriter, r *http.Request) {
}
