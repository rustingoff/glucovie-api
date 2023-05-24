package repositories

import (
	"context"
	"glucovie/internal/constants"
	"glucovie/internal/models"
	"glucovie/pkg/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type GlucoseRepositoryImpl interface {
	SaveGlucoseLevel(ctx context.Context, u *models.GlucoseLevel) error
	GetWeekGlucoseLevel(ctx context.Context) ([]*models.GlucoseLevel, error)
}

type glucoseRepository struct {
	db *mongo.Database
}

func NewGlucoseRepository(db *mongo.Database) GlucoseRepositoryImpl {
	return &glucoseRepository{db: db}
}

func (r *glucoseRepository) SaveGlucoseLevel(ctx context.Context, model *models.GlucoseLevel) error {
	_, err := r.db.
		Collection(constants.GlucoseCollection).
		InsertOne(ctx, model)

	if err != nil {
		logger.Log.Error("failed to insert glucose level", zap.Error(err))
		return err
	}

	return nil
}

func (r *glucoseRepository) GetWeekGlucoseLevel(ctx context.Context) ([]*models.GlucoseLevel, error) {
	var response = make([]*models.GlucoseLevel, 7)

	opts := &options.FindOptions{
		Sort:  bson.M{"date": -1},
		Skip:  &[]int64{0}[0],
		Limit: &[]int64{7}[0],
	}

	cursor, err := r.db.
		Collection(constants.GlucoseCollection).
		Find(ctx, bson.M{}, opts)

	if err != nil {
		logger.Log.Error("failed to find glucose level", zap.Error(err))
		return response, err
	}

	if err := cursor.All(ctx, &response); err != nil {
		logger.Log.Error("failed to decode glucose level", zap.Error(err))
		return response, err
	}

	return response, nil
}
