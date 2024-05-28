package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var rates []rate

func getRates(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, rates)
}

func getRateByName(c *gin.Context) {
	name := c.Param("name")
	for _, rate := range rates {
		if rate.Name == name {
			c.IndentedJSON(http.StatusOK, rate)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "rate not found"})
}

func deleteRateByName(c *gin.Context) {
	name := c.Param("name")
	for index, rate := range rates {
		if rate.Name == name {
			rates = append(rates[:index], rates[index+1:]...)
			c.IndentedJSON(http.StatusOK, rate)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "rate not found"})
}

func editRateByName(c *gin.Context) {
	name := c.Param("name")
	for index, rate := range rates {
		if rate.Name == name {
			rates = append(rates[:index], rates[index+1:]...)
			rates = append(rates, rate)
			c.IndentedJSON(http.StatusOK, rate)
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "rate not found"})
}

func routerInit() {
	router := gin.Default()
	router.GET("/rates", getRates)
	router.GET("/rates/get/:name", getRateByName)
	router.POST("/rates/delete/:name", deleteRateByName)
	router.POST("/rates/edit/:name", editRateByName)

	router.Run(":8080")
}
