package api

import (
	"context"
	"log/slog"
	"net/http"
)

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (a *API) signIn(w http.ResponseWriter, r *http.Request) {
	var req SignInRequest
	if err := DecodeJSON(r, &req); err != nil {
		BadRequest(w, err)
		return
	}
	token, err := a.users.Login(context.Background(), req.Email, req.Password)
	if err != nil {
		slog.Error(err.Error())
		InternalServerError(w, err)
		return
	}
	AnswerJson(w, token, http.StatusOK)
}
