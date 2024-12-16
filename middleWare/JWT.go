package middleWare

import (
	"MPT-CS/models"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strconv"
	"time"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWT(user models.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tokenString string) (bool, string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return privateKey, nil
	})

	if err != nil {
		return false, ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, fmt.Sprintf("%v", claims["id"])
	} else {
		return false, ""
	}
}
