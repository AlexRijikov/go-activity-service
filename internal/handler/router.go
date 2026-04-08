package handler

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine, h *Handler) {
	r.POST("/events", h.CreateEvent)
	r.GET("/events", h.GetEvents)
	r.GET("/stats", h.GetStats)
}
