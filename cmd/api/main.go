package main

import (
	"log"

	"github.com/AdiaElena/checkout-service/core/src/handler"
	"github.com/AdiaElena/checkout-service/core/src/repository"
	"github.com/AdiaElena/checkout-service/core/src/service"
	_ "github.com/AdiaElena/checkout-service/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	pricingRepository := repository.NewPricingRepository()

	checkout := func() service.ICheckout {
		return service.NewCheckoutService(pricingRepository)
	}

	apiHandler := handler.NewHandler(checkout)

	router := gin.Default()
	router.POST("/checkout", apiHandler.Checkout)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Server listening on :8080")
	log.Fatal(router.Run(":8080"))
}
