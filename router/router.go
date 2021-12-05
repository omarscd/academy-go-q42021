package router

import (
	"github.com/gin-gonic/gin"
	"github.com/omarscd/academy-go-q42021/interface/controller"
)

func NewRouter(c controller.AppController) *gin.Engine {
	router := gin.Default()
	router.GET("/pokemons", c.GetPokemons)
	router.GET("/pokemons/:id", c.GetPokemonById)
	router.GET("/pokemons/ext/:name", c.GetPokemonExt)
	return router
}
