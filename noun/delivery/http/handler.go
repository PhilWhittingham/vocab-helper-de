package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PhilWhittingham/vocab-helper-de/models"
	"github.com/PhilWhittingham/vocab-helper-de/noun"
)

type Noun struct {
	Id          string `json:"id"`
	Article     string `json:"article"`
	Word        string `json:"word"`
	Translation string `json:"translation"`
}

type NounCreate struct {
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

	c.IndentedJSON(http.StatusOK, toNouns(nouns))
}

func (h *Handler) Create(c *gin.Context) {
	var newNoun NounCreate

	if err := c.BindJSON(&newNoun); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	n, err := h.service.CreateNoun(c.Request.Context(), newNoun.Article, newNoun.Word, newNoun.Translation)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, toNoun(n))
}

func toNoun(n *models.Noun) *Noun {
	return &Noun{
		Id:          n.Id,
		Article:     n.Article,
		Word:        n.Word,
		Translation: n.Translation,
	}
}

func toNouns(ns []*models.Noun) []*Noun {
	out := make([]*Noun, len(ns))

	for i, n := range ns {
		out[i] = toNoun(n)
	}

	return out
}
