package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PhilWhittingham/vocab-helper-de/noun"
)

type Noun struct {
	Article     string `json:"article" binding:"required"`
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

func (h *Handler) Get(c *gin.Context) {
	nouns, err := h.service.GetNouns(c.Request.Context())

	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.IndentedJSON(http.StatusOK, nouns)
}

func (h *Handler) Create(c *gin.Context) {
	var newNoun Noun

	if err := c.BindJSON(&newNoun); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.service.CreateNoun(c.Request.Context(), newNoun.Article, newNoun.Word, newNoun.Translation)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, newNoun)
}
