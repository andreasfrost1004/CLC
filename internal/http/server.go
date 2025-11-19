package http

import (
	"CLC/internal/config"
	"net/http"
)

type Server struct {
	addr string
}

func NewServer(cfg config.Config) *Server {
	return &Server{
		addr: cfg.HTTP.Address,
	}
}
func (s *Server) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	return http.ListenAndServe(s.addr, mux)
}
