package presenter

import model "github.com/omarscd/academy-go-q42021/model"

type PokemonPresenter interface {
	ResponsePokemons(pks []*model.Pokemon) []*model.Pokemon
}
