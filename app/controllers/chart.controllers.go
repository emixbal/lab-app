package controllers

import (
	"fmt"
	"labqid/app/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func NewChart(c *fiber.Ctx) error {
	var chart models.Chart

	user_id := c.Locals("user_id")
	chart.UserId = int(user_id.(float64))
	chart.SampleName = c.FormValue("sample_name")
	chart.SampleDescription = c.FormValue("sample_description")
	chart.SampleState = c.FormValue("sample_state")
	chart.SampleWeight = c.FormValue("sample_weight")
	qty, _ := strconv.Atoi(c.FormValue("quantity"))
	product_id, _ := strconv.Atoi(c.FormValue("product_id"))
	chart.Quantity = qty
	chart.ProductId = product_id

	if chart.SampleName == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "sample_name is required"})
	}
	if chart.SampleDescription == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "sample_description is required"})
	}
	if chart.SampleState == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "sample_state is required"})
	}
	if chart.SampleWeight == "" {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "sample_weight is required"})
	}
	if chart.Quantity == 0 {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "quantity is required"})
	}
	if chart.ProductId == 0 {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": "product_id is required"})
	}
	if chart.UserId == 0 {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"message": "something went wrong"})
	}

	result, _ := models.NewChart(&chart)
	return c.Status(result.Status).JSON(result)
}

func ShowUserChart(c *fiber.Ctx) error {
	user_id := c.Locals("user_id")

	user_id_str := fmt.Sprintf("%v", user_id)
	result, _ := models.ShowUserChart(user_id_str)

	return c.Status(result.Status).JSON(result)
}

func RemoveChart(c *fiber.Ctx) error {
	result, _ := models.RemoveChart(c.Params("chart_id"))

	return c.Status(result.Status).JSON(result)
}

func UserChartDetail(c *fiber.Ctx) error {
	user_id := c.Locals("user_id")

	user_id_str := fmt.Sprintf("%v", user_id)
	result, _ := models.UserChartDetail(c.Params("chart_id"), user_id_str)

	return c.Status(result.Status).JSON(result)
}

func UserChartCheckout(c *fiber.Ctx) error {
	var charts_id []string

	user_id := c.Locals("user_id")
	user_id_str := fmt.Sprintf("%v", user_id)

	form, err_form := c.MultipartForm()

	if err_form != nil {
		log.Println(err_form)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "chart_id is required"})

	}
	chart_id_txt := form.Value["chart_id"]

	for _, v := range chart_id_txt {
		charts_id = append(charts_id, v)
	}

	result, _ := models.UserChartCheckout(user_id_str, charts_id)

	return c.Status(result.Status).JSON(result)
}
