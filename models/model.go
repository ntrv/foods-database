package models

import "time"

type Model struct {
	ID        uint `gorm:"not null;primary_key" json:-`
	CreatedAt time.Time `gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP" json:-`
	UpdatedAt time.Time `gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:-`
	DeletedAt *time.Time `sql:"index" json:-`
}
