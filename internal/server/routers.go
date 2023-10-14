package server

import (
	"notifier/internal/telegram"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the available routes for the server
func RegisterRoutes(r *gin.Engine, telegramClient *telegram.Client) {
	// Define a group for API endpoints
	api := r.Group("/api")
	{
		// POST endpoint for sending a message via Telegram
		api.POST("/send_message", func(c *gin.Context) {
			// Extracting text from the POST request
			var json struct {
				Text string `json:"text" binding:"required"`
			}
			if err := c.ShouldBindJSON(&json); err != nil {
				c.JSON(400, gin.H{"error": "Missing field"})
				return
			}

			// Send message using Telegram client
			err := telegramClient.SendMessage(json.Text)
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to send message"})
				return
			}

			c.JSON(200, gin.H{"status": "Message sent"})
		})
	}
}
