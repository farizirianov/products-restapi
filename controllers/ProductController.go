package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/farizirianov/products-restapi/configs"
	"github.com/farizirianov/products-restapi/data/request"
	"github.com/farizirianov/products-restapi/helper"
	"github.com/farizirianov/products-restapi/models"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func ProductCreate(c *gin.Context) {

	body := request.CreateProductRequestBody{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.NewValidatorError(err))
		return
	}
	product := &models.Product{Sku: body.Sku, Name: body.Name, Brand: body.Brand, Size: body.Size, Price: body.Price, MainImageUrl: body.MainImageUrl, OtherImagesUrl: body.OtherImagesUrl}

	result := configs.DB.Create(&product)

	if result.Error != nil {
		// verifica si el error es de tipo mysql.MySQLError
		var mysqlErr *mysql.MySQLError
		if errors.As(result.Error, &mysqlErr) {
			if mysqlErr.Number == 1062 && mysqlErr.Message == fmt.Sprintf("Duplicate entry '%s' for key 'PRIMARY'", product.Sku) {
				c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": fmt.Sprintf("The product with id %s already exists", product.Sku)})
				return
			}
		}
	}

	if result.Error == nil {
		c.JSON(http.StatusCreated, &product)
	}
}

func GetProducts(c *gin.Context) {
	var products []models.Product
	err := configs.DB.Find(&products).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	if err == nil {
		c.JSON(http.StatusOK, &products)
	}
}

func GetProductBySku(c *gin.Context) {
	sku := c.Param("sku")
	fmt.Println(sku)
	var product models.Product
	result := configs.DB.Where("sku = ?", sku).First(&product)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error": "Product not found"})
		return
	}

	c.JSON(200, &product)
}

func UpdateProduct(c *gin.Context) {
	sku := c.Param("sku")
	var product models.Product
	configs.DB.Where("sku = ?", sku).First(&product)

	body := request.UpdateProductRequestBody{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.NewValidatorError(err))
		return
	}

	data := &models.Product{Name: body.Name, Brand: body.Brand, Size: body.Size, Price: body.Price, MainImageUrl: body.MainImageUrl, OtherImagesUrl: body.OtherImagesUrl}

	result := configs.DB.Model(&product).Updates(data)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": "Product not found"})
		return
	}

	if result.Error == nil {
		c.JSON(http.StatusOK, &product)
	}
}

func DeleteProduct(c *gin.Context) {
	sku := c.Param("sku")
	var product models.Product
	result := configs.DB.Where("sku = ?", sku).Delete(&product)
	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error": "Product not found"})
		return
	}

	if result.RowsAffected > 0 {
		c.Status(http.StatusNoContent)
	}
}
