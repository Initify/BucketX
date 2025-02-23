package routes

import (
	"bucketX/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/", controllers.WelcomeController)
		// Files
		api.GET("/file/:file_key", controllers.FetchFileController)
		api.POST("/file", controllers.UploadFileController)

		// Buckets
		api.GET("/buckets", controllers.ListBucketsController)
	}
}
