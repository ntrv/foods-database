package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

type food struct {
	ID        uint       `sql:"not null;primary_key"`
	CreatedAt time.Time  `sql:"not null;DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `sql:"not null;DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `sql:"index"`
	Name      string     `sql:"size:255"`
}

// Up is executed when this migration is applied
func Up_20180715023549(txn *gorm.DB) {
	txn.Debug().CreateTable(&food{})
}

// Down is executed when this migration is rolled back
func Down_20180715023549(txn *gorm.DB) {
	txn.Debug().DropTable(&food{})
}
