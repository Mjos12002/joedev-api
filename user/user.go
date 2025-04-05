package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	userstruct "smartdeals/struct"
)

func GetUser(c *gin.Context) {
	user := []userstruct.User{
		{ID: "122", Name: "Riad", Age: 23, Role: "admin"},
		{ID: "132", Name: "Adil", Age: 23, Role: "admin"},
		{ID: "142", Name: "Mico", Age: 23, Role: "admin"},
		{ID: "152", Name: "Cyusa", Age: 23, Role: "admin"},
	}
	c.JSON(http.StatusOK, user)
}
