package service

import (
	"github.com/gustavouesso/brasileirao-go.git/model"
	"gorm.io/gorm"
)

// CoachService provides methods to manipulate the coach model

type CoachService struct {
	DB *gorm.DB
}

// NewCoachService creates a new coach service

func NewCoachService(db *gorm.DB) CoachService {

	return CoachService{DB: db}
}

// Save persists a coach in the database

func (service *CoachService) Save(coach model.Coach) model.Coach {
	
	service.DB.Create(&coach)
	return coach
}

// FindAll returns all coaches from the database

func (service *CoachService) FindAll() []model.Coach {

	var coaches []model.Coach
	service.DB.Find(&coaches)
	return coaches
}

// FindById returns a coach by its ID

func (service *CoachService) FindById(id uint) model.Coach {
	
	var coach model.Coach
	service.DB.First(&coach, id)
	return coach
}

// Update updates a coach in the database

func (service *CoachService) Update(id uint, coach model.Coach) model.Coach {
	
	service.DB.Save(&coach)
	return coach
}

// Delete removes a coach from the database

func (service *CoachService) Delete(id uint) {
	
	service.DB.Delete(&model.Coach{}, id)
}

// FindByTeam returns all coaches from a team

func (service *CoachService) FindByTeam(teamId uint) []model.Coach {
	
	var coaches []model.Coach
	service.DB.Joins("JOIN team_coaches ON coaches.id = team_coaches.coach_id").Where("team_coaches.team_id = ?", teamId).Find(&coaches)
	return coaches
}

// AddToTeam adds a coach to a team

func (service *CoachService) AddToTeam(teamId uint, coachId uint) {

	coach := service.FindById(coachId)
}
 
// RemoveFromTeam removes a coach from a team

func (service *CoachService) RemoveFromTeam(teamId uint, coachId uint) {

	coach := service.FindById(coachId)
	team := service.DB.First(&model.Team{}, teamId)
	team.Model.(*model.Team).Coaches = append(team.Model.(*model.Team).Coaches, coach)
	service.DB.Save(team.Model.(*model.Team))
}
