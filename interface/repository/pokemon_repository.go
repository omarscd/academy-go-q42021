package repository

import (
	"errors"

	"github.com/omarscd/academy-go-q42021/model"
)

type pokemonRepository struct {
	SUSMap map[uint64]model.Pokemon
}

type PokemonRepository interface {
	FindAll(sus []*model.Pokemon) ([]*model.Pokemon, error)
	FindByID(id uint64) (*model.Pokemon, error)
}

func NewPokemonRepository(susMap map[uint64]model.Pokemon) PokemonRepository {
	return &pokemonRepository{susMap}
}

func (sur *pokemonRepository) FindAll(sus []*model.Pokemon) ([]*model.Pokemon, error) {
	for _, su := range sur.SUSMap {
		tmp := su
		sus = append(sus, &tmp)
	}

	return sus, nil
}

func (sur *pokemonRepository) FindByID(id uint64) (*model.Pokemon, error) {
	su, ok := sur.SUSMap[id]
	if ok == true {
		return &su, nil
	}
	return nil, errors.New("Stand user not found")
}
