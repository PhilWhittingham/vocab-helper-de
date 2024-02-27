package mongo

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/PhilWhittingham/vocab-helper-de/models"
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

func TestToDatabaseModel(t *testing.T) {
	domainNoun := models.Noun{
		Id:          "some_id",
		Article:     "das",
		Word:        "word",
		Translation: "translation",
	}

	dbNoun := toDatabaseModel(&domainNoun)

	assert.Equal(t, dbNoun.Article, domainNoun.Article)
	assert.Equal(t, dbNoun.Word, domainNoun.Word)
	assert.Equal(t, dbNoun.Translation, domainNoun.Translation)
}
