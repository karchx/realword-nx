package server

import (
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	log "github.com/gothew/l-og"
)

type Server struct {
	server *http.Server
	router *mux.Router
}

func NewServer() *Server {
	s := Server{
		server: &http.Server{
			WriteTimeout: 5 * time.Second,
			ReadTimeout:  5 * time.Second,
			IdleTimeout:  5 * time.Second,
		},
		router: mux.NewRouter().StrictSlash(true),
	}

	return &s
}

func (s *Server) Run(port string) error {
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	s.server.Addr = port
	log.Info("Server starting on %s", port)
	return s.server.ListenAndServe()
}

func healtCheck() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := M{
			"status":  "available",
			"message": "healthy",
			"data":    M{"hello": "beautiful"},
		}

		writeJSON(w, http.StatusOK, resp)
	})
}
