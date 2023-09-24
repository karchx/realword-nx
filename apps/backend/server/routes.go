package server

import (
	"os"

	"github.com/rs/cors"
)

const (
	MustAuth     = true
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
		optionalAuth.Handle("/profiles/{username}", s.getProfile()).Methods("GET")
	}

	notAuth := apiRouter.PathPrefix("").Subrouter()
	{
		notAuth.Handle("/health", healtCheck())
		notAuth.Handle("/users", s.createUser()).Methods("POST")
		notAuth.Handle("/users/login", s.loginUser()).Methods("POST")
	}

	authApiRoutes := apiRouter.PathPrefix("").Subrouter()
	authApiRoutes.Use(s.authenticate(MustAuth))
	{
		authApiRoutes.Handle("/profiles/{username}/follow", s.followAction("follow")).Methods("POST")
	}
}
