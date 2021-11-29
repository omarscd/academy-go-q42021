package registry

import (
	"github.com/omarscd/academy-go-q42021/interface/controller"
	"github.com/omarscd/academy-go-q42021/model"
)

type registry struct {
	pkMap map[uint64]model.Pokemon
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(pkMap *map[uint64]model.Pokemon) Registry {
	return &registry{*pkMap}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewPokemonController()
}
