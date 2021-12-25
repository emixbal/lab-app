package models

import (
	"fmt"
	"labqid/config"
	"net/http"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name           string `json:"name"`
	Description    string `json:"description"`
	Price          int16  `json:"price"`
	ImageLink      string `json:"image_link"`
	ImageThumbLink string `json:"image_thumb_link"`
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
