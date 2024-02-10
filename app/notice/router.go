package notice

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/notice/controller"
	"github.com/pfjhyyj/ether/app/notice/service"
	"github.com/pfjhyyj/ether/middleware"
)

func SetMessageRouter(r *gin.RouterGroup) {
	messageService := service.NewMessageService()
	messageController := controller.NewMessageController(messageService)

	router := r.Group("/messages")
	router.Use(middleware.AuthMiddleware())
	{
		router.GET("", messageController.ListMyMessages)
		router.PUT("/:messageId/read", messageController.ReadMessage)
		router.PUT("/batchRead", messageController.BatchReadMessage)
	}
}

func SetRouter(r *gin.RouterGroup) {
	SetMessageRouter(r)
}
