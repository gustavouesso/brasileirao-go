package controller

import (
	"net/http"
	"strconv"

	"github.com/gustavouesso/brasileirao-go.git/model"
	"github.com/gustavouesso/brasileirao-go.git/service"
	"github.com/gin-gonic/gin"
)

type MatchController struct {
	service service.MatchService
}

func NewMatchController(service service.MatchService) MatchController {
	return MatchController{service: service}
}

func (controller *MatchController) initRoutes() {
	router := gin.Default()
	router.GET("/matches", controller.index)
	router.GET("/matches/:id", controller.show)
	router.POST("/matches", controller.create)
	router.PUT("/matches/:id", controller.update)
	router.DELETE("/matches/:id", controller.delete)
	router.Run(":8080")
}

func (controller *MatchController) index(c *gin.Context) {
	matches := controller.service.FindAll()
	c.JSON(http.StatusOK, matches)
}

func (controller *MatchController) show(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	match := controller.service.FindById(uint(id))
	c.JSON(http.StatusOK, match)
}

func (controller *MatchController) create(c *gin.Context) {
	var match model.Match
	c.BindJSON(&match)
	controller.service.Save(match)
	c.JSON(http.StatusCreated, match)
}

func (controller *MatchController) update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var match model.Match
	c.BindJSON(&match)
	controller.service.Update(uint(id), match)
	c.JSON(http.StatusOK, match)
}

func (controller *MatchController) delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	controller.service.Delete(uint(id))
	c.JSON(http.StatusNoContent, nil)
}
