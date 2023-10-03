package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Config struct {
    Port string
}

func LoadEnv() *Config {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Could not load env file")
    }

    cfg := &Config {
        Port: os.Getenv("PORT"),
    }

    if cfg.Port == "" {
        log.Fatal("PORT variable missing from env file")
    }

    return cfg
}

func main() {
    cfg := LoadEnv()

    router := mux.NewRouter()
	router.HandleFunc("/health_check", HandleFunc(checkHealth))

	log.Println("Frame running on port: ", cfg.Port)
    http.ListenAndServe(":" + cfg.Port, router)
}


