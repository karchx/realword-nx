package server

import "net/http"

func invalidAuthTokenError(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", "Token")
	msg := "invalid or missiong authentication token"
	errorReponse(w, http.StatusUnauthorized, msg)
}

func errorReponse[T any](w http.ResponseWriter, code int, errs T) {
	writeJSON(w, code, M{"errors": errs})
}