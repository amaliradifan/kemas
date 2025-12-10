package main

import (
	"github.com/gin-gonic/gin"

	"kemas/models"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/products", listProducts)
	router.POST("/products/transfer-stock", transferStock)

	router.Run(":3000")
}

func listProducts(c *gin.Context) {
	search := c.Query("search")
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")

	products, err := models.ListProductHandler(search, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func transferStock(c *gin.Context) {
	var input struct {
		SourceID      string `json:"source_id"`
		DestinationID string `json:"destination_id"`
		Quantity      int    `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

  err := models.TransferStockHandler(input.SourceID, input.DestinationID, input.Quantity)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
}
