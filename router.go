package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var rates []rate

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, rates)
}

func postAlbums(c *gin.Context) {
	var newAlbum rate

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	rates = append(rates, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
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

func routerInit() {
	router := gin.Default()
	router.GET("/", getAlbums)
	router.GET("/:name", getRateByName)
	router.POST("/", postAlbums)

	router.Run(":8080")
}
