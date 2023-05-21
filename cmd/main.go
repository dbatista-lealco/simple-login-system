package main

import (
	"github.com/dylanbatar/simple-login-system/internal/infra/factory/apiFactory"
	"log"
)

func main() {
	server, err := apiFactory.NewApiFactory("echo")

	if err != nil {
		log.Fatalln(err)
	}

	server.RunServer()
}
