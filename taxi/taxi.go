package taxi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_struct "smartdeals/struct"
)

func GetTaxi(context *gin.Context) {
	taxi := []_struct.Taxi{
		{ID: "1234", Name: "Mercedes Benz", Type: "VIP"},
		{ID: "3423", Name: "Toyota Corolla", Type: "Regular Cab"},
		{ID: "4783", Name: "Mazda C1", Type: "Ordinary"},
	}

	context.JSON(http.StatusOK, taxi)
}
