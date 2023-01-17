package main

import (
	"engineecore/demobank-server/infra/repository"
	"engineecore/demobank-server/server"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	var addr = flag.String("addr", ":8000", "The addr of the application.")
	flag.Parse()

	keys := []string{
		uuid.New().String(),
		uuid.New().String(),
	}

	// Display for user at started
	for _, key := range keys {
		fmt.Printf("api-key-available %v\r\n", key)
	}

	i := repository.NewInMemoryApiKeyStore()
	as := repository.NewInMemoryAccountsStore()
	ts := repository.NewInMemoryTransactionsStore()

	log.Fatal(http.ListenAndServe(*addr, server.NewServer(i, as, ts, keys)))
}
