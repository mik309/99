package db

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
  )

var DSN = "host=ec2-34-231-221-151.compute-1.amazonaws.com user=sdrvtcalonmcbg password=3e7de89ff1e97f62f8119350715f8a26115b3de63b585fe828dbd9c8fd3d77a8 dbname=df1eudoqib2k4 port=5432"

var DB *gorm.DB


func Connection(){
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil{
		log.Fatal(err)
	}else{
		log.Println("DB connected")
	}
}

