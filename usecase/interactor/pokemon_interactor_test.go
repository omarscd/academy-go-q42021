package interactor

import (
	"errors"
	"testing"

	"github.com/omarscd/academy-go-q42021/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var pokemons = []*model.Pokemon{
	{
		ID:       1,
		Name:     "bulbasaur",
		MainType: "grass",
	},
	{
		ID:       4,
		Name:     "charmander",
		MainType: "fire",
	},
	{
		ID:       9,
		Name:     "blastoise",
		MainType: "water",
	},
}

type mockPkRepo struct {
	mock.Mock
}

type mockExtApi struct {
	mock.Mock
}

type mockPkPresenter struct {
	mock.Mock
}

// Repository interface mocks
func (mPkRepo mockPkRepo) GetAll() ([]*model.Pokemon, error) {
	args := mPkRepo.Called()
	return args.Get(0).([]*model.Pokemon), args.Error(1)
}

func (mPkRepo mockPkRepo) GetById(id uint64) (*model.Pokemon, error) {
	args := mPkRepo.Called(id)
	return args.Get(0).(*model.Pokemon), args.Error(1)
}

func (mPkRepo mockPkRepo) InsertOne(model.Pokemon) error {
	return nil
}

func (mPkRepo mockPkRepo) GetOdds(items, itemsPerWorker int64) ([]*model.Pokemon, error) {
	return []*model.Pokemon{}, nil
}

func (mPkRepo mockPkRepo) GetEvens(items, itemsPerWorker int64) ([]*model.Pokemon, error) {
	return []*model.Pokemon{}, nil
}

// External api interface mocks
func (pkApi mockExtApi) GetByName(name string) (*model.Pokemon, error) {
	args := pkApi.Called(name)
	return args.Get(0).(*model.Pokemon), args.Error(1)
}

// Presenter interface mocks
func (mPkPres mockPkPresenter) ResponsePokemons(pks []*model.Pokemon) []*model.Pokemon {
	return []*model.Pokemon{}
}

// ------- Tests --------
func TestInteractor_GetById(t *testing.T) {
	tt := []struct {
		name       string
		id         uint64
		mockPk     *model.Pokemon
		mockErr    error
		want       *model.Pokemon
		expectsErr bool
	}{{
		"Valid ID",
		9,
		pokemons[2],
		nil,
		pokemons[2],
		false,
	}, {
		"ID not found",
		54,
		nil,
		errors.New(""),
		nil,
		true,
	}}

	for _, test := range tt {
		mockRepo := mockPkRepo{}
		mockRepo.On("GetById", test.id).Return(test.mockPk, test.mockErr)

		pki := pokemonInteractor{
			PokemonRepository: mockRepo,
		}

		got, err := pki.GetById(test.id)
		assert.Equal(t, got, test.want)
		if test.expectsErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}
