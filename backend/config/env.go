package config

import (
    "errors"
    "os"
)

type EnvVars struct {
    MONGODB_URI  string
    MONGODB_NAME string
    PORT         string
}

func LoadEnv() (*EnvVars, error) {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    mongoURI := os.Getenv("MONGODB_URI")
    if mongoURI == "" {
        return nil, errors.New("MONGODB_URI not specified")
    }

    dbName := os.Getenv("MONGODB_NAME")
    if dbName == "" {
        return nil, errors.New("MONGODB_NAME not specified")
    }

    return &EnvVars {
        MONGODB_URI: mongoURI,
        MONGODB_NAME: dbName,
        PORT: port,
    }, nil
}
