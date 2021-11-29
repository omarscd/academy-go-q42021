package datastore

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/omarscd/academy-go-q42021/model"
)

func NewPkMap() map[uint64]model.Pokemon {
	csvPath, _ := filepath.Abs("./db/pokes.csv")
	csvfile, err := os.Open(csvPath)
	if err != nil {
		log.Fatal(err)
		panic("Can't read file")
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = 3
	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	pkMap := make(map[uint64]model.Pokemon)
	for _, record := range rawCSVdata {
		id, err := strconv.ParseUint(record[0], 10, 32)
		if err != nil {
			continue
		}
		pkMap[id] = model.Pokemon{
			ID:       id,
			Name:     record[1],
			MainType: record[2],
		}
	}

	return pkMap
}
