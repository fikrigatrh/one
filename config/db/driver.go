package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"number1/config/env"
	"number1/models"
)

var DB *gorm.DB

func init() {
	var err error
	var configDb = env.Config

	if DB != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			log.Println("recovery error")
		}
	}()

	urlDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", configDb.DBUsername, configDb.DBPassword, configDb.DBHost, configDb.DBPort, configDb.DBName)
	//sqlServer := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", configDb.DBUserSqlServer, configDb.DBPassword, configDb.DBHostSqlServer, configDb.DBPortSqlServer, configDb.DBNameSqlServer)

	if configDb.DriverName == "postgres" {
		DB, err = gorm.Open(postgres.Open(urlDB), &gorm.Config{})
		if err != nil {
			log.Fatal("cannot connect to DB")
		}
	} else {
		DB, err = gorm.Open(mysql.Open(urlDB), &gorm.Config{})
		if err != nil {
			log.Fatal("cannot connect to DB")
		}
	}

	log.Print("success connect to db")

	DB.Debug().Migrator().DropTable(models.Auth{})
	err = DB.Debug().AutoMigrate(models.Auth{})
	if err != nil {
		return
	}

}
