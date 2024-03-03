package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/PhilWhittingham/vocab-helper-de/models"
)

type MemoryEvent struct {
	Id               string
	NounId           string
	TimeStamp        time.Time
	ExtentRemembered int64
}

type MemoryEventRepository struct {
	db *mongo.Collection
}

func NewMemoryEventRepository(db *mongo.Database, collection string) *MemoryEventRepository {
	return &MemoryEventRepository{
		db: db.Collection(collection),
	}
}

func (r MemoryEventRepository) GetMemoryEvents(ctx context.Context) ([]*models.MemoryEvent, error) {
	cur, err := r.db.Find(ctx, bson.M{})
	defer cur.Close(ctx)

	if err != nil {
		return nil, err
	}

	out := make([]*MemoryEvent, 0)

	for cur.Next(ctx) {
		noun := new(MemoryEvent)
		err := cur.Decode(noun)
		if err != nil {
			return nil, err
		}

		out = append(out, noun)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return toMemoryEvents(out), nil
}

func toMemoryEvent(e *MemoryEvent) *models.MemoryEvent {
	return &models.MemoryEvent{
		Id:               e.Id,
		NounId:           e.NounId,
		TimeStamp:        e.TimeStamp,
		ExtentRemembered: models.ExtentRemembered(e.ExtentRemembered),
	}
}

func toMemoryEvents(es []*MemoryEvent) []*models.MemoryEvent {
	out := make([]*models.MemoryEvent, len(es))

	for i, e := range es {
		out[i] = toMemoryEvent(e)
	}

	return out
}
