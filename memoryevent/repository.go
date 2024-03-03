package memoryevent

import (
	"context"

	"github.com/PhilWhittingham/vocab-helper-de/models"
)

type Repository interface {
	GetMemoryEvents(ctx context.Context) ([]*models.MemoryEvent, error)
	CreateMemoryEvent(ctx context.Context, noun *models.MemoryEvent) (*models.MemoryEvent, error)
}
