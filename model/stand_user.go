package model

import "errors"

type StandUser struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Stand string `json:"stand"`
}

func NewStandUser(id uint64, name string, stand string) (*StandUser, error) {
	if id == 0 {
		return nil, errors.New("ID can't be 0")
	}
	if name == "" {
		return nil, errors.New("Must provide a name")
	}
	if stand == "" {
		return nil, errors.New("Must provide a stand")
	}

	return &StandUser{
		id, name, stand,
	}, nil
}
