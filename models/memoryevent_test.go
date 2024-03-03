package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert.Equal(t, "none", None.String())
	assert.Equal(t, "all", All.String())
	assert.Equal(t, "only-article", OnlyArticle.String())
	assert.Equal(t, "only-word", OnlyWord.String())
}
