package main

import (
	"engineecore/demobank-server/server"
	"engineecore/demobank-server/server/security"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func main() {

	keys := []string{
		uuid.New().String(),
		uuid.New().String(),
	}

	// Display for user at started
	for _, key := range keys {
		fmt.Printf("api-key-available %v\r\n", key)
	}

	log.Fatal(http.ListenAndServe(":8000", server.NewServer(security.NewInMemoryApiKeyStore(), keys)))
}
