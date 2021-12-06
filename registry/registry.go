package registry

import (
	"github.com/omarscd/academy-go-q42021/datastore"
	"github.com/omarscd/academy-go-q42021/interface/controller"
)

type registry struct {
	pkDB datastore.PokemonDB
}

// contract for the dependency resolver
type Registry interface {
	NewAppController() controller.AppController
}

// takes a PokemonDB to pass it down to the repository
func NewRegistry(pkDB *datastore.PokemonDB) Registry {
	return &registry{*pkDB}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewPokemonController()
}
