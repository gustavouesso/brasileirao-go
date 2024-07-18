package infra

import (
	"log"

	"github.com/gustavouesso/brasileirao-go.git/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// CreateConnection creates a connection with the database
func CreateConnection() *gorm.DB {
	dsn :="host=localhost user=admin password=12345 dbname=brasileirao port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database")
	}

	db.AutoMigrate(&model.Championship{}, &model.Team{}, &model.Match{}, &model.ChampionshipFormat{})
	return db
}
	