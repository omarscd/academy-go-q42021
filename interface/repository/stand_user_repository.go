package repository

import (
	"errors"

	"github.com/omarscd/academy-go-q42021/model"
)

type standUserRepository struct {
	SUSMap map[uint64]model.StandUser
}

type StandUserRepository interface {
	FindAll(sus []*model.StandUser) ([]*model.StandUser, error)
	FindByID(id uint64) (*model.StandUser, error)
}

func NewStandUserRepository(susMap map[uint64]model.StandUser) StandUserRepository {
	return &standUserRepository{susMap}
}

func (sur *standUserRepository) FindAll(sus []*model.StandUser) ([]*model.StandUser, error) {
	for _, su := range sur.SUSMap {
		tmp := su
		sus = append(sus, &tmp)
	}

	return sus, nil
}

func (sur *standUserRepository) FindByID(id uint64) (*model.StandUser, error) {
	su, ok := sur.SUSMap[id]
	if ok == true {
		return &su, nil
	}
	return nil, errors.New("Stand user not found")
}
