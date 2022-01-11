package order

import (
	"api-foodmarket/helper"
	"api-foodmarket/services/product"
	"api-foodmarket/services/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type orderController struct {
	service        Service
	userService    user.Service
	productService product.Service
}

func NewController(service Service, userService user.Service, productService product.Service) *orderController {
	return &orderController{service: service, userService: userService, productService: productService}
}

func (controller *orderController) Route(app *gin.Engine) {
	route := app.Group("api/orders")
	route.GET("/", controller.List)
	route.POST("/", controller.Create)
	route.GET("/:id", controller.FindById)
	route.PUT("/:id/update-status", controller.UpdateStatus)
}

func (controller *orderController) List(c *gin.Context) {
	orders := controller.service.FindAll()

	c.JSON(http.StatusOK, helper.APIResponse("List Order", http.StatusOK, "success", OrdersFormat(orders), helper.Pagination{}))
	return
}

func (controller *orderController) Create(c *gin.Context) {
	var productData []ProductQuantityData
	var input CreateOrderRequest

	err := c.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	user := controller.userService.FindById(input.UserId)

	product := ProductQuantityData{
		ProductId: input.ProductId,
		Quantity:  input.Quantity,
	}
	productData = append(productData, product)

	order := controller.service.Create(user, productData)

	c.JSON(http.StatusOK, helper.APIResponse("Create Order Successfully", http.StatusOK, "success", OrderFormat(order), helper.Pagination{}))
	return
}

func (controller *orderController) FindById(c *gin.Context) {
	var input GetOrderDetail
	err := c.ShouldBindUri(&input)
	helper.PanicIfError(err)

	order := controller.service.FindById(input.Id)
	c.JSON(http.StatusOK, helper.APIResponse("Order Detail", http.StatusOK, "success", OrderFormat(order), helper.Pagination{}))
	return
}

func (controller *orderController) UpdateStatus(c *gin.Context) {
	var input UpdateStatusOrderRequest
	var inputParam GetOrderDetail

	err := c.ShouldBindJSON(&input)
	helper.PanicIfError(err)

	err = c.ShouldBindUri(&inputParam)
	helper.PanicIfError(err)

	user := controller.userService.FindById(input.UserId)

	order := controller.service.FindById(inputParam.Id)
	order.Status = input.Status
	order = controller.service.update(order, user)

	c.JSON(http.StatusOK, helper.APIResponse("Update Status Order Successfully", http.StatusOK, "success", OrderFormat(order), helper.Pagination{}))
	return
}
