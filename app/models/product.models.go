package models

import (
	"fmt"
	"labqid/app/helpers"
	"labqid/config"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name           string `json:"name" gorm:"index:idx_name,unique"`
	Description    string `json:"description"`
	Price          int    `json:"price"`
	ImageName      string `json:"image_name"`
	ImageThumbName string `json:"image_thumb_name"`
	UserId         int    `json:"user_id"`
	IsActive       bool   `json:"is_active" gorm:"default:true"`
	User           User
}

type ProductUpdate struct {
	Name        string `json:"name" gorm:"index:idx_name,unique"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type ProductResponse struct {
	Name           string `json:"name" gorm:"index:idx_name,unique"`
	Description    string `json:"description"`
	Price          int    `json:"price"`
	UserId         int    `json:"user_id"`
	UserName       string `json:"user_name"`
	UserEmail      string `json:"user_email"`
	ImageName      string `json:"image_name"`
	ImageThumbName string `json:"image_thumb_name"`
	ImageUrl       string `json:"image_url"`
	ImageThumbUrl  string `json:"image_thumb_url"`
}

func ActiveInActiveProduct(product_id string) (Response, error) {
	var res Response
	var product Product

	db := config.GetDBInstance()
	result := db.First(&product, product_id)
	if result.Error != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "can't find record"
		return res, result.Error
	}

	if product.IsActive {
		fmt.Println(product.IsActive)
		product.IsActive = false
	} else {
		fmt.Println(product.IsActive)
		product.IsActive = true
	}

	db.Save(&product)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = product

	return res, nil
}

func RemoveImage(product_id string) (Response, error) {
	var res Response
	var product Product

	db := config.GetDBInstance()
	result := db.First(&product, product_id)

	if result.Error != nil {
		res.Status = http.StatusOK
		res.Message = "can't find record"
		return res, result.Error
	}

	_, err_delete_img := helpers.MinioDeleteObject(product.ImageName)
	if err_delete_img != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "error delete image"
		return res, err_delete_img
	}
	_, err_delete_img_thumb := helpers.MinioDeleteObject(product.ImageThumbName)
	if err_delete_img_thumb != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "error delete image thumb"
		return res, err_delete_img_thumb
	}
	product.ImageName = ""
	product.ImageThumbName = ""
	db.Save(&product)

	return res, nil
}

func ProductDetail(product_id string) (Response, error) {
	var products_response ProductResponse
	var res Response
	var count int64

	db := config.GetDBInstance()

	rows, err_q := db.Table("products p").Where("p.id = ?", product_id).Select(
		"p.name, p.description, p.price, u.id as user_id, u.name as user_name, u.email as user_email, p.image_name, p.image_thumb_name",
	).Joins("left join users u on p.user_id = u.id").Count(&count).Rows()
	if err_q != nil {
		log.Println(err_q)
		fmt.Print("error FethAllProducts")
		fmt.Print(err_q)

		res.Status = http.StatusInternalServerError
		res.Message = "error fetchin records"
		return res, err_q
	}
	if count < 1 {
		res.Status = http.StatusOK
		res.Message = "no record found"
		return res, nil
	}

	for rows.Next() {
		db.ScanRows(rows, &products_response)
		products_response.ImageUrl, _ = helpers.MinioGetUrl(products_response.ImageName)
		products_response.ImageThumbUrl, _ = helpers.MinioGetUrl(products_response.ImageThumbName)
	}
	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = products_response
	return res, nil
}

func FethAllProducts() (Response, error) {
	var products_responses []ProductResponse
	var products_response ProductResponse
	var res Response

	db := config.GetDBInstance()

	rows, err_q := db.Table("products p").Select(
		"p.name, p.description, p.price, u.id as user_id, u.name as user_name, u.email as user_email, p.image_name, p.image_thumb_name",
	).Joins("left join users u on p.user_id = u.id").Rows()
	if err_q != nil {
		log.Println(err_q)
		fmt.Print("error FethAllProducts")
		fmt.Print(err_q)

		res.Status = http.StatusInternalServerError
		res.Message = "error fetchin records"
		return res, err_q
	}

	for rows.Next() {
		db.ScanRows(rows, &products_response)

		products_response.ImageUrl, _ = helpers.MinioGetUrl(products_response.ImageName)
		products_response.ImageThumbUrl, _ = helpers.MinioGetUrl(products_response.ImageThumbName)

		products_responses = append(products_responses, products_response)
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = products_responses

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

func UpdateProduct(product_id string, product_payload ProductUpdate) (Response, error) {
	var res Response
	var product Product

	db := config.GetDBInstance()
	result := db.First(&product, product_id)
	if result.Error != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "can't find record"
		return res, result.Error
	}

	if product_payload.Name != "" {
		product.Name = product_payload.Name
	}

	if product_payload.Description != "" {
		product.Description = product_payload.Description
	}

	if product_payload.Price != 0 {
		product.Price = product_payload.Price
	}

	db.Save(&product)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = product

	return res, nil
}
