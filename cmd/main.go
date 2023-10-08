package main

import (
	"fmt"
	"log"

	"github.com/dylanbatar/simple-login-system/internal/infra/factory/api"
)

func main() {
	fmt.Println("test commit 3")

	// TODO: Test reset password usecase
	server, err := api.NewApiFactory("echo")

	if err != nil {
		log.Fatalln(err)
	}

	server.RunServer()
}
