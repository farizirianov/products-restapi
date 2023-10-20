package main

import (
	"github.com/farizirianov/products-restapi/configs"
	"github.com/farizirianov/products-restapi/models"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Product{})
}
