package server

import "net/http"

type Server struct {
	http.Handler
}

func NewServer() *Server {
	server := new(Server)

	router := http.NewServeMux()
	router.Handle("/health", http.HandlerFunc(handleHealthCheck))
	server.Handler = router

	// server := &Server{make(chan bool)}
	return server
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
	w.WriteHeader(http.StatusOK)
}
