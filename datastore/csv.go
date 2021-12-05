package datastore

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/omarscd/academy-go-q42021/model"
)

type PokemonDB struct {
	pkMap map[uint64]model.Pokemon
}

// Find returns a slice of all the Pokemons that pass the test function
func (pkDB *PokemonDB) Find(test func(model.Pokemon) bool) ([]*model.Pokemon, error) {
	pks := make([]*model.Pokemon, 0)
	for _, pk := range pkDB.pkMap {
		if tmp := pk; test(tmp) {
			pks = append(pks, &tmp)
		}
	}

	return pks, nil
}

// FindOne returns the first element that passes the test function
func (pkDB *PokemonDB) FindOne(test func(model.Pokemon) bool) (*model.Pokemon, error) {
	for _, pk := range pkDB.pkMap {
		if tmp := pk; test(tmp) {
			return &tmp, nil
		}
	}
	return nil, errors.New("Pokemon not found")
}

// InsertOnt appends the element to the csv and adds it to the pkMap
func (pkDB *PokemonDB) InsertOne(model.Pokemon) error {
	// TODO:
	return nil
}

// NewPokemonDB creates a new PokemonDB instance
func NewPokemonDB(path string) (*PokemonDB, error) {
	csvPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	csvfile, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = 3
	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	pkMap := make(map[uint64]model.Pokemon)
	for _, record := range rawCSVdata {
		id, err := strconv.ParseUint(record[0], 10, 32)
		if err != nil {
			log.Println("Invalid ID for record: ", record)
			continue
		}

		pk, err := model.NewPokemon(id, record[1], record[2])
		if err != nil {
			log.Println("Invalid values for record: ", record)
			continue
		}
		pkMap[id] = *pk
	}

	return &PokemonDB{pkMap}, nil
}
