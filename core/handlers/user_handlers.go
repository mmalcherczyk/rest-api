package handlers

import (
	"github.com/mmalcherczyk/rest-api/core/models"
	"github.com/mmalcherczyk/rest-api/core/processors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	processor *processors.UserProcessor
}

func NewUserHandler(processor *processors.UserProcessor) *UserHandler {
	return &UserHandler{processor: processor}
}

func (handler *UserHandler) GetUsers(c *gin.Context) {
	users, err := handler.processor.ListUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest,
			map[string]string{
				"result": "error",
				"error":  err.Error(),
			})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (handler *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	usr, err := handler.processor.FindUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			map[string]string{
				"result": "error",
				"error":  err.Error(),
			})
		return
	}
	c.JSON(http.StatusOK, usr)

}

func (handler *UserHandler) PostUser(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest,
			map[string]string{
				"result": "error",
				"error":  err.Error(),
			})
		return
	}

	if err := handler.processor.CreateUser(newUser); err != nil {
		c.JSON(http.StatusBadRequest,
			map[string]string{
				"result": "error",
				"error":  err.Error(),
			})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}
