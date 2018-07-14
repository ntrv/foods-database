package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type dataSource struct {
	user    string
	pass    string
	host    string
	port    int
	db_name string
	option  string
}

func dataSourceName(config *dataSource) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		config.user,
		config.pass,
		config.host,
		config.port,
		config.db_name,
		config.option,
	)
}

func NewDBConnect() (*gorm.DB, error) {
	driverName := "mysql"

	db, err := gorm.Open(driverName, dataSourceName(
		&dataSource{
			user:    "root",
			pass:    "",
			host:    "127.0.0.1",
			port:    3306,
			db_name: "foods_db",
			option:  "parseTime=true",
		},
	))
	if err != nil {
		return nil, err
	}
	return db, nil
}
