package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type cartItem struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
	ImagePath   string  `json:"imagePath"`
}

var carItemsList = []cartItem{
	{ID: "1", Title: "iPhone 11 64 GB", Description: "Some test description", Price: 499.99, Currency: "EUR", ImagePath: "image/test.png"},
	{ID: "2", Title: "Mac Mini", Description: "Some test description #2", Price: 756, Currency: "USD", ImagePath: "image/test1.png"},
	{ID: "3", Title: "MacBook Air M1", Description: "Some test description #3", Price: 222.00, Currency: "RUB", ImagePath: "image/test2.png"},
}

func getItemsList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, carItemsList)
}

func getCartItemByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range carItemsList {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cart item not found"})
}

func addItemsToCart(c *gin.Context) {
	var newCartItem cartItem

	if err := c.BindJSON(&newCartItem); err != nil {
		return
	}

	carItemsList = append(carItemsList, newCartItem)
	c.IndentedJSON(http.StatusCreated, newCartItem)
}
