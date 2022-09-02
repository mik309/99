package routes

import ( 
	//"net/http"
	"github.com/gin-gonic/gin"
	"strings"
	"strconv"
	"api/99minutos/models"
	"api/99minutos/db"
	"time"
	//"errors"
	//"io/ioutil"
)


func ValidateCoord(coordinates string) bool {
	coordinates = strings.Replace(coordinates, " ", "", -1)
	coords := strings.Split(coordinates, ",")
	lat, err_lat := strconv.ParseFloat(coords[0], 64)
	lon, err_lon := strconv.ParseFloat(coords[1], 64)
	if err_lat != nil || err_lon != nil || lat > 90 || lat < -90 || lon > 180 || lon < -180{
		return false
	}else{
		return true
	}

}


func CreateOrder(c *gin.Context){

	var order, last_order models.Order 
	
	err := c.BindJSON(&order)
	if err != nil{
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
	}

	order.Status = "creado"
	//Setting current times
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	order.DestinationAddress.CreatedAt = time.Now()
	order.DestinationAddress.UpdatedAt = time.Now()

	products := order.Products
	var total_weight float64

	db.DB.Last(&last_order)
	last_id := last_order.ID

	for i:=0; i< len(products); i++{
		order.Products[i].CreatedAt = time.Now()
		order.Products[i].UpdatedAt = time.Now()
		order.Products[i].OrderID = last_id + 1
		total_weight += products[i].Weight
	}

	if total_weight <= 5{
		order.PackageSize = "S"
	}else if total_weight <= 15{
		order.PackageSize = "M"
	}else if total_weight <= 25{
		order.PackageSize = "L"
	}else {
		c.JSON(200, gin.H{
			"Message": "El peso excede los 25 kilos, por lo que deberás comunicarte atención al cliente",
			"Total_weight" : total_weight,
		})
		return
	}

	createdOrder := db.DB.Create(&order)

	if createdOrder.Error != nil{
		c.JSON(200, gin.H{
			"Error" : createdOrder.Error,
		})
		return
	}


	c.JSON(200, gin.H{
		"Order": order,
		"Total_weight" : total_weight,
		"last" : last_id,
	})
	
}

func GetOrder(c *gin.Context){
	var order models.Order
	db.DB.First(&order, c.Param("id"))
	db.DB.Model(&order).Association("DestinationAddress").Find(&order.DestinationAddress)
	db.DB.Model(&order).Association("Products").Find(&order.Products)

	if order.ID == 0{
		c.JSON(404, gin.H{
			"Error": "Order not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"Order": order,
	})
}

func ValidStatus(status string) bool {
	valid_statuses := []string {"creado", "recolectado", "en_estacion", "en_ruta", "entregado", "cancelado"}

	for _, item := range valid_statuses {
		if item == status{
			return true
		}
	}

	return false
}


func UpdateOrderStatus(c *gin.Context){
	var order models.Order
	db.DB.First(&order, c.Param("id"))
	db.DB.Model(&order).Association("DestinationAddress").Find(&order.DestinationAddress)
	db.DB.Model(&order).Association("Products").Find(&order.Products)
	new_status := c.Param("new_status")
	if ValidStatus(new_status){
		if new_status == "cancelado" && (order.Status == "en_ruta" || order.Status == "entregada"){
			c.JSON(404, gin.H{
				"Error": "No se puede cancelar una orden en ruta o entregada",
			})
			return
		}else{
			order.Status = new_status
			order.UpdatedAt = time.Now()
			db.DB.Save(&order)
			c.JSON(404, gin.H{
				"Order": order,
			})
		}
	}
	
	order.Status = new_status
}