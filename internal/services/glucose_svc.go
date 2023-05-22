package services

import (
	"context"
	"glucovie/internal/models"
	"glucovie/internal/repositories"
	"time"
)

type GlucoseServiceImpl interface {
	SaveGlucoseLevel(model *models.GlucoseLevel) error
	GetWeekGlucoseLevel() ([]*models.GlucoseLevel, error)
}

type glucoseService struct {
	repo repositories.GlucoseRepositoryImpl
}

func NewGlucoseService(repo repositories.GlucoseRepositoryImpl) GlucoseServiceImpl {
	return &glucoseService{repo: repo}
}

func (s glucoseService) SaveGlucoseLevel(model *models.GlucoseLevel) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	model.Date = time.Now()

	return s.repo.SaveGlucoseLevel(ctx, model)
}

func (s glucoseService) GetWeekGlucoseLevel() ([]*models.GlucoseLevel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return s.repo.GetWeekGlucoseLevel(ctx)
}
