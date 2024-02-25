package service

import (
	"context"

	"github.com/PhilWhittingham/vocab-helper-de/models"
	"github.com/PhilWhittingham/vocab-helper-de/noun"
)

type NounService struct {
	nounRepo noun.Repository
}

func NewNounservice(nounRepo noun.Repository) *NounService {
	return &NounService{
		nounRepo: nounRepo,
	}
}

func (n NounService) GetNouns(ctx context.Context) ([]*models.Noun, error) {
	return n.nounRepo.GetNouns(ctx)
}
