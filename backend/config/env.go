package config

import (
    "errors"
    "os"
)

type EnvVars struct {
    PRODUCTION     bool
    MONGODB_URI    string
    MONGODB_NAME   string
    PORT           string
    SECRET_KEY     string
    ORIGIN_ALLOWED string
}

func LoadEnv() (*EnvVars, error) {
    envMode := os.Getenv("MODE")

    mongoURI := os.Getenv("MONGODB_URI")
    if mongoURI == "" {
        return nil, errors.New("MONGODB_URI not specified")
    }

    dbName := os.Getenv("MONGODB_NAME")
    if dbName == "" {
        return nil, errors.New("MONGODB_NAME not specified")
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    secretKey := os.Getenv("SECRET_KEY")
    if secretKey == "" {
        return nil, errors.New("SECRET_KEY not specified")
    }

    originAllowed := os.Getenv("ORIGIN_ALLOWED")
    if originAllowed == "" {
        return nil, errors.New("ORIGIN_ALLOWED not specified")
    }

    return &EnvVars {
        PRODUCTION: (envMode == "production"),
        MONGODB_URI: mongoURI,
        MONGODB_NAME: dbName,
        PORT: port,
        SECRET_KEY: secretKey,
        ORIGIN_ALLOWED: originAllowed,
    }, nil
}
