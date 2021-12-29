package models

import (
	"fmt"
	"labqid/config"
	"net/http"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	IsCompleted bool `json:"is_completed" gorm:"default:false;"`
}

func CreateNewTransaction(transaction *Transaction) (Response, error) {
	var res Response
	db := config.GetDBInstance()

	if result := db.Create(&transaction); result.Error != nil {
		fmt.Print("error CreateNewTransaction")
		fmt.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error save new record"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = transaction

	return res, nil
}
