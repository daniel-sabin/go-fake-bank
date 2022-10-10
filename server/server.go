package server

import (
	"log"
	"net/http"

	_ "engineecore/demobank-server/statik" // path to generated statik.go

	"github.com/rakyll/statik/fs"
)

type Server struct {
	http.Handler
}

func NewServer() *Server {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	server := new(Server)

	router := http.NewServeMux()

	router.Handle("/health", http.HandlerFunc(handleHealthCheck))
	router.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", http.FileServer(statikFS)))

	server.Handler = router

	return server
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
	w.WriteHeader(http.StatusOK)
}
