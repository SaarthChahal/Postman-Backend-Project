package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type headphone struct {
	ID    string  `json:"id"`
	Item  string  `json:"item"`
	Brand string  `json:"brand"`
	Price float64 `json:"price"`
}

func getHeadphones(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, headphones)
}

var headphones = []headphone{
	{ID: "1", Item: "Wireless Headphone", Brand: " Sennheiser", Price: 8990},
	{ID: "2", Item: "Foldable Headphone", Brand: "Zebronics", Price: 783},
	{ID: "3", Item: " Wired Headphone", Brand: "JBL", Price: 599},
	{ID: "4", Item: "Bluetooth Headphone", Brand: "boAT", Price: 999},
	{ID: "5", Item: "Noise Cancellation Headphones", Brand: "Sony", Price: 9990},
}

func main() {
	router := gin.Default()
	router.GET("/headphones", getHeadphones)
	router.GET("/headphones/:id", getHeadphoneByID)
	router.POST("/headphones", postHeadphones)

	router.Run("localhost:8080")
}

func postHeadphones(c *gin.Context) {
	var newHeadphone headphone

	if err := c.BindJSON(&newHeadphone); err != nil {
		return
	}

	headphones = append(headphones, newHeadphone)
	c.IndentedJSON(http.StatusCreated, newHeadphone)
}

func getHeadphoneByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range headphones {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "headphone not found"})
}
