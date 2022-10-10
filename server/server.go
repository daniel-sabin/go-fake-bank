package server

import "net/http"

type Server struct {
	http.Handler
}

func NewServer() *Server {
	server := new(Server)

	router := http.NewServeMux()

	router.Handle("/health", http.HandlerFunc(handleHealthCheck))
	router.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swaggerui"))))

	server.Handler = router

	return server
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
	w.WriteHeader(http.StatusOK)
}
