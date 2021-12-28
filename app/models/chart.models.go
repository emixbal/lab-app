package models

import (
	"fmt"
	"labqid/config"
	"log"
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

type ChartResponse struct {
	Id                string `json:"id"`
	SampleName        string `json:"sample_name"`
	SampleDescription string `json:"sample_description"`
	SampleState       string `json:"sample_state"`
	SampleWeight      int    `json:"sample_weight"`
	Quantity          int    `json:"quantity"`
	ProductName       string `json:"product_name"`
	ProductId         int    `json:"product_id"`
	ProductPrice      int    `json:"product_price"`
	IsActive          bool   `json:"is_active" gorm:"default:true;"`
	IsCompleted       bool   `json:"is_completed" gorm:"default:false;"`
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

func ShowUserChart(user_id string) (Response, error) {
	var res Response
	var chart_response ChartResponse
	var charts_response []ChartResponse

	db := config.GetDBInstance()

	rows, err_q := db.Table("charts c").Where("c.user_id = ?", user_id).Where("c.is_active = ?", true).Select(
		"c.id, c.sample_name, c.sample_description, c.sample_state, c.sample_weight, c.quantity, p.name AS product_name, p.id AS product_id, p.price AS product_price, c.is_active, c.is_completed",
	).Joins("left join products p on c.product_id = p.id").Rows()

	if err_q != nil {
		log.Println(err_q)
		fmt.Print("error FethAllProducts")
		fmt.Print(err_q)

		res.Status = http.StatusInternalServerError
		res.Message = "error fetchin records"
		return res, err_q
	}

	for rows.Next() {
		db.ScanRows(rows, &chart_response)

		charts_response = append(charts_response, chart_response)
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = charts_response

	return res, nil
}
