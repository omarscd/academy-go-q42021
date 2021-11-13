package registry

import (
	"github.com/omarscd/academy-go-q42021/interface/controller"
	"github.com/omarscd/academy-go-q42021/model"
)

type registry struct {
	susMap map[uint64]model.StandUser
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(susMap *map[uint64]model.StandUser) Registry {
	return &registry{*susMap}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewStandUserController()
}
