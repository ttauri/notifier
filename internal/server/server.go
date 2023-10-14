package server

import (
	"github.com/gin-gonic/gin"
	"notifier/internal/telegram"
)

// NewServer creates and configures a new Gin server
func NewServer(telegramClient *telegram.Client) *gin.Engine {
	r := gin.Default()
	RegisterRoutes(r, telegramClient)
	return r
}
