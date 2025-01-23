package jwt

import (
	"github.com/golang-jwt/jwt"
	"kursachDB/internal/domain/models"
	"time"
)

func NewToken(user models.User, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["phone"] = user.PhoneNumber
	claims["fio"] = user.FIO
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte("is0tjg9i5f4jgjtg8j"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
