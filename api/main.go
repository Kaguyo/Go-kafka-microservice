package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.PATCH("/api/v1", func(c *gin.Context) {
		var body map[string]interface{}

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "invalid JSON"})
			return
		}
		c.JSON(200, gin.H{
			"message": "PATCH request received",
			"data":    body})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
