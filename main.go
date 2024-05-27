package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

const URL string = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

func main() {
	currencyRates := ExtractQuotes(URL)
	fmt.Println(reflect.TypeOf(currencyRates.Currencies))
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, currencyRates)
	})

	router.POST("/add", func(c *gin.Context) {
		var newCurrency Currency

		if err := c.ShouldBind(&newCurrency); err != nil {
			return
		}
		currencyRates.Currencies = append(currencyRates.Currencies, newCurrency)
	})
	_ = router.Run("localhost:8080")
}
