package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dns = "root:@tcp(127.0.0.1:3306)/store?charset=utf8mb4&parseTime=True&loc=Local"
var DB *gorm.DB

func DbConnection() {
	var err error

	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("No se ha podido acceder a la bdd", err)
	} else {
		log.Println("Se ha conectado a la bdd")
	}
}
