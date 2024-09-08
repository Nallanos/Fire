package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nallanos/fire/internal/db"
)

type userKeyType struct{}

var UserKey userKeyType = userKeyType{}

func (a *API) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session, err := a.users.GetSession(context.Background(), r.Header.Get("Authorization"))
		if err != nil {
			UnauthorizedError(w, err)
			return
		}

		user, err := a.users.GetUserById(context.Background(), session.UserID)
		if err != nil {
			UnauthorizedError(w, err)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), UserKey, user.ID)))
	})
}

func GetUserFromCtx(ctx context.Context) (*db.User, error) {
	u := ctx.Value(UserKey)
	if u == nil {
		return nil, fmt.Errorf("user not found in context")
	}
	user, ok := u.(db.User)
	if !ok {
		return nil, fmt.Errorf("session not found in context")
	}
	return &user, nil
}
