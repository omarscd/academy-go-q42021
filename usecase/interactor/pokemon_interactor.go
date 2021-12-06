package interactor

import (
	"github.com/omarscd/academy-go-q42021/model"
	"github.com/omarscd/academy-go-q42021/usecase/presenter"
	"github.com/omarscd/academy-go-q42021/usecase/repository"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
	PokemonExtApi     repository.PokemonExtApi
	PokemonPresenter  presenter.PokemonPresenter
}

type PokemonInteractor interface {
	GetAll() ([]*model.Pokemon, error)
	GetById(uint64) (*model.Pokemon, error)
	GetExtPokeByName(string) (*model.Pokemon, error)
	GetOdds(items, itemsPerWorker int64) ([]*model.Pokemon, error)
	GetEvens(items, itemsPerWorker int64) ([]*model.Pokemon, error)
}

func NewPokemonInteractor(r repository.PokemonRepository, e repository.PokemonExtApi, p presenter.PokemonPresenter) PokemonInteractor {
	return &pokemonInteractor{r, e, p}
}

// gets all pokemons from the repository
func (pki *pokemonInteractor) GetAll() ([]*model.Pokemon, error) {
	pk, err := pki.PokemonRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return pki.PokemonPresenter.ResponsePokemons(pk), nil
}

// gets a pokemon by id from the repository
func (pki *pokemonInteractor) GetById(id uint64) (*model.Pokemon, error) {
	pk, err := pki.PokemonRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return pk, nil
}

// gets a pokemon by name from an external api and saves it to the repository
func (pki *pokemonInteractor) GetExtPokeByName(name string) (*model.Pokemon, error) {
	pk, err := pki.PokemonExtApi.GetByName(name)
	if err != nil {
		return nil, err
	}

	err = pki.PokemonRepository.InsertOne(*pk)
	if err != nil {
		return nil, err
	}

	return pk, nil
}

// gets as many pokemons with odd ids as specified by 'items'
func (pki *pokemonInteractor) GetOdds(items, itemsPerWorker int64) ([]*model.Pokemon, error) {
	return pki.PokemonRepository.GetOdds(items, itemsPerWorker)
}

// gets as many pokemons with even ids as specified by 'items'
func (pki *pokemonInteractor) GetEvens(items, itemsPerWorker int64) ([]*model.Pokemon, error) {
	return pki.PokemonRepository.GetEvens(items, itemsPerWorker)
}
