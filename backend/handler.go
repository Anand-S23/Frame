package main

import "net/http"

func checkHealth(w http.ResponseWriter, r *http.Request) error {
    return writeJSON(w, http.StatusOK, "Ok")
}
