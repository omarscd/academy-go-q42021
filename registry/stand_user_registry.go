package registry

import (
	"github.com/omarscd/academy-go-q42021/interface/controller"
	ip "github.com/omarscd/academy-go-q42021/interface/presenters"
	ir "github.com/omarscd/academy-go-q42021/interface/repository"
	"github.com/omarscd/academy-go-q42021/usecase/interactor"
	up "github.com/omarscd/academy-go-q42021/usecase/presenter"
	ur "github.com/omarscd/academy-go-q42021/usecase/repository"
)

func (r *registry) NewStandUserController() controller.StandUserController {
	return controller.NewStandUserController(r.NewStandUserInteractor())
}

func (r *registry) NewStandUserInteractor() interactor.StandUserInteractor {
	return interactor.NewStandUserInteractor(r.NewStandUserRepository(), r.NewStandUserPresenter())
}

func (r *registry) NewStandUserRepository() ur.StandUserRepository {
	return ir.NewStandUserRepository(r.susMap)
}

func (r *registry) NewStandUserPresenter() up.StandUserPresenter {
	return ip.NewStandUserPresenter()
}
