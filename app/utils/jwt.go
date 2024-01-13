package utils

import (
	"errors"
	"github.com/JamesAndresCM/fiber-api/lib"
	"github.com/dgrijalva/jwt-go"
	"github.com/subosito/gotenv"
	"os"
	"time"
)

var jwtSecret []byte

func readSecret() {

	file, err := os.Open("./.env")
	lib.Fatal(err)
	defer file.Close()

	gotenv.Load()
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
}

func GenerateJWT(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expira en 24 horas
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Método de firma no válido")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
