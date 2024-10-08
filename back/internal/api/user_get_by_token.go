package api

import (
	"context"
	"log/slog"
	"net/http"
)

type gettingUserRequest struct {
	Token string `json:"token"`
}

func (a *API) getUserByToken(w http.ResponseWriter, r *http.Request) {
	var req gettingUserRequest
	if err := DecodeJSON(r, &req); err != nil {
		BadRequest(w, err)
		return
	}
	session, err := a.users.GetSession(context.Background(), req.Token)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		InternalServerError(w, err)
	}
	user, err := a.users.GetUserById(context.Background(), session.UserID)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		InternalServerError(w, err)
	}
	AnswerJson(w, user, 200)
}
