package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joaovds/diocese-santos/internal/liturgy"
)

func main() {
	mainMux := http.NewServeMux()
	muxV1 := http.NewServeMux()
	mainMux.Handle("/api/v1/", http.StripPrefix("/api/v1", muxV1))

	liturgy.NewHandlers(muxV1).SetupRoutes()

	log.Println("Server running on port", 3333)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", "3333"), mainMux); err != nil {
		log.Fatalf("could not start server: %s", err)
	}
}
