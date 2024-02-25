package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/PhilWhittingham/vocab-helper-de/models"
)

type Noun struct {
	Article     string `bson:"article"`
	Word        string `bson:"word"`
	Translation string `bson:"translation"`
}

type NounRepository struct {
	db *mongo.Collection
}

func NewNounRepository(db *mongo.Database, collection string) *NounRepository {
	return &NounRepository{
		db: db.Collection(collection),
	}
}

func (r NounRepository) GetNouns(ctx context.Context) ([]*models.Noun, error) {
	cur, err := r.db.Find(ctx, bson.M{})
	defer cur.Close(ctx)

	if err != nil {
		return nil, err
	}

	out := make([]*Noun, 0)

	for cur.Next(ctx) {
		noun := new(Noun)
		err := cur.Decode(noun)
		if err != nil {
			return nil, err
		}

		out = append(out, noun)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return toNouns(out), nil
}

func (r NounRepository) CreateNoun(ctx context.Context, noun *models.Noun) error {
	model := toModel(noun)

	_, err := r.db.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	return nil
}

func toModel(n *models.Noun) *Noun {
	return &Noun{
		Article:     n.Article,
		Word:        n.Word,
		Translation: n.Translation,
	}
}

func toNoun(n *Noun) *models.Noun {
	return &models.Noun{
		Article:     n.Article,
		Word:        n.Word,
		Translation: n.Translation,
	}
}

func toNouns(ns []*Noun) []*models.Noun {
	out := make([]*models.Noun, len(ns))

	for i, n := range ns {
		out[i] = toNoun(n)
	}

	return out
}
