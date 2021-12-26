package controllers

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"labqid/app/helpers"
	"labqid/app/models"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nfnt/resize"
)

func FetchAllProducts(c *fiber.Ctx) error {
	result, _ := models.FethAllProducts()
	return c.Status(result.Status).JSON(result)
}

func ActiveInActiveProduct(c *fiber.Ctx) error {
	result, _ := models.ActiveInActiveProduct(c.FormValue("product_id"))
	return c.Status(result.Status).JSON(result)
}

func UpdateProduct(c *fiber.Ctx) error {
	var product models.ProductUpdate

	product_id := c.FormValue("product_id")
	product.Name = c.FormValue("name")
	product.Description = c.FormValue("description")
	price, _ := strconv.Atoi(c.FormValue("price"))
	product.Price = price

	result, _ := models.UpdateProduct(product_id, product)
	return c.Status(result.Status).JSON(result)
}

func ProductDetail(c *fiber.Ctx) error {
	result, _ := models.ProductDetail(c.FormValue("product_id"))
	return c.Status(result.Status).JSON(result)
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	var img image.Image
	var err_decode error
	var err_encode error

	user_id := c.Locals("user_id")
	product.UserId = int(user_id.(float64))
	product.Name = c.FormValue("name")
	product.Description = c.FormValue("description")
	price, _ := strconv.Atoi(c.FormValue("price"))
	product.Price = price

	file, err_file := c.FormFile("image")
	if err_file != nil {
		log.Println(err_file)
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "image is required"})
	}
	if product.Name == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "name is required"})
	}
	if product.Price == 0 {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "price is required"})
	}

	final_image_file_name := helpers.ReplaceSpace(product.Name)
	final_image_file_thumb_name := helpers.ReplaceSpace(product.Name) + "_thumb"
	fileExtension := filepath.Ext(file.Filename)
	file_path_ori := os.Getenv("PATH_UPLOAD_TMP") + "/" + file.Filename
	file_path_renamed := os.Getenv("PATH_UPLOAD_TMP") + "/" + final_image_file_name + fileExtension
	file_thumb_path_renamed := os.Getenv("PATH_UPLOAD_TMP") + "/" + final_image_file_thumb_name + fileExtension

	if err_upload := c.SaveFile(file, fmt.Sprintf(file_path_ori)); err_upload != nil {
		log.Println(err_upload)
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"message": "Something went wrong!"})
	}

	product_image_file, _ := os.Open(file_path_ori)
	if err_rename := os.Rename(file_path_ori, file_path_renamed); err_rename != nil {
		log.Println(err_rename)
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"message": "Something went wrong!"})
	}

	switch fileExtension {
	case ".jpeg", ".jpg":
		img, err_decode = jpeg.Decode(product_image_file)
	case ".png":
		img, err_decode = png.Decode(product_image_file)
	default:
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "file type is not allowed!"})
	}

	if err_decode != nil {
		log.Fatal(err_decode)
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"message": "Somethinge went wrong"})
	}

	m := resize.Resize(270, 0, img, resize.Lanczos3)

	out, errCreate := os.Create(file_thumb_path_renamed)
	if errCreate != nil {
		log.Fatal(errCreate)
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"message": "Somethinge went wrong"})
	}

	switch fileExtension {
	case ".png":
		err_encode = png.Encode(out, m)
	default:
		err_encode = jpeg.Encode(out, m, nil)
	}

	if err_encode != nil {
		log.Fatal(err_encode)
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"message": "Somethinge went wrong"})
	}

	_, err_minio := helpers.MinioUpload(file_path_renamed, final_image_file_name, fileExtension)
	if err_minio != nil {
		log.Println(err_minio)
	}

	_, err_minio_thumb := helpers.MinioUpload(file_path_renamed, final_image_file_thumb_name, fileExtension)
	if err_minio_thumb != nil {
		log.Println(err_minio_thumb)
	}

	if _, err := os.Stat(file_path_renamed); err == nil {
		if err := os.Remove(file_path_renamed); err != nil {
			log.Println(err)
		}
	}

	if _, err := os.Stat(file_thumb_path_renamed); err == nil {
		if err := os.Remove(file_thumb_path_renamed); err != nil {
			log.Println(err)
		}
	}

	if _, err := os.Stat(file_path_ori); err == nil {
		if err := os.Remove(file_path_ori); err != nil {
			log.Println(err)
		}
	}

	product.ImageName = final_image_file_name + fileExtension
	product.ImageThumbName = final_image_file_thumb_name + fileExtension

	result, _ := models.CreateAProduct(&product)
	return c.Status(result.Status).JSON(result)
}
