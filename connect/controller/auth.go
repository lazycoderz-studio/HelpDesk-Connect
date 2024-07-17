package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Auth struct{}

func (c *Auth) Register(r chi.Router) {
}

func (c *Auth) Login(w http.ResponseWriter, r *http.Request) {
}

func (c *Auth) RefreshToken(w http.ResponseWriter, r *http.Request) {
}

func (c *Auth) Logout(w http.ResponseWriter, r *http.Request) {
}

func (c *Auth) Me(w http.ResponseWriter, r *http.Request) {
}

func (c *Auth) ChangePassword(w http.ResponseWriter, r *http.Request) {
}

func (c *Auth) ForgotPassword(w http.ResponseWriter, r *http.Request) {
}

func (c *Auth) ResetPassword(w http.ResponseWriter, r *http.Request) {
}
