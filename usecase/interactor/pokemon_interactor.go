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
	GetAll() ([]*model.Pokemon, error)
	GetById(uint64) (*model.Pokemon, error)
}

func NewPokemonInteractor(r repository.PokemonRepository, p presenter.PokemonPresenter) PokemonInteractor {
	return &pokemonInteractor{r, p}
}

func (pki *pokemonInteractor) GetAll() ([]*model.Pokemon, error) {
	pk, err := pki.PokemonRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return pki.PokemonPresenter.ResponsePokemons(pk), nil
}

func (pki *pokemonInteractor) GetById(id uint64) (*model.Pokemon, error) {
	pk, err := pki.PokemonRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return pk, nil
}
