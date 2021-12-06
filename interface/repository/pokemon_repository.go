package repository

import (
	"github.com/omarscd/academy-go-q42021/datastore"
	"github.com/omarscd/academy-go-q42021/model"
)

type pokemonRepository struct {
	pkDB datastore.PokemonDB
}

type PokemonRepository interface {
	GetAll() ([]*model.Pokemon, error)
	GetById(id uint64) (*model.Pokemon, error)
	InsertOne(model.Pokemon) error
	GetOdds(items, itemsPerWorker int64) ([]*model.Pokemon, error)
	GetEvens(items, itemsPerWorker int64) ([]*model.Pokemon, error)
}

func NewPokemonRepository(pkDB datastore.PokemonDB) PokemonRepository {
	return &pokemonRepository{pkDB}
}

func (pkr *pokemonRepository) GetAll() ([]*model.Pokemon, error) {
	pks, err := pkr.pkDB.Find(func(model.Pokemon) bool {
		return true
	})
	if err != nil {
		return nil, err
	}
	return pks, nil
}

func (pkr *pokemonRepository) GetById(id uint64) (*model.Pokemon, error) {
	pk, err := pkr.pkDB.FindOne(func(pk model.Pokemon) bool {
		return pk.ID == id
	})
	if err != nil {
		return nil, err
	}
	return pk, nil
}

func (pkr *pokemonRepository) InsertOne(pk model.Pokemon) error {
	err := pkr.pkDB.InsertOne(pk)
	return err
}

func (pkr *pokemonRepository) GetOdds(items, itemsPerWorker int64) ([]*model.Pokemon, error) {
	test := func(pk model.Pokemon) bool {
		return pk.ID%2 == 1
	}
	pks, err := pkr.pkDB.FindWP(test, items, itemsPerWorker)
	if err != nil {
		return nil, err
	}
	return pks, nil
}

func (pkr *pokemonRepository) GetEvens(items, itemsPerWorker int64) ([]*model.Pokemon, error) {
	test := func(pk model.Pokemon) bool {
		return pk.ID%2 == 0
	}
	pks, err := pkr.pkDB.FindWP(test, items, itemsPerWorker)
	if err != nil {
		return nil, err
	}
	return pks, nil
}
