package service

import (
	"github.com/gustavouesso/brasileirao-go.git/model"
	"gorm.io/gorm"
)

// PlayerService provides methods to manipulate the player model

type PlayerService struct {
	DB *gorm.DB
}

// NewPlayerService creates a new player service

func NewPlayerService(db *gorm.DB) PlayerService {
	return PlayerService{DB: db}
}

// Save persists a player in the database

func (service *PlayerService) Save(player model.Player) model.Player {
	service.DB.Create(&player)
	return player
}

// FindAll returns all players from the database

func (service *PlayerService) FindAll() []model.Player {
	var players []model.Player
	service.DB.Find(&players)
	return players
}

// FindById returns a player by its ID

func (service *PlayerService) FindById(id uint) model.Player {
	var player model.Player
	service.DB.First(&player, id)
	return player
}

// Update updates a player in the database

func (service *PlayerService) Update(id uint, player model.Player) model.Player {
	
	service.DB.Save(&player)
	return player
}

// Delete removes a player from the database

func (service *PlayerService) Delete(id uint) {
	service.DB.Delete(&model.Player{}, id)
}

// FindByTeam returns all players from a team

func (service *PlayerService) FindByTeam(teamId uint) []model.Player {
	var players []model.Player
	service.DB.Joins("JOIN team_players ON players.id = team_players.player_id").Where("team_players.team_id = ?", teamId).Find(&players)
	return players
}

// AddToTeam adds a player to a team

func (service *PlayerService) AddToTeam(teamId uint, playerId uint) {
	player := service.FindById(playerId)
	team := service.DB.First(&model.Team{}, teamId)
	team.Model.(*model.Team).Players = append(team.Model.(*model.Team).Players, player)
	service.DB.Save(team.Model.(*model.Team))
}

// RemoveFromTeam removes a player from a team

func (service *PlayerService) RemoveFromTeam(teamId uint, playerId uint) {
	player := service.FindById(playerId)
	team := service.DB.First(&model.Team{}, teamId)
	team.Model.(*model.Team).Players = removePlayer(team.Model.(*model.Team).Players, player)
	service.DB.Save(team.Model.(*model.Team))
}

func removePlayer(players []model.Player, player model.Player) []model.Player {
	for i, p := range players {
		if p.ID == player.ID {
			return append(players[:i], players[i+1:]...)
		}
	}
	return players
}

// FindByNationality returns all players from

func (service *PlayerService) FindByNationality(nationality string) []model.Player {
	var players []model.Player
	service.DB.Where("national = ?", nationality).Find(&players)
	return players
}

// FindByPosition returns all players from a position

func (service *PlayerService) FindByPosition(position string) []model.Player {
	var players []model.Player
	service.DB.Where("position = ?", position).Find(&players)
	return players
}

// FindByAgeRange returns all players from an age range

func (service *PlayerService) FindByAgeRange(minAge uint, maxAge uint) []model.Player {
	var players []model.Player
	service.DB.Where("age BETWEEN ? AND ?", minAge, maxAge).Find(&players)
	return players
}
