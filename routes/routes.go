package routes

import (
	"net/http"

	"github.com/farizirianov/products-restapi/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine) {

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	baseRouter := router.Group("/api/v1")
	productsRouter := baseRouter.Group("/products")
	productsRouter.POST("", controllers.ProductCreate)
	productsRouter.GET("", controllers.GetProducts)
	productsRouter.GET("/:sku", controllers.GetProductBySku)
	productsRouter.PUT("/:sku", controllers.UpdateProduct)
	productsRouter.DELETE("/:sku", controllers.DeleteProduct)

}
