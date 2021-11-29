package presenter

import (
	"github.com/omarscd/academy-go-q42021/model"
)

type pokemonPresenter struct{}

type PokemonPresenter interface {
	ResponsePokemons(su []*model.Pokemon) []*model.Pokemon
}

func NewPokemonPresenter() PokemonPresenter {
	return &pokemonPresenter{}
}

func (sup *pokemonPresenter) ResponsePokemons(sus []*model.Pokemon) []*model.Pokemon {
	return sus
}
