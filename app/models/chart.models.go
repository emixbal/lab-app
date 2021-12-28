package models

import (
	"fmt"
	"labqid/config"
	"net/http"

	"gorm.io/gorm"
)

type Chart struct {
	gorm.Model
	SampleName        string `json:"sample_name"`
	SampleDescription string `json:"sample_description"`
	SampleState       string `json:"sample_state"`
	SampleWeight      string `json:"sample_weight"`
	Quantity          int    `json:"quantity"`
	ProductId         int    `json:"product_id"`
	Product           Product
	UserId            int `json:"user_id"`
	User              User
	IsActive          bool `json:"is_active" gorm:"default:true;"`
	IsCompleted       bool `json:"is_completed" gorm:"default:false;"`
}

func NewChart(chart *Chart) (Response, error) {
	var res Response
	db := config.GetDBInstance()

	if result := db.Create(&chart); result.Error != nil {
		fmt.Print("error CreateAChart")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error save new record"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = chart

	return res, nil
}
