package interactor

import (
	"github.com/omarscd/academy-go-q42021/model"
	"github.com/omarscd/academy-go-q42021/usecase/presenter"
	"github.com/omarscd/academy-go-q42021/usecase/repository"
)

type standUserInteractor struct {
	StandUserRepository repository.StandUserRepository
	StandUserPresenter  presenter.StandUserPresenter
}

type StandUserInteractor interface {
	GetAll(su []*model.StandUser) ([]*model.StandUser, error)
	GetById(uint64) (*model.StandUser, error)
}

func NewStandUserInteractor(r repository.StandUserRepository, p presenter.StandUserPresenter) StandUserInteractor {
	return &standUserInteractor{r, p}
}

func (sui *standUserInteractor) GetAll(sus []*model.StandUser) ([]*model.StandUser, error) {
	su, err := sui.StandUserRepository.FindAll(sus)
	if err != nil {
		return nil, err
	}

	return sui.StandUserPresenter.ResponseStandUsers(su), nil
}

func (sui *standUserInteractor) GetById(id uint64) (*model.StandUser, error) {
	s, err := sui.StandUserRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return s, nil
}
