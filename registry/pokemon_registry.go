package registry

import (
	"github.com/omarscd/academy-go-q42021/interface/controller"
	ip "github.com/omarscd/academy-go-q42021/interface/presenters"
	ir "github.com/omarscd/academy-go-q42021/interface/repository"
	"github.com/omarscd/academy-go-q42021/usecase/interactor"
	up "github.com/omarscd/academy-go-q42021/usecase/presenter"
	ur "github.com/omarscd/academy-go-q42021/usecase/repository"
)

// generates a controller with interactor
func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor())
}

// returns an interactor with repositories and presenter
func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonRepository(), r.NewPokemonExtApi(), r.NewPokemonPresenter())
}

// returns a pokemon repository which fullfills usecase interface
func (r *registry) NewPokemonRepository() ur.PokemonRepository {
	return ir.NewPokemonRepository(r.pkDB)
}

// returns a pokemon presenter which fullfills usecase interface

func (r *registry) NewPokemonPresenter() up.PokemonPresenter {
	return ip.NewPokemonPresenter()
}

func (r *registry) NewPokemonExtApi() ur.PokemonExtApi {
	return ir.NewPokemonExtApi("https://pokeapi.co/api/v2/")
}
