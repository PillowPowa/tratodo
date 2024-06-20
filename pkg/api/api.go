package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func MakeHandlerFunc(fn HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			resErr, ok := err.(Error)
			if !ok {
				resErr = NewApiError(http.StatusInternalServerError, "ISE Occured")
			}
			log.Println(err.Error())
			WriteJSON(w, resErr.StatusCode, resErr)
		}
	}
}
