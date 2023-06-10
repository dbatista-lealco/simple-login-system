package main

import (
	"encoding/json"
	"github.com/dylanbatar/simple-login-system/internal/infra/factory/api"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

type myCustomClaim struct {
	jwt.RegisteredClaims
	Foo map[string]interface{} `json:"data"`
}

func main() {
	server, err := api.NewApiFactory("echo")

	if err != nil {
		log.Fatalln(err)
	}

	server.RunServer()

}

func conver(data interface{}) map[string]interface{} {
	var maper map[string]interface{}

	a, _ := json.Marshal(data)

	json.Unmarshal(a, &maper)

	return maper
}
