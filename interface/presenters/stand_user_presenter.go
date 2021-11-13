package presenter

import (
	"github.com/omarscd/academy-go-q42021/model"
)

type standUserPresenter struct{}

type StandUserPresenter interface {
	ResponseStandUsers(su []*model.StandUser) []*model.StandUser
}

func NewStandUserPresenter() StandUserPresenter {
	return &standUserPresenter{}
}

func (sup *standUserPresenter) ResponseStandUsers(sus []*model.StandUser) []*model.StandUser {
	return sus
}
