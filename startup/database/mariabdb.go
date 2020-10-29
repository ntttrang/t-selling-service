package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ntttrang/t-selling-service/startup/config"
)

var mariadb *gorm.DB

func initMariaDb(configuration config.MariaDBConfig) *gorm.DB {
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configuration.Username,
		configuration.Password,
		configuration.Host,
		configuration.Port,
		configuration.Name)

	mariadb, err := gorm.Open("mysql", connectString)
	if nil != err {
		panic("DB Connection error: " + err.Error())
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	mariadb.DB().SetMaxIdleConns(20)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	mariadb.DB().SetMaxOpenConns(10)
	mariadb.LogMode(true)
	mariadb.SingularTable(false)

	// mariadb.AutoMigrate(
	// 	entity.User{},
	// )
	return mariadb
}

// GetMariaDb init mariaDB
func GetMariaDb() *gorm.DB {
	if nil == mariadb {
		mariadb = initMariaDb(config.GetMariaDbConf())
	}
	return mariadb
}
