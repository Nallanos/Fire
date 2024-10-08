package api

import (
	"context"
	"log/slog"
	"net/http"
)

func (a *API) AuthWebSocketMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := a.users.GetSession(r.Context(), r.URL.Query().Get("token"))
		if err != nil {
			UnauthorizedError(w, err)
			return
		}

		user, err := a.users.GetUserById(context.Background(), session.UserID)
		if err != nil {
			slog.Error(err.Error())
			UnauthorizedError(w, err)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), UserKey, user.ID)))
	})
}
