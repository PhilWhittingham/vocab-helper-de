package noun

import (
	"context"

	"github.com/PhilWhittingham/vocab-helper-de/models"
)

type UseCase interface {
	GetNouns(ctx context.Context) ([]*models.Noun, error)
}
