package user

import (
	"api-foodmarket/helper"
	"api-foodmarket/services/role"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	service Service
}

func NewController(service Service) *UserController {
	return &UserController{service: service}
}

func (controller *UserController) Route(app *gin.Engine) {
	route := app.Group("api/users")
	route.GET("/", controller.List)
	route.GET("/:id", controller.GetById)
	route.POST("/", controller.Create)
	route.PUT("/:id", controller.Update)
	route.DELETE("/:id", controller.Delete)
}

func (controller *UserController) List(c *gin.Context) {
	users := controller.service.FindAll()
	c.JSON(http.StatusOK, helper.APIResponse("List user", http.StatusOK, "success", UsersFormat(users), helper.Pagination{}))
	return
}

func (controller *UserController) GetById(c *gin.Context) {
	var input GetUserDetailParam
	err := c.ShouldBindUri(&input)
	helper.PanicIfError(err)

	user := controller.service.FindById(input.Id)
	c.JSON(http.StatusOK, helper.APIResponse("user detail", http.StatusOK, "success", UserFormat(user), helper.Pagination{}))
	return
}

func (controller *UserController) Create(c *gin.Context) {
	var input CreateUserRequest
	err := c.ShouldBind(&input)
	helper.PanicIfError(err)

	user := User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
		Address:  input.Address,
		City:     input.City,
		Photo:    input.Photo,
	}

	user = controller.service.Create(user, role.Customer)
	c.JSON(http.StatusOK, helper.APIResponse("Create User Successfully", http.StatusOK, "success", UserFormat(user), helper.Pagination{}))
	return
}

func (controller *UserController) Update(c *gin.Context) {
	var inputParam GetUserDetailParam
	err := c.ShouldBindUri(&inputParam)
	helper.PanicIfError(err)

	var input UpdateUserRequest
	err = c.ShouldBind(&input)
	helper.PanicIfError(err)

	user := controller.service.FindById(inputParam.Id)
	user = mapUserUpdate(input, user)

	user = controller.service.Update(user)
	c.JSON(http.StatusOK, helper.APIResponse("Create User Successfully", http.StatusOK, "success", UserFormat(user), helper.Pagination{}))
	return
}

func (controller *UserController) Delete(c *gin.Context) {
	var inputParam GetUserDetailParam
	err := c.ShouldBindUri(&inputParam)
	helper.PanicIfError(err)

	user := controller.service.FindById(inputParam.Id)
	controller.service.Delete(user)
	c.JSON(http.StatusOK, helper.APIResponse("Delete User Successfully", http.StatusOK, "success", nil, helper.Pagination{}))
	return
}

func mapUserUpdate(request UpdateUserRequest, user User) User {

	if request.Name != "" {
		user.Name = request.Name
	}
	if request.Phone != "" {
		user.Phone = request.Phone
	}
	if request.Address != "" {
		user.Address = request.Address
	}
	if request.City != "" {
		user.City = request.City
	}
	if request.Photo != "" {
		user.Phone = request.Phone
	}

	return user
}
