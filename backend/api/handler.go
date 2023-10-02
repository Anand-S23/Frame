package api

import (
    "net/http"
)

func (s *Server) healthCheck(w http.ResponseWriter, r *http.Request) error {
    return writeJSON(w, http.StatusOK, map[string]string{"message": "Frame backend is running"})
}

