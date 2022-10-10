package main

import (
	"engineecore/demobank-server/server"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8000", server.NewServer()))
}
