package main

import (
    "encoding/json" 
    "net/http"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

func HandleFunc(fn apiFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        err := fn(w, r)
        if err != nil {
            errMsg := map[string]string { "error": err.Error() }
            writeJSON(w, http.StatusInternalServerError, errMsg)
        }
    }
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

