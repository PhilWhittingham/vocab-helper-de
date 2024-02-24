package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/PhilWhittingham/vocab-helper-de/models"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetNouns(t *testing.T) {
	r := SetUpRouter()
	r.GET("/nouns", getNouns)

	req, _ := http.NewRequest("GET", "/nouns", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostNouns(t *testing.T) {

	t.Run("With correctly formatted input returns success code", func(t *testing.T) {
		r := SetUpRouter()
		r.POST("/nouns", postNouns)

		newNoun := models.Noun{
			Article:     "Das",
			Word:        "word",
			Translation: "translation",
		}
		jsonValue, _ := json.Marshal(newNoun)
		req, _ := http.NewRequest("POST", "/nouns", bytes.NewBuffer(jsonValue))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("With incorrectly formatted input returns failure code", func(t *testing.T) {
		r := SetUpRouter()
		r.POST("/nouns", postNouns)

		badBody := []byte(`{
			"any_key": "any_value"
		}`)

		req, _ := http.NewRequest("POST", "/nouns", bytes.NewBuffer(badBody))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
