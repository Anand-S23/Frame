package controller

import (
	"net/http"

	"github.com/Anand-S23/frame/internal/storage"
)

type Controller struct {
    store *storage.MongoStore
}

func NewController(store *storage.MongoStore) *Controller {
    return &Controller{store: store}
}

func (c *Controller) Ping(w http.ResponseWriter, r *http.Request) error {
    return WriteJSON(w, http.StatusOK, "Pong")
}

