package server

import (
	"os"

	"github.com/rs/cors"
)

const (
	versionApi   = "v1"
	OptionalAuth = false
)

func (s *Server) routes() {
	s.router.Use(cors.AllowAll().Handler)
	s.router.Use(Logger(os.Stdout))

	apiRouter := s.router.PathPrefix("/api/" + versionApi).Subrouter()
	optionalAuth := apiRouter.PathPrefix("").Subrouter()

	optionalAuth.Use(s.authenticate(OptionalAuth))
	{
		optionalAuth.Handle("/profiles", s.getProfile()).Methods("GET")
	}

	notAuth := apiRouter.PathPrefix("").Subrouter()
	{
		notAuth.Handle("/health", healtCheck())
	}
}
