package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/dylanbatar/simple-login-system/internal/application/ports"
	jwtLib "github.com/golang-jwt/jwt/v5"
)

type JwtAdapter struct {
}

type CustomClaims struct {
	jwtLib.RegisteredClaims
	Data map[string]interface{} `json:"data"`
}

func NewJwtAdapter() *JwtAdapter {
	return &JwtAdapter{}
}

func (adapter *JwtAdapter) Generate(data interface{}) ports.Token {
	var dataClaims map[string]interface{}
	dataBytes, _ := json.Marshal(data)
	json.Unmarshal(dataBytes, &dataClaims)

	token := jwtLib.NewWithClaims(jwtLib.SigningMethodHS256, CustomClaims{
		Data: dataClaims,
	})

	jwtString, _ := token.SignedString([]byte("secret"))

	return ports.Token(jwtString)
}

func (adapter *JwtAdapter) Verify(tokenString ports.Token) (interface{}, error) {
	tokenDecode, err := jwtLib.Parse(string(tokenString), func(token *jwtLib.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtLib.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid method to sign")
		}

		return []byte("secret"), nil
	})

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	claims, _ := tokenDecode.Claims.(jwtLib.MapClaims)

	return claims, nil
}
