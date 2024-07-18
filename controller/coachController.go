package controller

import (
	"net/http"
	"strconv"

	"github.com/gustavouesso/brasileirao-go.git/model"
	"github.com/gustavouesso/brasileirao-go.git/service"
	"github.com/gin-gonic/gin"
)

type CoachController struct {
	service service.CoachService
}

func NewCoachController(service service.CoachService) CoachController {
	return CoachController{service: service}
}

func (controller *CoachController) initRoutes() {
	router := gin.Default()
	router.GET("/coaches", controller.index)
	router.GET("/coaches/:id", controller.show)
	router.POST("/coaches", controller.create)
	router.PUT("/coaches/:id", controller.update)
	router.DELETE("/coaches/:id", controller.delete)
	router.POST("/coaches/:id/teams/:teamId", controller.addCoachToTeam)
	router.DELETE("/coaches/:id/teams/:teamId", controller.removeCoachFromTeam)
	router.Run(":8080")
}

func (controller *CoachController) index(c *gin.Context) {
	coaches := controller.service.FindAll()
	c.JSON(http.StatusOK, coaches)
}

func (controller *CoachController) show(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	coach := controller.service.FindById(uint(id))
	c.JSON(http.StatusOK, coach)
}

func (controller *CoachController) create(c *gin.Context) {
	var coach model.Coach
	c.BindJSON(&coach)
	controller.service.Save(coach)
	c.JSON(http.StatusCreated, coach)
}

func (controller *CoachController) update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var coach model.Coach
	c.BindJSON(&coach)
	controller.service.Update(uint(id), coach)
	c.JSON(http.StatusOK, coach)
}

func (controller *CoachController) delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	controller.service.Delete(uint(id))
	c.JSON(http.StatusNoContent, nil)
}

func (controller *CoachController) addCoachToTeam(c *gin.Context) {
	coachId, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	teamId, _ := strconv.ParseUint(c.Param("teamId"), 10, 64)
	controller.service.AddCoachToTeam(uint(coachId), uint(teamId))
	c.JSON(http.StatusOK, nil)
}

func (controller *CoachController) removeCoachFromTeam(c *gin.Context) {
	coachId, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	teamId, _ := strconv.ParseUint(c.Param("teamId"), 10, 64)
	controller.service.RemoveCoachFromTeam(uint(coachId), uint(teamId))
	c.JSON(http.StatusOK, nil)
}
