package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Anand-S23/frame/database"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

var DB *database.MongoDB

func main() {
    port := os.Getenv("Port")
    if port == "" {
        port = "8080"
    }

    mongoURI := os.Getenv("MONGODB_URI")
    if mongoURI == "" {
        log.Fatal("MONGODB_URI variable must be specified in env file")
    }

    dbName := os.Getenv("MONGODB_NAME")
    if dbName == "" {
        log.Fatal("MONGODB_NAME variable must be specified in env file")
    }

    db, err := database.InitDB(mongoURI, dbName, []string{"test", "users"}) // TODO: test is temporary need to remove
    if err != nil {
        log.Fatal(err)
    }
    DB = db

    router := mux.NewRouter()
	router.HandleFunc("/health_check", HandleFunc(checkHealth))

	log.Println("Frame running on port: ", port)
    http.ListenAndServe(":" + port, router)
}


