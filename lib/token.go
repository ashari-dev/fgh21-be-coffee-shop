package lib

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateUserTokenById(id int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"iat": jwt.NumericDate{
			Time: time.Now(),
		},
	})
	tokenSignedString, _ := token.SignedString([]byte("secret"))

	return tokenSignedString
}

func ValidateToken(token string) (bool, int) {
	validated, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if err != nil {
		panic("Error: token invalid")
	}

	if claims, ok := validated.Claims.(jwt.MapClaims); ok {
		userId := int(claims["id"].(float64))

		return true, userId
	} else {
		panic("Error: token invalid")
	}
}
