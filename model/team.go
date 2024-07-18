package model

type Team struct {
	ID uint `gorm:"primary_key,auto_increment"`
	Name string `gorm:"type:varchar(100);not null"`
	Abbreviation string `gorm:"type:varchar(10);not null"`
	Founded uint `gorm:"type:smallint;not null"`
	Stadium string `gorm:"type:varchar(100);not null"`
	City string `gorm:"type:varchar(100);not null"`
	State string `gorm:"type:varchar(2);not null"`
	Players []Player `gorm:"foreignKey:TeamID"`
	Coach Coach `gorm:"foreignKey:CoachID"`
}