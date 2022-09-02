package routes


import ( 
	"github.com/gin-gonic/gin"
	"api/99minutos/models"
	"api/99minutos/db"
	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func ComparePassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}


func CreateUserHandler(c *gin.Context){
	var user models.User
	
	err := c.BindJSON(&user)
	user.Password, _ = HashPassword(user.Password)

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




func LoginHandler(c *gin.Context){
	var userLogin models.UserLogin
	var user models.User
	err := c.BindJSON(&userLogin)

	if err != nil{
		c.JSON(400, gin.H{
			"Error": err,
		})
		return
	}

	db.DB.Where("email = ?", userLogin.Email).First(&user)

	match := ComparePassword(userLogin.Password, user.Password)
	if match == true{
		c.JSON(200, gin.H{
			"message": "Logged in",
			"token" : "token",
		})
	}else{
		c.JSON(401, gin.H{
			"message": "Logged in",
			"token" : "token",
		})
	}
	

}


