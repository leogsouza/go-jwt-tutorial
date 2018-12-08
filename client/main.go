package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("MySecredtKey")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Leonardo Souza"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())

		return "", err
	}

	return tokenString, nil

}

func main() {
	fmt.Println("My Simple Client")

	tokenString, err := GenerateJWT()

	if err != nil {
		fmt.Println("error generating token string")
	}

	fmt.Println(tokenString)
}
