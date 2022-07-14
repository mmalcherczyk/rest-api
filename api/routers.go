package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mmalcherczyk/rest-api/core/handlers"
)

func CreateRouter(handler *handlers.UserHandler) *gin.Engine {
	router := gin.Default()
	router.GET("/users", handler.GetUsers())
	router.POST("/users", handler.PostUser())
	router.GET("users/:id", handler.GetUserByID())

	return router

}
