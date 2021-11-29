package presenter

import (
	"github.com/omarscd/academy-go-q42021/model"
)

type pokemonPresenter struct{}

type PokemonPresenter interface {
	ResponsePokemons([]*model.Pokemon) []*model.Pokemon
}

func NewPokemonPresenter() PokemonPresenter {
	return &pokemonPresenter{}
}

func (pkp *pokemonPresenter) ResponsePokemons(pks []*model.Pokemon) []*model.Pokemon {
	return pks
}
