package controller

import (
	"net/http"
	"strconv"

	"github.com/gustavouesso/brasileirao-go.git/model"
	"github.com/gustavouesso/brasileirao-go.git/service"
	"github.com/gin-gonic/gin"
)

type ChampionshipController struct {
	service service.ChampionshipService
}

func NewChampionshipController(service service.ChampionshipService) ChampionshipController {
	return ChampionshipController{service: service}
}

func (controller *ChampionshipController) initRoutes() {
	router := gin.Default()
	router.GET("/championships", controller.index)
	router.GET("/championships/:id", controller.show)
	router.POST("/championships", controller.create)
	router.PUT("/championships/:id", controller.update)
	router.DELETE("/championships/:id", controller.delete)
	router.Run(":8080")
}

func (controller *ChampionshipController) index(c *gin.Context) {
	championships := controller.service.FindAll()
	c.JSON(http.StatusOK, championships)
}

func (controller *ChampionshipController) show(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	championship := controller.service.FindById(uint(id))
	c.JSON(http.StatusOK, championship)
}

func (controller *ChampionshipController) create(c *gin.Context) {
	var championship model.Championship
	c.BindJSON(&championship)
	controller.service.Save(championship)
	c.JSON(http.StatusCreated, championship)
}

func (controller *ChampionshipController) update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var championship model.Championship
	c.BindJSON(&championship)
	controller.service.Update(uint(id), championship)
	c.JSON(http.StatusOK, championship)
}

func (controller *ChampionshipController) delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	controller.service.Delete(uint(id))
	c.JSON(http.StatusNoContent, nil)
}
