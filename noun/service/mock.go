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

func (m NounServiceMock) CreateNoun(ctx context.Context, article string, word string, translation string) (*models.Noun, error) {
	args := m.Called(article, word, translation)

	return args.Get(0).(*models.Noun), args.Error(1)
}
