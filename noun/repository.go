package noun

import (
	"context"

	"github.com/PhilWhittingham/vocab-helper-de/models"
)

type Repository interface {
	GetNouns(ctx context.Context) ([]*models.Noun, error)
}
