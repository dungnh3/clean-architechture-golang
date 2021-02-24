package api

import (
	"github.com/dungnh3/clean-architechture/app/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterServiceHandler(svc *service.Service) *gin.Engine {
	router := gin.Default()

	userRouterGroup := router.Group("/users")
	{
		userRouterGroup.GET("", svc.GetUsers)
		userRouterGroup.GET("/:id", svc.GetUserByID)
		userRouterGroup.POST("", svc.CreateUser)
		userRouterGroup.PUT("", svc.UpdateUser)
	}
	return router
}
