package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/nallanos/fire/internal/errdefs"
)

func AnswerJson(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DecodeJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func BadRequest(w http.ResponseWriter, err error) {
	AnswerJson(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
}

func InternalServerError(w http.ResponseWriter, err error) {
	AnswerJson(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
}

func UnauthorizedError(w http.ResponseWriter, err error) {
	AnswerJson(w, map[string]string{"error": err.Error()}, http.StatusUnauthorized)
}

func Error(w http.ResponseWriter, err error) {
	var httpError errdefs.Error
	if errors.As(err, &httpError) {
		AnswerJson(w, map[string]string{"error": httpError.Error()}, httpError.HTTPCode())
		return
	}
}
