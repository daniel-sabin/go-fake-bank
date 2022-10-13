package server

import (
	"log"
	"net/http"

	"engineecore/demobank-server/domain/security"
	_ "engineecore/demobank-server/statik" // path to generated statik.go

	"github.com/rakyll/statik/fs"
)

type Server struct {
	http.Handler
}

func storeApiKeys(i security.ApiKeyStore, apiKeys []string) {
	store := security.ApiKeyStoreFactory(i)
	for _, key := range apiKeys {
		store(key)
	}
}

func NewServer(i security.ApiKeyStore, k []string) *Server {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	server := new(Server)
	storeApiKeys(i, k)

	router := http.NewServeMux()

	router.Handle("/health", http.HandlerFunc(handleHealthCheck))
	router.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", http.FileServer(statikFS)))
	router.Handle("/applications", http.HandlerFunc(handleApplicationsFactory(i)))
	server.Handler = router

	return server
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func handleApplicationsFactory(i security.ApiKeyStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		isKeyAllowed := security.IsKeyAllowedFactory(i)
		allowed, _ := isKeyAllowed(r.Header.Get("x-api-key"))
		if !allowed {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.Write([]byte("{applications: ok}"))
		}
	}
}
