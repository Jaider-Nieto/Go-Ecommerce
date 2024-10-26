package config

import (
	"os"

	"github.com/jaider-nieto/ecommerce-go/products-service/internal/controller"
	"github.com/jaider-nieto/ecommerce-go/products-service/internal/repository"
	"github.com/jaider-nieto/ecommerce-go/products-service/internal/service"
)

func NewContainer() *controller.ProductController {

	mongoURI := os.Getenv("MONGO_URI")
	clientMongo := InitMongoDB(mongoURI)
	clientRedis := InitRedisClient(os.Getenv("REDIS_ADR"), os.Getenv("REDIS_PASSWORD"))
	clientS3 := InitS3Client()

	productCacheRepository := repository.NewProductRedisRepository(clientRedis)
	productRepository := repository.NewProductRepository(GetMongoCollection(clientMongo, "products_db", "products"))
	s3Repository := repository.NewS3Repository(clientS3, os.Getenv("AWS_BUCKET"))
	productService := service.NewProductService(productRepository, productCacheRepository, s3Repository)
	productController := controller.NewProductController(productService)

	return productController
}
