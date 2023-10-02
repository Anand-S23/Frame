package main

import "github.com/Anand-S23/frame/api"

func main() {
    server := api.NewServer(":3000")
    server.Run()
}


