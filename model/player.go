package model

type Player struct {
	ID uint `gorm:"primary_key,auto_increment"`
	Name string `gorm:"type:varchar(100);not null"`
	Position string `gorm:"type:varchar(100);not null"`
	Number uint `gorm:"type:smallint;not null"`
	TeamID uint `gorm:"not null"`
	Height uint `gorm:"type:smallint;not null"`
	Weight uint `gorm:"type:smallint;not null"`
	Age uint `gorm:"type:smallint;not null"`
	Foot string `gorm:"type:varchar(10);not null"`
	Nationality string `gorm:"type:varchar(100);not null"`
	MarketValue uint `gorm:"type:smallint;not null"`
}
