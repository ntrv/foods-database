package main

import "github.com/jinzhu/gorm"

// Up is executed when this migration is applied
func Up_20180715184202(txn *gorm.DB) {
	txn.Debug().Table("foods").AddUniqueIndex("ui_foods_name", "name")
}

// Down is executed when this migration is rolled back
func Down_20180715184202(txn *gorm.DB) {
	txn.Debug().Table("foods").RemoveIndex("ui_foods_name")
}
