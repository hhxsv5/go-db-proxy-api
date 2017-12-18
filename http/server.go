package http

import (
	"net/http"
	"fmt"
)

type Server struct {
	listen   string
	handlers map[string]func(w http.ResponseWriter, r *http.Request)
}

func (s *Server) Run() {
	for pattern, handler := range s.handlers {
		http.HandleFunc(pattern, handler)
	}
	err := http.ListenAndServe(s.listen, nil)
	if err != nil {
		panic(err)
	}
}

func NewServer(handlers map[string]func(w http.ResponseWriter, r *http.Request)) *Server {
	cfg := GetConfig()
	fmt.Println(cfg)
	return &Server{cfg.Listen, handlers}
}
