package repository

import "github.com/omarscd/academy-go-q42021/model"

type PokemonRepository interface {
	FindAll([]*model.Pokemon) ([]*model.Pokemon, error)
	FindByID(id uint64) (*model.Pokemon, error)
}
