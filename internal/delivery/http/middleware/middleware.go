package middleware

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func LoadUpperLower() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.Request.URL.Path

		log.Printf("[UPPER] %s %s", strings.ToUpper(method), strings.ToUpper(path))
		log.Printf("[LOWER] %s %s", strings.ToLower(method), strings.ToLower(path))
		c.Next()
	}
}
