package router

import (
	"github.com/gin-gonic/gin"
	"github.com/omarscd/academy-go-q42021/interface/controller"
)

func NewRouter(c controller.AppController) *gin.Engine {
	router := gin.Default()
	router.GET("/stand_users", func(ctx *gin.Context) { _ = c.GetStandUsers(ctx) })
	router.GET("/stand_users/:id", func(ctx *gin.Context) { _ = c.GetStandUserById(ctx) })
	return router
}