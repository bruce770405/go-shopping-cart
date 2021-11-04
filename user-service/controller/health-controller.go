package controller

import (
	"github.com/gin-gonic/gin"
)

// User manages
type Health struct {
}

func (h *Health) HealthGET(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "UP",
	})
}

func (h *Health) ReadyGET(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "UP",
	})
}
