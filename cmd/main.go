package main

import (
	"github.com/dylanbatar/simple-login-system/internal/infra/factory/api"
	"log"
)

func main() {
	// TODO: Test reset password usecase
	server, err := api.NewApiFactory("echo")

	if err != nil {
		log.Fatalln(err)
	}

	server.RunServer()
}
