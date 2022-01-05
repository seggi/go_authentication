package utils

import "github.com/gin-gonic/gin"

// Response struct
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Json error response function
// ErrorJON
func ErrorJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"error": data})
}

// SuccessJSON
func SuccessJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"msg": data})
}
