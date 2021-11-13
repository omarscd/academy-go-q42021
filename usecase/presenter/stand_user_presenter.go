package presenter

import model "github.com/omarscd/academy-go-q42021/model"

type StandUserPresenter interface {
	ResponseStandUsers(sus []*model.StandUser) []*model.StandUser
}
