package registry

import (
	"github.com/omarscd/academy-go-q42021/datastore"
	"github.com/omarscd/academy-go-q42021/interface/controller"
)

type registry struct {
	pkDB datastore.PokemonDB
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(pkDB *datastore.PokemonDB) Registry {
	return &registry{*pkDB}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewPokemonController()
}
