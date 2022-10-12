package main

import (
	"fmt"
	"github.com/FOS-Lover/Golang-Study-Notes/service"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello world")
	service.TestUserService()
	service.TestCustomerService()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
}
