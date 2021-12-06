package repository

import "github.com/omarscd/academy-go-q42021/model"

type PokemonExtApi interface {
	GetByName(string) (*model.Pokemon, error)
}
