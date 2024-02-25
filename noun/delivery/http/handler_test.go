package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/PhilWhittingham/vocab-helper-de/noun/service"
)

func TestGetNouns(t *testing.T) {
	router := gin.Default()

	s := new(service.NounServiceMock)

	group := router.Group("/api")
	RegisterHTTPEndpoints(group, s)

	req, _ := http.NewRequest("GET", "/api/nouns", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostNouns(t *testing.T) {

	t.Run("With correctly formatted input returns success code", func(t *testing.T) {
		router := gin.Default()

		service := new(service.NounServiceMock)

		group := router.Group("/api")
		RegisterHTTPEndpoints(group, service)

		newNoun := Noun{
			Article:     "Das",
			Word:        "word",
			Translation: "translation",
		}
		jsonValue, _ := json.Marshal(newNoun)
		req, _ := http.NewRequest("POST", "/api/nouns", bytes.NewBuffer(jsonValue))

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("With incorrectly formatted input returns failure code", func(t *testing.T) {
		router := gin.Default()

		service := new(service.NounServiceMock)

		group := router.Group("/api")
		RegisterHTTPEndpoints(group, service)

		badBody := []byte(`{
			"any_key": "any_value"
		}`)

		req, _ := http.NewRequest("POST", "/api/nouns", bytes.NewBuffer(badBody))

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}