package main

import (
	"github.com/gustavouesso/brasileirao-go.git/controller",
	"github.com/gustavouesso/brasileirao-go.git/infra",
	"github.com/gustavouesso/brasileirao-go.git/service",
)

func main() {
	db := infra.CreateConnection()
	championshipService := service.NewChampionshipService(db)
	championshipController := controller.NewChampionshipController(championshipService)

	championshipController.initRoutes()
}