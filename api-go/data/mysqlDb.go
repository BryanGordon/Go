package data

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "root:@tcp(127.0.0.1:3306)/moviebd?charset=utf8mb4&parseTime=True&loc=Local"

var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func ConnectionDB() {
	if err != nil {
		log.Fatal("No se ha podido conectar a la base de datos", err)
	} else {
		log.Println("Se ha conectado a la bdd")
	}
}
