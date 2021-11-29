package router

import (
	"github.com/gin-gonic/gin"
	"github.com/omarscd/academy-go-q42021/interface/controller"
)

func NewRouter(c controller.AppController) *gin.Engine {
	router := gin.Default()
	router.GET("/pokemons", func(ctx *gin.Context) { _ = c.GetPokemons(ctx) })
	router.GET("/pokemons/:id", func(ctx *gin.Context) { _ = c.GetPokemonById(ctx) })
	return router
}
