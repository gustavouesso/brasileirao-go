package controller

import (
	"net/http"
	"strconv"

	"github.com/gustavouesso/brasileirao-go.git/model"
	"github.com/gustavouesso/brasileirao-go.git/service"
	"github.com/gin-gonic/gin"
)

type TeamController struct {
	service service.TeamService
}

func NewTeamController(service service.TeamService) TeamController {
	return TeamController{service: service}
}

func (controller *TeamController) initRoutes() {
	router := gin.Default()
	router.GET("/teams", controller.index)
	router.GET("/teams/:id", controller.show)
	router.POST("/teams", controller.create)
	router.PUT("/teams/:id", controller.update)
	router.DELETE("/teams/:id", controller.delete)
	router.Run(":8080")
}

func (controller *TeamController) index(c *gin.Context) {
	teams := controller.service.FindAll()
	c.JSON(http.StatusOK, teams)
}

func (controller *TeamController) show(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	team := controller.service.FindById(uint(id))
	c.JSON(http.StatusOK, team)
}

func (controller *TeamController) create(c *gin.Context) {
	var team model.Team
	c.BindJSON(&team)
	controller.service.Save(team)
	c.JSON(http.StatusCreated, team)
}

func (controller *TeamController) update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var team model.Team
	c.BindJSON(&team)
	controller.service.Update(uint(id), team)
	c.JSON(http.StatusOK, team)
}

func (controller *TeamController) delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	controller.service.Delete(uint(id))
	c.JSON(http.StatusNoContent, nil)
}
