package controller

import (
	"log"
	"net/http"

	"github.com/Anand-S23/frame/internal/storage"
)

type Controller struct {
    store *storage.MongoStore
}

func NewController(store *storage.MongoStore) *Controller {
    return &Controller{store: store}
}

func (c *Controller) CheckHealth(w http.ResponseWriter, r *http.Request) error {
    log.Println("Hit")
    return WriteJSON(w, http.StatusOK, "Ok")
}

