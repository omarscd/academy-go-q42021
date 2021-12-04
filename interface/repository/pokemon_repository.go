package repository

import (
	"errors"

	"github.com/omarscd/academy-go-q42021/model"
)

type pokemonRepository struct {
	pkMap map[uint64]model.Pokemon
}

type PokemonRepository interface {
	Find(func(model.Pokemon) bool) ([]*model.Pokemon, error)
	FindByID(id uint64) (*model.Pokemon, error)
}

func NewPokemonRepository(pkMap map[uint64]model.Pokemon) PokemonRepository {
	return &pokemonRepository{pkMap}
}

func (pkr *pokemonRepository) Find(test func(model.Pokemon) bool) ([]*model.Pokemon, error) {
	pks := make([]*model.Pokemon, 0)
	for _, pk := range pkr.pkMap {
		if tmp := pk; test(tmp) {
			pks = append(pks, &tmp)
		}
	}

	return pks, nil
}

func (pkr *pokemonRepository) FindByID(id uint64) (*model.Pokemon, error) {
	if pk, ok := pkr.pkMap[id]; ok {
		return &pk, nil
	}
	return nil, errors.New("Pokemon ID not found")
}
