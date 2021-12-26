package models

import (
	"context"
	"fmt"
	"labqid/config"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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
	IsActive       bool   `json:"is_active" gorm:"default:true"`
	User           User
}

type ProductResponse struct {
	Name           string `json:"name" gorm:"index:idx_name,unique"`
	Description    string `json:"description"`
	Price          int16  `json:"price"`
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

func FethAllProducts() (Response, error) {
	var products_responses []ProductResponse
	var products_response ProductResponse
	var res Response

	// ======
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACEESS_ID")
	secretAccessKey := os.Getenv("MINIO_SECRET_ACEESS_ID")
	useSSL := false
	minioClient, err_minio_conn := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err_minio_conn != nil {
		log.Println(err_minio_conn)
	}
	// =======

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

		// Set request parameters for content-disposition.
		reqParams := make(url.Values)
		// Generates a presigned url which expires in a day.
		image_url, err := minioClient.PresignedGetObject(context.Background(), os.Getenv("MINIO_BUCKET_NAME"), products_response.ImageName, time.Second*24*60*60, reqParams)
		if err != nil {
			fmt.Println(err)
		}
		image_thumb_url, err_thumb := minioClient.PresignedGetObject(context.Background(), os.Getenv("MINIO_BUCKET_NAME"), products_response.ImageThumbName, time.Second*24*60*60, reqParams)
		if err_thumb != nil {
			fmt.Println(err_thumb)
		}
		products_response.ImageUrl = fmt.Sprintf("%v", image_url)
		products_response.ImageThumbUrl = fmt.Sprintf("%v", image_thumb_url)

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
