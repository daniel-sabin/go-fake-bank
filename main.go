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

func uuidGenerator() string {
	id := uuid.New()
	return id.String()
}

func main() {
	var addr = flag.String("addr", ":8000", "The addr of the application.")
	flag.Parse()

	keys := []string{
		uuidGenerator(),
		uuidGenerator(),
	}

	// Display for user at started
	for _, key := range keys {
		fmt.Printf("api-key-available %v\r\n", key)
	}

	i := repository.NewInMemoryApiKeyStore()
	as := repository.NewInMemoryAccountsStore()
	ts := repository.NewInMemoryTransactionsStore()
	cs := repository.NewInMemoryClientsStore(uuidGenerator)

	log.Fatal(http.ListenAndServe(*addr, server.NewServer(i, as, ts, cs, keys)))
}
