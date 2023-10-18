package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Anand-S23/frame/config"
	"github.com/Anand-S23/frame/internal/controller"
	"github.com/Anand-S23/frame/internal/database"
	"github.com/Anand-S23/frame/internal/router"
	"github.com/Anand-S23/frame/internal/storage"
)

func main() {
    env, err := config.LoadEnv()
    if err != nil {
        log.Fatal(err)
    }

    db, err := database.InitDB(env.MONGODB_URI, env.MONGODB_NAME, 10 * time.Second)
    if err != nil {
        log.Fatal(err)
    }
    defer database.CloseDB(db)

    mongoStore := storage.NewMongoStore(db)
    controller := controller.NewController(mongoStore)
    router := router.NewRouter(controller)

	log.Println("Frame running on port: ", env.PORT)
    http.ListenAndServe(":" + env.PORT, router)
}

