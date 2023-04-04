package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DIALECT             = "mysql"
	MYSQL_USER_NAME     = "root"
	MYSQL_USER_PASSWORD = "rootroot"
	MYSQL_DB_NAME       = "springdb"
	MYSQL_REQUIREMENTS  = "charset=utf8&parseTime=True&loc=Local"
)

var (
	db *gorm.DB
	// "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	dbConnectData = MYSQL_USER_NAME + ":" + MYSQL_USER_PASSWORD + "@/" + MYSQL_DB_NAME + "?" + MYSQL_REQUIREMENTS
)

func Connect() {
	d, err := gorm.Open(DIALECT, dbConnectData)

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
