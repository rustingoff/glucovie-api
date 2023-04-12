package glucovieapi

import (
	"log"
	"net/http"
)

type server struct {
	addr    string
	handler http.Handler
}

func NewHTTPServer(addr string, handler http.Handler) *server {
	return &server{addr: addr, handler: handler}
}

func (s *server) Start() {
	srv := http.Server{
		Addr:    s.addr,
		Handler: s.handler,
	}

	log.Fatal(srv.ListenAndServe())
}
