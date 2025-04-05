package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "User are coming",
		})
	})
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
