package model

import (
	"testing"
)

func TestModel_NewPokemon(t *testing.T) {
	tests := []struct {
		id    uint64
		name  string
		mtype string
	}{
		{
			0,
			"a",
			"a",
		},
		{
			1,
			"",
			"b",
		},
		{
			2,
			"c",
			"",
		},
	}

	for _, tt := range tests {
		_, err := NewPokemon(tt.id, tt.name, tt.mtype)
		if err == nil {
			t.Error("Did not get expected error from Pokemon initializer")
		}
	}
}
