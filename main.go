package main

import ( 
	"github.com/gin-gonic/gin"
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
	host := os.Getenv("DBHOST")
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("DBPORT")

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
	r.Run()
}

