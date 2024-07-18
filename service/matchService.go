package service

import (
	"github.com/gustavouesso/brasileirao-go.git/model"
	"gorm.io/gorm"
)

type MatchService struct {
	DB *gorm.DB
}

func NewMatchService(db *gorm.DB) MatchService {
	return MatchService{DB: db}
}

// Save persists a championship in the database

func (service *MatchService) Save(match model.Match) model.Match {
	service.DB.Create(&match)
	return match
}

// FindAll returns all championships from the database

func (service *MatchService) FindAll() []model.Match {
	var matches []model.Match
	service.DB.Preload("HomeTeam").Preload("AwayTeam").Find(&matches)
	return matches
}

// FindById returns a championship by its ID

func (service *MatchService) FindById(id uint) model.Match {
	var match model.Match
	service.DB.Preload("HomeTeam").Preload("AwayTeam").First(&match, id)
	return match
}

// Update updates a championship in the database

func (service *MatchService) Update(id uint, match model.Match) model.Match {
	service.DB.Save(&match)
	return match
}

// Delete removes a championship from the database

func (service *MatchService) Delete(id uint) {
	service.DB.Delete(&model.Match{}, id)
}
