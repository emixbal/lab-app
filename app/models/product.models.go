package models

import (
	"fmt"
	"labqid/config"
	"net/http"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name           string `json:"name" gorm:"index:idx_name,unique"`
	Description    string `json:"description"`
	Price          int16  `json:"price"`
	ImageName      string `json:"image_name"`
	ImageThumbName string `json:"image_thumb_name"`
	UserId         int    `json:"user_id"`
	User           User
}

func FethAllProducts() (Response, error) {
	var products []Product
	var res Response

	db := config.GetDBInstance()

	if result := db.Find(&products); result.Error != nil {
		fmt.Print("error FethAllProducts")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error fetchin records"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = products

	return res, nil
}

func CreateAProduct(product *Product) (Response, error) {
	var res Response
	db := config.GetDBInstance()

	if result := db.Create(&product); result.Error != nil {
		fmt.Print("error CreateAProduct")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error save new record"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = product

	return res, nil
}
