package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(InjectAppName())
	r.Use(ErrorHandler())

	r.GET("/ping", Pong)
	r.GET("/error", ErrorEndpoint)

	r.Run("0.0.0.0:8888")
}

func InjectAppName() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-App-Name", "testapp")
		c.Set("X-App-Name", "testapp")
		c.Next()
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func (_c *gin.Context) {
		defer func(c *gin.Context) {
			if rec := recover(); rec != nil {
				c.JSON(500, gin.H{
					"error": rec,
				})
			}
		}(_c)
		_c.Next()
	}
}

func Pong(c *gin.Context) {
	fmt.Println(c.Get("X-App-Name"))
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func ErrorEndpoint(c *gin.Context) {
	panic("die")
}
