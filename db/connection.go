package db

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
  )


var DB *gorm.DB


func Connection(host string, user string, password string, dbname string, port string){
	var err error
	
	DSN := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", host, user, password, dbname, port)
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	
	if err != nil{
		log.Fatal(err)
	}else{
		log.Println("DB connected")
	}
}

