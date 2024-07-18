service controller

import (
	"net/http"
	"strconv"

	"github.com/gustavouesso/brasileirao-go.git/model"
	"github.com/gustavouesso/brasileirao-go.git/service"
	"github.com/gin-gonic/gin"
)

type PlayerController struct {
	service service.PlayerService
}

func NewPlayerController(service service.PlayerService) PlayerController {
	return PlayerController{service: service}
}

func (controller *PlayerController) initRoutes() {
	router := gin.Default()
	router.GET("/players", controller.index)
	router.GET("/players/:id", controller.show)
	router.POST("/players", controller.create)
	router.PUT("/players/:id", controller.update)
	router.DELETE("/players/:id", controller.delete)
	router.POST("/players/:id/teams/:teamId", controller.addPlayerToTeam)
	router.DELETE("/players/:id/teams/:teamId", controller.removePlayerFromTeam)
	router.GET("players/:nation", controller.FindByNationality)
	router.GET("players/:position", controller.FindByPosition)
	router.GET("players/min/:min/max/:max", controller.FindByAgeRange)
	router.Run(":8080")
}

func (controller *PlayerController) index(c *gin.Context) {
	players := controller.service.FindAll()
	c.JSON(http.StatusOK, players)
}

func (controller *PlayerController) show(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	player := controller.service.FindById(uint(id))
	c.JSON(http.StatusOK, player)
}

func (controller *PlayerController) create(c *gin.Context) {
	var player model.Player
	c.BindJSON(&player)
	controller.service.Save(player)
	c.JSON(http.StatusCreated, player)
}

func (controller *PlayerController) update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var player model.Player
	c.BindJSON(&player)
	controller.service.Update(uint(id), player)
	c.JSON(http.StatusOK, player)
}

func (controller *PlayerController) delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	controller.service.Delete(uint(id))
	c.JSON(http.StatusNoContent, nil)
}

func (controller *PlayerController) addPlayerToTeam(c *gin.Context) {
	playerId, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	teamId, _ := strconv.ParseUint(c.Param("teamId"), 10, 64)
	controller.service.AddPlayerToTeam(uint(playerId), uint(teamId))
	c.JSON(http.StatusOK, nil)
}

func (controller *PlayerController) removePlayerFromTeam(c *gin.Context) {
	playerId, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	teamId, _ := strconv.ParseUint(c.Param("teamId"), 10, 64)
	controller.service.RemovePlayerFromTeam(uint(playerId), uint(teamId))
	c.JSON(http.StatusOK, nil)
}

func (controller *PlayerController) FindByNationality(c *gin.Context) {
	nation := c.Param("nation")
	players := controller.service.FindByNationality(nation)
	c.JSON(http.StatusOK, players)
}

func (controller *PlayerController) FindByPosition(c *gin.Context) {
	position := c.Param("position")
	players := controller.service.FindByPosition(position)
	c.JSON(http.StatusOK, players)
}

func (controller *PlayerController) FindByAgeRange(c *gin.Context) {
	min, _ := strconv.ParseUint(c.Param("min"), 10, 64)
	max, _ := strconv.ParseUint(c.Param("max"), 10, 64)
	players := controller.service.FindByAgeRange(uint(min), uint(max))
	c.JSON(http.StatusOK, players)
}
