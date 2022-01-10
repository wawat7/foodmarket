package main

import (
	"api-foodmarket/app"
	"api-foodmarket/exception"
	"api-foodmarket/helper"
	"api-foodmarket/services/product"
	"api-foodmarket/services/role"
	"api-foodmarket/services/user"
	"github.com/gin-gonic/gin"
)

func main() {

	configuration := app.New()
	db := app.NewDB(configuration)

	roleRepository := role.NewRepository(db)
	roleService := role.NewService(roleRepository)
	roleController := role.NewController(roleService)

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository, roleRepository)
	userController := user.NewController(userService)

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productController := product.NewController(productService)

	// Setup Gin
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.CustomRecovery(exception.ErrorHandler))

	// Setup Routing
	userController.Route(router)
	roleController.Route(router)
	productController.Route(router)

	// Start App
	err := router.Run(":3000")
	helper.PanicIfError(err)
}
