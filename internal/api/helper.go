package api

import (
	"encoding/json"
	"net/http"
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