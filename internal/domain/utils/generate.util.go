package utils

import (
	"math/rand"
	"time"
)

func generateLetters() []string {
	var lettersSlice []string

	for l := 'a'; l <= 'z'; l++ {
		lettersSlice = append(lettersSlice, string(l))
	}

	for l := 'A'; l <= 'Z'; l++ {
		lettersSlice = append(lettersSlice, string(l))
	}

	return lettersSlice
}

func GenerateVerifyCode(codeLen int) string {
	letters := generateLetters()
	var code string

	for i := 0; i < codeLen; i++ {
		seed := time.Now().UnixMilli()
		rand.NewSource(seed)

		pick := rand.Intn((52 - 1) + 1)
		code += letters[pick]
	}

	return code
}
