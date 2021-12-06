package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/omarscd/academy-go-q42021/model"
)

type pokemonExtApi struct {
	baseUrl string
}

type extPokemon struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

// contract for an external API client
type PokemonExtApi interface {
	GetByName(string) (*model.Pokemon, error)
}

// creates a new instance of PokemonExtApi
func NewPokemonExtApi(name string) *pokemonExtApi {
	return &pokemonExtApi{name}
}

// gets a pokemon by its name
func (api *pokemonExtApi) GetByName(name string) (*model.Pokemon, error) {
	resp, err := http.Get(api.baseUrl + "pokemon/" + name)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result extPokemon
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
		return nil, err
	}

	pk, err := model.NewPokemon(uint64(result.ID), result.Name, result.Types[0].Type.Name)
	if err != nil {
		return nil, err
	}

	return pk, nil
}
