package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	log "github.com/gothew/l-og"
	"github.com/karchx/realword-nx/conduit"
)

// M is a generic map
type M map[string]interface{}

func readJson[T any](body io.Reader, input T) error {
	return json.NewDecoder(body).Decode(input)
}

func writeJSON(w http.ResponseWriter, code int, data interface{}) {
	jsonBytes, err := json.Marshal(data)

	if err != nil {
		log.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(jsonBytes)

	if err != nil {
		log.Error(err)
	}
}

var hmacSampleSecret = []byte("sample-secret")

func generateUserToken(user *conduit.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
	})

	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
