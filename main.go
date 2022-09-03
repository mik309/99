package main

import ( 
	//"net/http"
	"github.com/gin-gonic/gin"
	//"errors"
	"api/99minutos/routes"
	"api/99minutos/db"
	"api/99minutos/models"
	"github.com/joho/godotenv"
	"os"
	"log"
)





func main(){

	err := godotenv.Load()
  	if err != nil {
    	log.Fatal(err)
  	}
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("PORT")

	db.Connection(host,user,password,dbname,port)
	//Order migrations
	db.DB.AutoMigrate(models.Order{}, models.Product{}, models.Address{})
	db.DB.AutoMigrate(models.User{})

	r := gin.Default()
	//Orders
	r.POST("/orders/create", routes.CreateOrder)
	r.GET("/orders/:id", routes.GetOrder)
	r.PUT("/orders/:id/:new_status", routes.UpdateOrderStatus)

	
	//Users
	r.POST("/users/create", routes.CreateUserHandler)
	//r.GET("/users/login", routes.LoginHandler)
	r.Run()
}

