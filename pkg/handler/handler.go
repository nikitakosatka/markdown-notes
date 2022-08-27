package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nikitakosatka/markdown-notes/pkg/service"
)

func IncludeRouter(r *gin.Engine) {
	apiGroup := r.Group("/api")
	{
		apiGroup.POST("/", service.Create)
		apiGroup.GET("/", service.GetAll)
		apiGroup.GET("/:id", service.Read)
		apiGroup.PUT("/:id", service.Update)
		apiGroup.DELETE("/:id", service.Remove)
	}
}
