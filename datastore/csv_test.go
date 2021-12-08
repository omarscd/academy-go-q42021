package datastore

import (
	"testing"

	"github.com/omarscd/academy-go-q42021/model"
)

var mockPokemonDB = PokemonDB{
	pkMap: map[uint64]model.Pokemon{
		1: {
			ID:       1,
			Name:     "bulbasaur",
			MainType: "grass",
		},
		4: {
			ID:       4,
			Name:     "charmander",
			MainType: "fire",
		},
		9: {
			ID:       9,
			Name:     "blastoise",
			MainType: "water",
		},
	},
	path: "./test.csv",
}

func TestDatastore_Find(t *testing.T) {
	tests := []struct {
		testFunc       func(pk model.Pokemon) bool
		expectedLength int
	}{
		{
			func(pk model.Pokemon) bool { return true },
			3,
		},
		{
			func(pk model.Pokemon) bool { return pk.ID == 1 },
			1,
		},
		{
			func(pk model.Pokemon) bool { return pk.Name == "Romario" },
			0,
		},
	}

	for _, tt := range tests {
		gotPks, gotErr := mockPokemonDB.Find(tt.testFunc)
		if gotErr != nil {
			t.Error("Got unexpected error: ", gotErr)
		}
		if len(gotPks) != tt.expectedLength {
			t.Errorf("Should get len 3, instead got %v\n", len(gotPks))
		}
	}
}

func TestDatastore_FindOne(t *testing.T) {
	tests := []struct {
		testFunc     func(pk model.Pokemon) bool
		expectedPk   *model.Pokemon
		expectsError bool
	}{
		{
			func(pk model.Pokemon) bool { return pk.ID == 1 },
			&model.Pokemon{
				ID:       1,
				Name:     "bulbasaur",
				MainType: "grass",
			},
			false,
		},
		{
			func(pk model.Pokemon) bool { return pk.Name == "Romario" },
			nil,
			true,
		},
	}

	for _, tt := range tests {
		gotPk, gotErr := mockPokemonDB.FindOne(tt.testFunc)
		if gotErr != nil {
			if !tt.expectsError {
				t.Error("Got unexpected error: ", gotErr)
			}
		} else if gotPk.ID != tt.expectedPk.ID {
			t.Errorf("expected %v, instead got %v\n", tt.expectedPk, gotPk)
		}
	}
}

func TestDatastore_FindWP(t *testing.T) {
	tests := []struct {
		testFunc       func(pk model.Pokemon) bool
		expectedLength int
	}{
		{
			func(pk model.Pokemon) bool { return true },
			3,
		},
		{
			func(pk model.Pokemon) bool { return pk.ID == 1 },
			1,
		},
		{
			func(pk model.Pokemon) bool { return pk.ID%2 == 0 },
			2,
		},
		{
			func(pk model.Pokemon) bool { return pk.MainType == "water" },
			1,
		},
	}

	mockDB, _ := NewPokemonDB("../db/test_sample.csv")
	for _, tt := range tests {
		gotPks, gotErr := mockDB.FindWP(tt.testFunc, 5, 5)
		if gotErr != nil {
			t.Error("Got unexpected error: ", gotErr)
		}
		if len(gotPks) != tt.expectedLength {
			t.Errorf("Should get len 3, instead got %v\n", len(gotPks))
		}
	}
}
