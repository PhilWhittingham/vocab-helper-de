package noun

import (
	"context"

	"github.com/PhilWhittingham/vocab-helper-de/models"
)

type Service interface {
	GetNouns(ctx context.Context) ([]*models.Noun, error)
	CreateNoun(ctx context.Context, article string, word string, translation string) (*models.Noun, error)
}
