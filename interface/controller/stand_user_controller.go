package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/omarscd/academy-go-q42021/model"
	"github.com/omarscd/academy-go-q42021/usecase/interactor"
)

type standUserController struct {
	standUserInteractor interactor.StandUserInteractor
}

type StandUserController interface {
	GetStandUsers(c *gin.Context) error
	GetStandUserById(c *gin.Context) error
}

func NewStandUserController(sui interactor.StandUserInteractor) StandUserController {
	return &standUserController{sui}
}

func (suc *standUserController) GetStandUsers(c *gin.Context) error {
	var u []*model.StandUser

	u, err := suc.standUserInteractor.GetAll(u)
	if err != nil {
		return err
	}

	c.IndentedJSON(http.StatusOK, u)
	return nil
}

func (suc *standUserController) GetStandUserById(c *gin.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	u, err := suc.standUserInteractor.GetById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return err
	}
	c.IndentedJSON(http.StatusOK, u)
	return nil
}
