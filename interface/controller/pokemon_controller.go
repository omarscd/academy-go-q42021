package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/omarscd/academy-go-q42021/model"
	"github.com/omarscd/academy-go-q42021/usecase/interactor"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

type PokemonController interface {
	GetPokemons(c *gin.Context) error
	GetPokemonById(c *gin.Context) error
}

func NewPokemonController(pki interactor.PokemonInteractor) PokemonController {
	return &pokemonController{pki}
}

func (pkc *pokemonController) GetPokemons(c *gin.Context) error {
	var pk []*model.Pokemon

	pk, err := pkc.pokemonInteractor.GetAll(pk)
	if err != nil {
		return err
	}

	c.IndentedJSON(http.StatusOK, pk)
	return nil
}

func (pkc *pokemonController) GetPokemonById(c *gin.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Pokemon ID. Value must be numeric."})
		return err
	}

	pk, err := pkc.pokemonInteractor.GetById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return err
	}

	c.IndentedJSON(http.StatusOK, pk)
	return nil
}
