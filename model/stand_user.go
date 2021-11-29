package model

import "errors"

type Pokemon struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	MainType string `json:"type"`
}

func NewPokemon(id uint64, name string, mt string) (*Pokemon, error) {
	if id == 0 {
		return nil, errors.New("ID can't be 0")
	}
	if name == "" {
		return nil, errors.New("Must provide a name")
	}
	if mt == "" {
		return nil, errors.New("Must provide a type")
	}

	return &Pokemon{
		id, name, mt,
	}, nil
}
