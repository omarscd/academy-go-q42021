package repository

import "github.com/omarscd/academy-go-q42021/model"

type StandUserRepository interface {
	FindAll([]*model.StandUser) ([]*model.StandUser, error)
	FindByID(id uint64) (*model.StandUser, error)
}
