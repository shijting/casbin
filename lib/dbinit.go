package lib

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func init()  {
	DB = initDB()
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "mysql57:shijinting0510@tcp(106.53.5.146:3306)/test?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize: 256, // default size for string fields
		DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err !=nil {
		log.Fatal(err)
	}
	return db
}
