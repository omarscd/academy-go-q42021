package datastore

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/omarscd/academy-go-q42021/model"
)

type PokemonDB struct {
	pkMap map[uint64]model.Pokemon
	path  string
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

// InsertOne appends the element to the csv and adds it to the pkMap
func (pkDB *PokemonDB) InsertOne(pk model.Pokemon) error {
	// if record already exists, do not write again
	if _, ok := pkDB.pkMap[pk.ID]; ok {
		return nil
	}

	csvPath, err := filepath.Abs(pkDB.path)
	if err != nil {
		log.Printf("Could not access path: %v", pkDB.path)
		return err
	}

	file, err := os.OpenFile(csvPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Printf("Could not open file: %v", pkDB.path)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		strconv.FormatUint(pk.ID, 10),
		pk.Name,
		pk.MainType,
	}

	writer.Write(record)
	if err := writer.Error(); err != nil {
		log.Printf("Could not write record: %v", pk)
		return err
	}

	pkDB.pkMap[pk.ID] = pk
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

	return &PokemonDB{pkMap, path}, nil
}

// FindWP returns a slice of all the Pokemons that pass the test function
// similar to Find, the difference is FindWP uses a WorkerPool under the hood
func (pkDB *PokemonDB) FindWP(test func(model.Pokemon) bool, items, itemsPerWorker int64) ([]*model.Pokemon, error) {
	pks := []*model.Pokemon{}

	csvPath, err := filepath.Abs(pkDB.path)
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

	src := make(chan []string)
	out := make(chan *model.Pokemon, items)

	var wg sync.WaitGroup
	nWorkers := int(items/itemsPerWorker) + 1

	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go func(out chan *model.Pokemon, src chan []string) {
			defer wg.Done()
			var addedByWorker int64 = 0
			for record := range src {
				if cap(out) == len(out) {
					return
				}

				id, err := strconv.ParseUint(record[0], 10, 32)
				if err != nil {
					continue
				}

				pk, err := model.NewPokemon(id, record[1], record[2])
				if err != nil {
					continue
				}

				if test(*pk) {
					select {
					case out <- pk:
						if addedByWorker++; addedByWorker >= itemsPerWorker {
							return
						}
					default:
						return
					}
				}
			}
		}(out, src)
	}

	go func() {
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				continue
			}
			src <- record
		}
		close(src)
	}()

	wg.Wait()
	close(out)

	for pk := range out {
		pks = append(pks, pk)
	}

	return pks, nil
}
