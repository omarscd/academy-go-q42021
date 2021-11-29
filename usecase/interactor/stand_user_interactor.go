package interactor

import (
	"github.com/omarscd/academy-go-q42021/model"
	"github.com/omarscd/academy-go-q42021/usecase/presenter"
	"github.com/omarscd/academy-go-q42021/usecase/repository"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
	PokemonPresenter  presenter.PokemonPresenter
}

type PokemonInteractor interface {
	GetAll(su []*model.Pokemon) ([]*model.Pokemon, error)
	GetById(uint64) (*model.Pokemon, error)
}

func NewPokemonInteractor(r repository.PokemonRepository, p presenter.PokemonPresenter) PokemonInteractor {
	return &pokemonInteractor{r, p}
}

func (sui *pokemonInteractor) GetAll(sus []*model.Pokemon) ([]*model.Pokemon, error) {
	su, err := sui.PokemonRepository.FindAll(sus)
	if err != nil {
		return nil, err
	}

	return sui.PokemonPresenter.ResponsePokemons(su), nil
}

func (sui *pokemonInteractor) GetById(id uint64) (*model.Pokemon, error) {
	s, err := sui.PokemonRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return s, nil
}
