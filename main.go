package main

import (
	"log"

	"github.com/omarscd/academy-go-q42021/datastore"
	"github.com/omarscd/academy-go-q42021/registry"
	"github.com/omarscd/academy-go-q42021/router"
)

func main() {
	pkDB, err := datastore.NewPokemonDB("./db/pokes.csv")
	if err != nil {
		log.Fatalf("Error initializing DB: %v\n", err)
	}

	r := registry.NewRegistry(pkDB)

	router := router.NewRouter(r.NewAppController())
	router.Run("localhost:8888")
}
