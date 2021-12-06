package repository

import "github.com/omarscd/academy-go-q42021/model"

type PokemonRepository interface {
	GetAll() ([]*model.Pokemon, error)
	GetById(id uint64) (*model.Pokemon, error)
	InsertOne(model.Pokemon) error
	GetOdds(items, itemsPerWorker int64) ([]*model.Pokemon, error)
	GetEvens(items, itemsPerWorker int64) ([]*model.Pokemon, error)
}
