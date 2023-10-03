package main

import (
	"net/http"
    "log"
	"github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
	router.HandleFunc("/health_check", HandleFunc(checkHealth))

    // TODO: Read value from .env file
    port := "8080"
	log.Println("Frame running on port: ", port)
    http.ListenAndServe(":" + port, router)
}


