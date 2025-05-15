package api

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// Routes
	r.GET("/health", healthCheck)
	r.POST("/hazards", reportHazard)
	r.GET("/route", getSafeRoute)
	r.GET("/alerts", handleAlerts) // WebSocket

	return r
}
