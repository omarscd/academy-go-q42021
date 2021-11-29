package datastore

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/omarscd/academy-go-q42021/model"
)

func NewSUSMap() map[uint64]model.Pokemon {
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
		fmt.Println(err)
		os.Exit(1)
	}

	susMap := make(map[uint64]model.Pokemon)
	for _, record := range rawCSVdata {
		id, err := strconv.ParseUint(record[0], 10, 32)
		if err != nil {
			continue
		}
		susMap[id] = model.Pokemon{
			ID:       id,
			Name:     record[1],
			MainType: record[2],
		}
	}

	return susMap
}
