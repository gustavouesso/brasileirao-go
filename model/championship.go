package model

type Championship struct {
	ID uint `gorm:"primary_key,auto_increment"`
	Name string `gorm:"type:varchar(100);not null"`
	Year uint `gorm:"type:smallint;not null"`
	Teams []Team `gorm:"many2many:championship_teams;"`
	Matches []Match `gorm:"foreignKey:ChampionshipID"`
	ChampionshipFormat ChampionshipFormat `gorm:"foreignKey:ChampionshipFormatID"`
}
