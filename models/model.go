package models

import "time"

type Model struct {
	ID        uint `gorm:"not null;primary_key"`
	CreatedAt time.Time `gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `sql:"index"`
}
