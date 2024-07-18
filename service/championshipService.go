package service

import (
	"github.com/gustavouesso/brasileirao-go.git/model"
	"gorm.io/gorm"
)

// ChampionshipService provides methods to manipulate the championship model

type ChampionshipService struct {
	DB *gorm.DB
}

// NewChampionshipService creates a new championship service

func NewChampionshipService(db *gorm.DB) ChampionshipService {
	return ChampionshipService{DB: db}
}

// Save persists a championship in the database

func (service *ChampionshipService) Save(championship model.Championship) model.Championship {
	service.DB.Create(&championship)
	return championship
}

// FindAll returns all championships from the database

func (service *ChampionshipService) FindAll() []model.Championship {
	var championships []model.Championship
	service.DB.Preload("Teams").Preload("Matches").Preload("ChampionshipFormat").Find(&championships)
	return championships
}

// FindById returns a championship by its ID

func (service *ChampionshipService) FindById(id uint) model.Championship {
	var championship model.Championship
	service.DB.Preload("Teams").Preload("Matches").Preload("ChampionshipFormat").First(&championship, id)
	return championship
}

// Update updates a championship in the database

func (service *ChampionshipService) Update(id uint, championship model.Championship) model.Championship {
	service.DB.Save(&championship)
	return championship
}

// Delete removes a championship from the database

func (service *ChampionshipService) Delete(id uint) {
	service.DB.Delete(&model.Championship{}, id)
}
