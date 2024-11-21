package main

import (
	"log"
	"net/http"

	"github.com/gowebexamples/goreddit/postgres"
	"github.com/gowebexamples/goreddit/web"
)

func main() {
	dsn := "postgres://postgres:example@192.168.106.2/postgres?sslmode=disable"

	store, err := postgres.NewStore(dsn)
	if err != nil {
		log.Fatal(err)
	}

	sessions, err := web.NewSessionManager(dsn)
	if err != nil {
		log.Fatal(err)
	}

	csrfKey := []byte("01234567890123456789012345678901")
	h := web.NewHandler(store, sessions, csrfKey)
	http.ListenAndServe(":8080", h)
}