package api

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/nallanos/fire/internal/errdefs"
)

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Token string `json:"token"`
}

func (a *API) signUp(w http.ResponseWriter, r *http.Request) {
	var req SignUpRequest
	if err := DecodeJSON(r, &req); err != nil {
		Error(w, err)
		return
	}
	err := a.users.CreateUser(context.Background(), req.Name, req.Email, req.Password)
	if err != nil {
		if err == errdefs.ErrAlreadyExist {
			slog.Error(err.Error())
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		slog.Error(err.Error())
		InternalServerError(w, err)
		return
	}
	user, err := a.users.GetUserByEmail(context.Background(), req.Email)
	if err != nil {
		slog.Error(err.Error())
		InternalServerError(w, err)
		return
	}
	_, err = a.users.Login(context.Background(), user.Email, req.Password)
	if err != nil {
		slog.Error(err.Error())
		InternalServerError(w, err)
		return
	}
}
