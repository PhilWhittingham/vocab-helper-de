package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PhilWhittingham/vocab-helper-de/models"
)

var nouns = []models.Noun{
	{Article: "Das", Word: "Meeting", Translation: "Meeting"},
}

func getNouns(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nouns)
}

func postNouns(c *gin.Context) {
	var newNoun models.Noun

	if err := c.BindJSON(&newNoun); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	nouns = append(nouns, newNoun)
	c.IndentedJSON(http.StatusCreated, newNoun)
}

func Run(port string) {
	router := gin.Default()
	router.GET("/nouns", getNouns)
	router.POST("/nouns", postNouns)

	router.Run("localhost:" + port)
}
