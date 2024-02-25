package mongo

import (
	"testing"

	"github.com/PhilWhittingham/vocab-helper-de/models"
	"github.com/stretchr/testify/assert"
)

func TestToNoun(t *testing.T) {
	dbNoun := Noun{
		Article:     "das",
		Word:        "word",
		Translation: "translation",
	}

	domainNoun := toNoun(&dbNoun)

	assert.Equal(t, domainNoun.Article, dbNoun.Article)
	assert.Equal(t, domainNoun.Word, dbNoun.Word)
	assert.Equal(t, domainNoun.Translation, dbNoun.Translation)
}

func TestToNouns(t *testing.T) {
	dbNoun := Noun{
		Article:     "das",
		Word:        "word",
		Translation: "translation",
	}

	dbNouns := []*Noun{&dbNoun}

	domainNouns := toNouns(dbNouns)

	assert.Equal(t, len(domainNouns), 1)
}

func TestToModel(t *testing.T) {
	domainNoun := models.Noun{
		Article:     "das",
		Word:        "word",
		Translation: "translation",
	}

	dbNoun := toModel(&domainNoun)

	assert.Equal(t, dbNoun.Article, domainNoun.Article)
	assert.Equal(t, dbNoun.Word, domainNoun.Word)
	assert.Equal(t, dbNoun.Translation, domainNoun.Translation)
}
