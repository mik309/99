package routes


import ( 
	"github.com/gin-gonic/gin"
	"api/99minutos/models"
	"api/99minutos/db"
	"api/99minutos/utils"
)



func CreateUserHandler(c *gin.Context){
	var user models.User
	
	err := c.BindJSON(&user)
	user.Password, _ = utils.HashPassword(user.Password)

	if err != nil{
		c.JSON(400, gin.H{
			"Error": err,
		})
	}

	createdUser := db.DB.Create(&user)
	if createdUser.Error != nil {
		c.JSON(400, gin.H{
			"Error": createdUser.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Account Created Succesfully",
	})
}







