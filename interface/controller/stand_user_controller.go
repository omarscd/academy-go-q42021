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

func NewPokemonController(sui interactor.PokemonInteractor) PokemonController {
	return &pokemonController{sui}
}

func (suc *pokemonController) GetPokemons(c *gin.Context) error {
	var u []*model.Pokemon

	u, err := suc.pokemonInteractor.GetAll(u)
	if err != nil {
		return err
	}

	c.IndentedJSON(http.StatusOK, u)
	return nil
}

func (suc *pokemonController) GetPokemonById(c *gin.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	u, err := suc.pokemonInteractor.GetById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return err
	}
	c.IndentedJSON(http.StatusOK, u)
	return nil
}
