package exception

import (
	"api-foodmarket/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ErrorHandler(c *gin.Context, recovered interface{}) {

	if err, ok := recovered.(validator.ValidationErrors); ok {
		if ok {
			ValidationErrors(c, err)
			return
		}
	}
	if err, ok := recovered.(string); ok {
		errMessage := map[string]string{
			"message": err,
		}
		c.JSON(http.StatusInternalServerError, errMessage)
		return
	}

	if error, ok := recovered.(error); ok {
		errMessage := map[string]string{
			"message": error.Error(),
		}
		c.JSON(http.StatusInternalServerError, errMessage)
		return
	}
	c.AbortWithStatus(http.StatusInternalServerError)
	return
}

func ValidationErrors(c *gin.Context, err error) {
	response := helper.APIResponse("BAD REQUEST", http.StatusBadRequest, "failed", err.Error(), helper.Pagination{})
	c.JSON(http.StatusBadRequest, response)
	return
}
