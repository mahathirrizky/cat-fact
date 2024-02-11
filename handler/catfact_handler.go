package handler

import (
	"cat-fact/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler function for the /cat-fact endpoint
func GetCatFact(c *gin.Context) {
	// Call the external API from the services package
	response, err := services.GetCatFact()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cat fact"})
		return
	}

	c.JSON(http.StatusOK, response.Fact)
}
