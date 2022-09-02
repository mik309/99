package main

import ( 
	//"net/http"
	"github.com/gin-gonic/gin"
	//"errors"
	"api/99minutos/routes"
	"api/99minutos/db"
	"api/99minutos/models"
)





func main(){

	db.Connection()
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
	r.GET("/users/login", routes.LoginHandler)
	r.Run()
}

