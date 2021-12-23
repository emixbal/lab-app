package models

import (
	"fmt"
	"labqid/app/helpers"
	"labqid/config"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	IsAdmin  bool   `json:"is_admin,omitempty" gorm:"default:false"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserRegister(user *User) (Response, error) {
	var res Response
	db := config.GetDBInstance()

	if result := db.Create(&user); result.Error != nil {
		fmt.Print("error CreateABook")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error save new record"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"

	return res, nil
}

func CheckLogin(email, passwordTxt string) (bool, string, error) {
	var user User
	db := config.GetDBInstance()

	if result := db.Where(&User{Email: email}).First(&user); result.Error != nil {
		fmt.Print(result.Error)
		return false, "", result.Error
	}

	match, _ := helpers.CheckPasswordHash(user.Password, passwordTxt)
	if !match {
		return false, "", nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    email,
		"userId":   user.ID,
		"is_admin": user.IsAdmin,
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println(err)
	}

	return true, tokenString, nil
}
