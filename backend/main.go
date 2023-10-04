package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
    Port string
    MongoURI string
}

func LoadEnv() *Config {
    port := os.Getenv("Port")
    if port == "" {
        port = "8080"
    }

    mongoURI := os.Getenv("MONGODB_URI")
    if mongoURI == "" {
        log.Fatal("MONGODB_URI variable must be specified in env file")
    }

    return &Config {
        Port: port,
        MongoURI: mongoURI,
    }
}

func main() {
    cfg := LoadEnv()

    opts := options.Client().ApplyURI(cfg.MongoURI)
    client, err := mongo.Connect(context.TODO(), opts)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }

    router := mux.NewRouter()
	router.HandleFunc("/health_check", HandleFunc(checkHealth))

	log.Println("Frame running on port: ", cfg.Port)
    http.ListenAndServe(":" + cfg.Port, router)
}


