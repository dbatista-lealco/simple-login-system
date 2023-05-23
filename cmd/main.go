package main

import (
	"github.com/dylanbatar/simple-login-system/internal/infra/factory/api"
	"log"
)

func main() {
	server, err := api.NewApiFactory("echo")

	if err != nil {
		log.Fatalln(err)
	}

	server.RunServer()
}
