package main

import (
	"engineecore/demobank-server/infra/repository"
	"engineecore/demobank-server/server"
	"flag"
	"fmt"
	"os"

	"github.com/google/uuid"
	"goyave.dev/goyave/v4"
)

func main() {
	//var addr = flag.String("addr", ":8000", "The addr of the application.")
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
	accountsStore := repository.NewInMemoryAccountsStore()
	router := server.RegisterRoutes(i, accountsStore, keys)

	if err := goyave.Start(router); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}
}
