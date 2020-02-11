package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func (s *server) router() http.Handler {
	r := chi.NewRouter()

	r.Get("/*", s.handleRequest)

	return r
}

func (s *server) handleRequest(w http.ResponseWriter, r *http.Request) {
	resp, err := s.proxyService.ProcessRequest(r)
	if err != nil {
		fmt.Fprint(w, "unknown server error")
		return
	}
	for k, values := range resp.Header {
		for _, v := range values {
			w.Header().Add(k, v)
		}
	}
	fmt.Fprint(w, resp.Body)
}
