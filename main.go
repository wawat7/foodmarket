package main

import (
	"api-foodmarket/app"
	"api-foodmarket/exception"
	"api-foodmarket/helper"
	"api-foodmarket/services/order"
	"api-foodmarket/services/product"
	"api-foodmarket/services/role"
	"api-foodmarket/services/user"
	"github.com/gin-gonic/gin"
)

func main() {

	configuration := app.New()
	db := app.NewDB(configuration)

	app.SeederRun()

	roleRepository := role.NewRepository(db)
	roleService := role.NewService(roleRepository)
	roleController := role.NewController(roleService)

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository, roleRepository)
	userController := user.NewController(userService)

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productController := product.NewController(productService)

	orderRepository := order.NewRepository(db)
	orderService := order.NewService(orderRepository, productService)
	orderController := order.NewController(orderService, userService, productService)

	// Setup Gin
	if configuration.Get("APP_MODE") == "RELEASE" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.CustomRecovery(exception.ErrorHandler))

	// Setup Routing
	userController.Route(router)
	roleController.Route(router)
	productController.Route(router)
	orderController.Route(router)

	// Start App
	err := router.Run(":3000")
	helper.PanicIfError(err)
}
