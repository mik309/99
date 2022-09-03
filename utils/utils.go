package utils

import (
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


func LoginVerifier(email string, pw string) (bool, models.User) {
	var user, nullUser models.User 

	db.DB.Where("email = ?", email).First(&user)

	match := ComparePassword(pw, user.Password)
	if match == true{
		return true, user
	}else{
		return false, nullUser
	}
	

}
