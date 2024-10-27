package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jaider-nieto/ecommerce-go/products-service/internal/controller"
	"github.com/jaider-nieto/ecommerce-go/products-service/internal/middleware"
)

func ProductRoutes(router *gin.Engine, productsController *controller.ProductController) {
	productGroup := router.Group("/products")
	{
		productGroup.GET("/", productsController.GetProducts)
		productGroup.GET("/product_id", productsController.GetProduct)
	}

	//rutas de admin o protegidas
	adminGroup := productGroup.Group("admin")
	adminGroup.Use(middleware.AuthMiddleware)
	{
		adminGroup.POST("/", productsController.PostProduct)
		adminGroup.PATCH("/:product_id", productsController.UpdateProduct)
		adminGroup.DELETE("/:product_id", productsController.DeleteProduct)
		adminGroup.POST("/upload-file", productsController.UploadFile)
	}

	// 		GET /search
	// 		GET /category/:category_id
	// 		POST /:product_id/reviews
	// 		GET /:product_id/reviews
	// 		GET /featured
	// 		GET /:product_id/related

}
