package registry

import (
	"github.com/omarscd/academy-go-q42021/interface/controller"
	ip "github.com/omarscd/academy-go-q42021/interface/presenters"
	ir "github.com/omarscd/academy-go-q42021/interface/repository"
	"github.com/omarscd/academy-go-q42021/usecase/interactor"
	up "github.com/omarscd/academy-go-q42021/usecase/presenter"
	ur "github.com/omarscd/academy-go-q42021/usecase/repository"
)

func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor())
}

func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonRepository(), r.NewPokemonPresenter())
}

func (r *registry) NewPokemonRepository() ur.PokemonRepository {
	return ir.NewPokemonRepository(r.pkDB)
}

func (r *registry) NewPokemonPresenter() up.PokemonPresenter {
	return ip.NewPokemonPresenter()
}
