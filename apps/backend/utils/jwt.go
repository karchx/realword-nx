package utils

import (
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
)

var secretString = []byte("!!SECRET!!")
var JWTSecret = jwtware.SigningKey{Key: secretString}

func GenerateJWT(id uuid.UUID) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id.String()
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString(secretString)
	return t
}
