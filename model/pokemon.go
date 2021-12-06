package model

import "errors"

// Pokemon model
type Pokemon struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	MainType string `json:"type"`
}

// creates a new instance of Pokemon
func NewPokemon(id uint64, name string, mType string) (*Pokemon, error) {
	if id == 0 {
		return nil, errors.New("ID can't be 0")
	}
	if name == "" {
		return nil, errors.New("Must provide a name")
	}
	if mType == "" {
		return nil, errors.New("Must provide a type")
	}

	return &Pokemon{
		id, name, mType,
	}, nil
}
