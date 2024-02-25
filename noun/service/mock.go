package service

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/PhilWhittingham/vocab-helper-de/models"
)

type NounServiceMock struct {
	mock.Mock
}

func (m NounServiceMock) GetNouns(ctx context.Context) ([]*models.Noun, error) {
	args := m.Called()

	return args.Get(0).([]*models.Noun), args.Error(1)
}
