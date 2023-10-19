package controller

import (
	"net/http"

	"github.com/Anand-S23/frame/internal/storage"
)

type Controller struct {
    store         *storage.MongoStore
    jwtSecretKey string
}

func NewController(store *storage.MongoStore, secretKey string) *Controller {
    return &Controller {
        store: store,
        jwtSecretKey: secretKey,
    }
}

func (c *Controller) Ping(w http.ResponseWriter, r *http.Request) error {
    return WriteJSON(w, http.StatusOK, "Pong")
}

