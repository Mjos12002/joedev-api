package main

import (
	"github.com/gin-gonic/gin"
	"smartdeals/taxi"
	"smartdeals/user"
)

func main() {

	router := gin.Default()
	router.GET("/user", user.GetUser)
	router.GET("/taxi", taxi.GetTaxi)
	router.Run("localhost:8080")

}
