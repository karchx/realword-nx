package server

import (
	"encoding/json"
	"io"
	"net/http"

	log "github.com/gothew/l-og"
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
