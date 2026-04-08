package handler

import "github.com/gin-gonic/gin"

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/green/instance/:idInstance")
	{
		api.GET("/settings/:apiToken", h.GetSettings)
		api.GET("/state/:apiToken", h.GetStateInstance)
		api.POST("/send-message/:apiToken", h.SendMessage)
		api.POST("/send-file/:apiToken", h.SendFileByURL)
	}
}
