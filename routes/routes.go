package routes

import ( 
	"github.com/gin-gonic/gin"
	"strings"
	"strconv"
	"api/99minutos/models"
	"api/99minutos/db"
	"api/99minutos/utils"
	"time"
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
	var user models.User
	var logged bool
	email, password, hasAuth := c.Request.BasicAuth()
	if hasAuth{
		logged, user = utils.LoginVerifier(email, password)
		if logged == false{
			c.JSON(404, gin.H{
				"Error": "Cannot acces this information",
			})
			return 
		}
	}else{
			c.JSON(404, gin.H{
				"Error": "Cannot acces this information",
			})
			return
		}
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
	switch {
	case total_weight <= 5:
		order.PackageSize = "S"
		break
	case total_weight <= 15:
		order.PackageSize = "M"
		break
	case total_weight <= 25:
		order.PackageSize = "L"
		break
	case total_weight > 25 && user.XlEnabled == true:
		order.PackageSize = "XL"
		break
	default:
		c.JSON(200, gin.H{
			"Message": "El peso excede los 25 kilos, por lo que deberás comunicarte a atención al cliente",
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
	})
	
}

func GetOrder(c *gin.Context){
	email, password, hasAuth := c.Request.BasicAuth()
	if hasAuth{
		logged, _ := utils.LoginVerifier(email, password)
		if logged == false{
			c.JSON(404, gin.H{
				"Error": "Cannot acces this information",
			})
			return 
		}
	}else{
		c.JSON(404, gin.H{
			"Error": "Cannot acces this information",
		})
		return
	}
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
	email, password, hasAuth := c.Request.BasicAuth()
	if hasAuth{
		logged, user := utils.LoginVerifier(email, password)
		if logged == false || user.IsAdmin == false{
			c.JSON(401, gin.H{
				"Error": "Cannot acces this information",
			})
			return 
		}
	}else{
			c.JSON(401, gin.H{
				"Error": "Cannot acces this information",
			})
			return
	}

	db.DB.First(&order, c.Param("id"))
	db.DB.Model(&order).Association("DestinationAddress").Find(&order.DestinationAddress)
	db.DB.Model(&order).Association("Products").Find(&order.Products)
	new_status := c.Param("new_status")
	if ValidStatus(new_status){
		if (new_status == "cancelado" && order.Status == "en_ruta") || (new_status == "cancelado"  && order.Status == "entregado"){
			c.JSON(404, gin.H{
				"Error": "No se puede cancelar una orden en ruta o entregada",
			})
			return
		}else if order.Status == "cancelado" && new_status != "cancelado"{
			c.JSON(404, gin.H{
				"Error" : "No se puede modificar una orden cancelada",
			})
		}else{
			if new_status == "cancelado"{
				currentTime := time.Now()
				deltaTime := currentTime.Sub(order.CreatedAt)
				if deltaTime <= time.Minute*2{
					order.Status = new_status
					order.Refund = true
					order.UpdatedAt = time.Now()
					db.DB.Save(&order)
					c.JSON(200, gin.H{
						"Order": order,
						"refund" :true,
						"Message": "La orden cumple para el reembolso",
					})
					return
				}else{
					order.Status = new_status
					order.Refund = false
					order.UpdatedAt = time.Now()
					db.DB.Save(&order)
					c.JSON(200, gin.H{
						"Order": order,
						"Refund": false,
						"Message": "La orden no cumple para el reembolso",
					})
					return
				}				
			}else{
				order.Status = new_status
				order.Refund = false
				order.UpdatedAt = time.Now()
				db.DB.Save(&order)
				c.JSON(200, gin.H{
					"Order": order,
				})
				return
			}

		}
	}else{
		c.JSON(400, gin.H{
			"Error": "This status is not valid",
			"Status" : new_status,
		})
		return
	}
}