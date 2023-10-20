package main

import (
	"log"

	"github.com/farizirianov/products-restapi/configs"
	"github.com/farizirianov/products-restapi/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	r := gin.Default()

	routes.NewRouter(r)

	r.Run()
	if err := r.Run("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}
