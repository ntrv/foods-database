package models

type Food struct {
	Model
	Name string `json:name gorm:"size:255"`
}
