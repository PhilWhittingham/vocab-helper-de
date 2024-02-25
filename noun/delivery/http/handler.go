package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PhilWhittingham/vocab-helper-de/noun"
)

type Noun struct {
	Article     string `json:"artical" binding:"required"`
	Word        string `json:"word" binding:"required"`
	Translation string `json:"translation" binding:"required"`
}

type Handler struct {
	service noun.Service
}

func NewHandler(service noun.Service) *Handler {
	return &Handler{
		service: service,
	}
}

var nouns = []Noun{
	{Article: "Das", Word: "Meeting", Translation: "Meeting"},
}

func (h *Handler) Get(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nouns)
}

func (h *Handler) Create(c *gin.Context) {
	var newNoun Noun

	if err := c.BindJSON(&newNoun); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	nouns = append(nouns, newNoun)
	c.IndentedJSON(http.StatusCreated, newNoun)
}
