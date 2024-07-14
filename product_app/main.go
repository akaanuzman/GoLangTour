package main

import (
	"context"
	"golangtour/common/app"
	"golangtour/common/postresql"
	"golangtour/controller"
	"golangtour/persistence"
	"golangtour/service"

	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	e := echo.New()

	configurationManager := app.NewConfigurationManager()

	dbPool := postresql.GetConnectionPool(ctx, configurationManager.PostreSqlConfig)

	productRepository := persistence.NewProductRepository(dbPool)

	productService := service.NewProductService(productRepository)

	productController := controller.NewProductController(productService)

	productController.RegisterRoutes(e)

	e.Start("localhost:8080")
}
