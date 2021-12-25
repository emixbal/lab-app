package models

import (
	"fmt"
	"labqid/config"
	"net/http"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

func FethAllCategories() (Response, error) {
	var categories []Category
	var res Response

	db := config.GetDBInstance()

	if result := db.Find(&categories); result.Error != nil {
		fmt.Print("error FethAllCategories")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error fetchin records"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = categories

	return res, nil
}

func CreateACategory(category *Category) (Response, error) {
	var res Response
	db := config.GetDBInstance()

	if result := db.Create(&category); result.Error != nil {
		fmt.Print("error CreateACategory")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error save new record"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = category

	return res, nil
}
