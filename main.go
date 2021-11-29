package main

import (
	"fmt"

	"github.com/omarscd/academy-go-q42021/datastore"
	"github.com/omarscd/academy-go-q42021/registry"
	"github.com/omarscd/academy-go-q42021/router"
)

func main() {
	pkMap := datastore.NewPkMap()

	r := registry.NewRegistry(&pkMap)

	router := router.NewRouter(r.NewAppController())

	fmt.Println("Listening on port 8888")
	router.Run("localhost:8888")
}
