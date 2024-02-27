package service

import (
	"context"

	"github.com/PhilWhittingham/vocab-helper-de/models"
	"github.com/PhilWhittingham/vocab-helper-de/noun"
)

type NounService struct {
	nounRepo noun.Repository
}

func NewNounService(nounRepo noun.Repository) *NounService {
	return &NounService{
		nounRepo: nounRepo,
	}
}

func (n NounService) GetNouns(ctx context.Context) ([]*models.Noun, error) {
	return n.nounRepo.GetNouns(ctx)
}

func (n NounService) CreateNoun(ctx context.Context, article string, word string, translation string) (*models.Noun, error) {
	noun := &models.Noun{
		Article:     article,
		Word:        word,
		Translation: translation,
	}

	return n.nounRepo.CreateNoun(ctx, noun)
}
