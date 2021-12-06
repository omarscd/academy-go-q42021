package controller

import (
	"log"
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
	GetPokemons(c *gin.Context)
	GetPokemonById(c *gin.Context)
	GetPokemonExt(c *gin.Context)
	GetPokemonsByType(c *gin.Context)
}

func NewPokemonController(pki interactor.PokemonInteractor) PokemonController {
	return &pokemonController{pki}
}

func (pkc *pokemonController) GetPokemons(c *gin.Context) {
	pk, err := pkc.pokemonInteractor.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error."})
		return
	}

	c.IndentedJSON(http.StatusOK, pk)
}

func (pkc *pokemonController) GetPokemonById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Pokemon ID. Value must be numeric."})
		return
	}

	pk, err := pkc.pokemonInteractor.GetById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, pk)
}

func (pkc *pokemonController) GetPokemonExt(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Must provide a pokemon name."})
		return
	}

	pk, err := pkc.pokemonInteractor.GetExtPokeByName(name)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, pk)
}

func (pkc *pokemonController) GetPokemonsByType(c *gin.Context) {
	t := c.Query("type")

	items, err := strconv.ParseInt(c.DefaultQuery("items", "15"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Parameter 'items' must be an integer"})
		return
	}

	itemsPerWorker, err := strconv.ParseInt(c.DefaultQuery("items_per_worker", "10"), 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Parameter 'items_per_worker' must be an integer"})
		return
	}
	log.Println(items, itemsPerWorker)

	pks := []*model.Pokemon{}

	switch t {
	case "odd":
		pks, err = pkc.pokemonInteractor.GetOdds()
		c.IndentedJSON(http.StatusOK, gin.H{"data": pks})
		return
	case "even":
		pks, err = pkc.pokemonInteractor.GetEvens()
		c.IndentedJSON(http.StatusOK, gin.H{"data": pks})
		return
	default:
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid type filter"})
		return
	}
}
