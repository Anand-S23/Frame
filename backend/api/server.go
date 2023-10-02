package api

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
    return &Server {
        listenAddr: listenAddr,
    }
}

func (s *Server) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/api/health_check", makeAPIHandleFunc(s.healthCheck))

	log.Println("Frame running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

