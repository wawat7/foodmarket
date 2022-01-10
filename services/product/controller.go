package product

import (
	"api-foodmarket/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type productController struct {
	service Service
}

func NewController(service Service) *productController {
	return &productController{service: service}
}

func (controller *productController) Route(app *gin.Engine) {
	route := app.Group("api/products")
	route.GET("/", controller.List)
	route.GET("/:id", controller.GetById)
	route.POST("/", controller.Create)
	route.PUT("/:id", controller.Update)
	route.DELETE("/:id", controller.Delete)
}

func (controller *productController) List(c *gin.Context) {
	products := controller.service.FindAll()

	c.JSON(http.StatusOK, helper.APIResponse("List Product", http.StatusOK, "success", ProductsFormat(products), helper.Pagination{}))
	return
}

func (controller *productController) GetById(c *gin.Context) {
	var input GetProductDetail
	err := c.ShouldBindUri(&input)
	helper.PanicIfError(err)

	product := controller.service.FindById(input.Id)
	c.JSON(http.StatusOK, helper.APIResponse("Product Detail", http.StatusOK, "success", ProductFormat(product), helper.Pagination{}))
	return
}

func (controller *productController) Create(c *gin.Context) {
	var input CreateProductRequest
	err := c.ShouldBind(&input)
	helper.PanicIfError(err)

	product := Product{
		Name:        input.Name,
		Description: input.Description,
		Ingredient:  helper.ConvertDataToJson(input.Ingredient),
		Price:       input.Price,
		Rate:        0,
		Type:        input.Type,
		Image:       input.Image,
	}

	product = controller.service.Create(product)
	c.JSON(http.StatusOK, helper.APIResponse("Create Product Successfully", http.StatusOK, "success", ProductFormat(product), helper.Pagination{}))
	return
}

func (controller *productController) Update(c *gin.Context) {
	var inputParam GetProductDetail
	err := c.ShouldBindUri(&inputParam)
	helper.PanicIfError(err)

	var input UpdateProductRequest
	err = c.ShouldBind(&input)
	helper.PanicIfError(err)

	product := controller.service.FindById(inputParam.Id)
	product = mapRequestToProduct(input, product)
	product = controller.service.Update(product)

	c.JSON(http.StatusOK, helper.APIResponse("Update Product Successfully", http.StatusOK, "success", ProductFormat(product), helper.Pagination{}))
	return
}

func (controller *productController) Delete(c *gin.Context) {
	var input GetProductDetail
	err := c.ShouldBindUri(&input)
	helper.PanicIfError(err)

	product := controller.service.FindById(input.Id)
	controller.service.Delete(product)
	c.JSON(http.StatusOK, helper.APIResponse("Delete Product Successfully", http.StatusOK, "success", nil, helper.Pagination{}))
	return
}

func mapRequestToProduct(request UpdateProductRequest, product Product) Product {
	if request.Name != "" {
		product.Name = request.Name
	}
	if request.Description != "" {
		product.Description = request.Description
	}
	if len(request.Ingredient) > 0 {
		product.Ingredient = helper.ConvertDataToJson(request.Name)
	}
	if request.Price != 0 {
		product.Price = request.Price
	}
	if request.Type != "" {
		product.Type = request.Type
	}
	if request.Image != "" {
		product.Image = request.Image
	}

	return product
}
