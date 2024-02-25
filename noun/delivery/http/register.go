package http

import (
	"github.com/gin-gonic/gin"

	"github.com/PhilWhittingham/vocab-helper-de/noun"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, s noun.Service) {
	h := NewHandler(s)

	nouns := router.Group("/nouns")
	{
		nouns.GET("", h.Get)
		nouns.POST("", h.Create)
	}
}
