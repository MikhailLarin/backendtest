package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/cart/items", getItemsList)
	router.GET("/cart/items/:id", getCartItemByID)
	router.POST("/cart/items", addItemsToCart)
	router.Run("localhost:8080")
}
