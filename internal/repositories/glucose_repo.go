package repositories

import (
	"context"
	"glucovie/internal/constants"
	"glucovie/internal/models"
	"glucovie/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
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

func (repo *glucoseRepository) GetWeekGlucoseLevel(ctx context.Context) ([]*models.GlucoseLevel, error) {
	return nil, nil
}
