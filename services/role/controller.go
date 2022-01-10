package role

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoleController struct {
	service Service
}

func NewController(service Service) *RoleController {
	return &RoleController{service: service}
}

func (controller *RoleController) Route(app *gin.Engine) {
	route := app.Group("api/roles")
	route.GET("/", controller.List)
}

func (controller *RoleController) List(c *gin.Context) {
	roles := controller.service.FindAll()

	c.JSON(http.StatusOK, roles)
	return
}
