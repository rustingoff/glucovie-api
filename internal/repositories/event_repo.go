package repositories

import (
	"context"
	"glucovie/internal/constants"
	"glucovie/internal/models"
	"glucovie/pkg/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type EventRepositoryImpl interface {
	SaveEvent(ctx context.Context, u *models.EventModel) error
	GetEvents(ctx context.Context, userID string) ([]*models.EventResponse, error)
}

type eventRepository struct {
	db *mongo.Database
}

func NewEventRepository(db *mongo.Database) EventRepositoryImpl {
	return &eventRepository{db: db}
}

func (r *eventRepository) SaveEvent(ctx context.Context, model *models.EventModel) error {
	_, err := r.db.
		Collection(constants.EventCollection).
		InsertOne(ctx, model)

	if err != nil {
		logger.Log.Error("failed to insert glucose level", zap.Error(err))
		return err
	}

	return nil
}

func (r *eventRepository) GetEvents(ctx context.Context, userID string) ([]*models.EventResponse, error) {
	var response = []*models.EventResponse{}

	cursor, err := r.db.
		Collection(constants.EventCollection).
		Find(ctx, bson.M{"user_id": userID})

	if err != nil {
		logger.Log.Error("failed to find events", zap.Error(err))
		return response, err
	}

	if err := cursor.All(ctx, &response); err != nil {
		logger.Log.Error("failed to decode events", zap.Error(err))
		return response, err
	}

	return response, nil
}
