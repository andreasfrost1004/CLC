package http

import (
	"CLC/internal/config"
	"CLC/internal/database"
	"net/http"
)

type Server struct {
	addr string
	db   *database.Database
}

func NewServer(cfg config.Config, db *database.Database) *Server {
	return &Server{
		addr: cfg.HTTP.Address,
		db:   db,
	}
}
func (s *Server) Start() error {
	r := s.routes()
	return http.ListenAndServe(s.addr, r)
}
