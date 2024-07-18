package model 

type Match struct {
	ID uint `gorm:"primary_key,auto_increment"`
	HomeTeam Team `gorm:"foreignKey:HomeTeamID"`
	AwayTeam Team `gorm:"foreignKey:AwayTeamID"`
	HomeTeamScore uint `gorm:"type:smallint;not null"`
	AwayTeamScore uint `gorm:"type:smallint;not null"`
	ChampionshipID uint `gorm:"not null"`
	Stadium string `gorm:"type:varchar(100);not null"`
}