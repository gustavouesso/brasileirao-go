package service

import (
	"github.com/gustavouesso/brasileirao-go.git/model"
	"gorm.io/gorm"
)

type TeamService struct {
	DB *gorm.DB
}

func NewTeamService(db *gorm.DB) TeamService {
	return TeamService{DB: db}
}

// Save persists a championship in the database

func (service *TeamService) Save(team model.Team) model.Team {
	service.DB.Create(&team)
	return team
}

// FindAll returns all championships from the database

func (service *TeamService) FindAll() []model.Team {
	var teams []model.Team
	service.DB.Find(&teams)
	return teams
}

// FindById returns a championship by its ID

func (service *TeamService) FindById(id uint) model.Team {
	var team model.Team
	service.DB.First(&team, id)
	return team
}

// Update updates a championship in the database

func (service *TeamService) Update(id uint, team model.Team) model.Team {

	service.DB.Save(&team)
	return team
}

// Delete removes a championship from the database

func (service *TeamService) Delete(id uint) {
	service.DB.Delete(&model.Team{}, id)
}

// FindByChampionship returns all teams from a championship

func (service *TeamService) FindByChampionship(championshipId uint) []model.Team {
	var teams []model.Team
	service.DB.Joins("JOIN championship_teams ON teams.id = championship_teams.team_id").Where("championship_teams.championship_id = ?", championshipId).Find(&teams)
	return teams
}

// AddToChampionship adds a team to a championship

func (service *TeamService) AddToChampionship(championshipId uint, teamId uint) {
	team := service.FindById(teamId)
	championship := model.Championship{}
	service.DB.First(&championship, championshipId)
	service.DB.Model(&championship).Association("Teams").Append(&team)
}

// RemoveFromChampionship removes a team from a championship

func (service *TeamService) RemoveFromChampionship(championshipId uint, teamId uint) {
	team := service.FindById(teamId)
	championship := model.Championship{}
	service.DB.First(&championship, championshipId)
	service.DB.Model(&championship).Association("Teams").Delete(&team)
}

// FindByMatch returns all teams from a match

func (service *TeamService) FindByMatch(matchId uint) []model.Team {
	var teams []model.Team
	service.DB.Joins("JOIN matches ON teams.id = matches.home_team_id OR teams.id = matches.away_team_id").Where("matches.id = ?", matchId).Find(&teams)
	return teams
}

// FindByPlayer returns all teams from a player

func (service *TeamService) FindByPlayer(playerId uint) []model.Team {
	var teams []model.Team
	service.DB.Joins("JOIN team_players ON teams.id = team_players.team_id").Where("team_players.player_id = ?", playerId).Find(&teams)
	return teams
}

// FindByCoach returns all teams from a coach

func (service *TeamService) FindByCoach(coachId uint) []model.Team {
	var teams []model.Team
	service.DB.Joins("JOIN team_coaches ON teams.id = team_coaches.team_id").Where("team_coaches.coach_id = ?", coachId).Find(&teams)
	return teams
}

// FindByStadium returns all teams from a stadium

func (service *TeamService) FindByStadium(stadiumId uint) []model.Team {
	var teams []model.Team
	service.DB.Joins("JOIN stadiums ON teams.stadium_id = stadiums.id").Where("stadiums.id = ?", stadiumId).Find(&teams)
	return teams
}

// FindByCity returns all teams from a city

func (service *TeamService) FindByCity(cityId uint) []model.Team {
	var teams []model.Team
	service.DB.Joins("JOIN stadiums ON teams.stadium_id = stadiums.id").Where("stadiums.city_id = ?", cityId).Find(&teams)
	return teams
}

// FindByState returns all teams from a state

func (service *TeamService) FindByState(stateId uint) []model.Team {
	var teams []model.Team
	service.DB.Joins("JOIN stadiums ON teams.stadium_id = stadiums.id").Where("stadiums.city_id = ?", cityId).Find(&teams)
	return teams
}
