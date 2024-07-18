package model

type Coach struct {
	ID uint `gorm:"primary_key,auto_increment"`
	Name string `gorm:"type:varchar(100);not null"`
	Age uint `gorm:"type:smallint;not null"`
	Nationality string `gorm:"type:varchar(100);not null"`
	Teams []Team `gorm:"foreignKey:CoachID"`
}
