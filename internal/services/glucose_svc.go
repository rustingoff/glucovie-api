package services

import (
	"context"
	"glucovie/internal/models"
	"glucovie/internal/repositories"
	"time"
)

type GlucoseServiceImpl interface {
	SaveGlucoseLevel(model *models.GlucoseLevel) error
	GetWeekGlucoseLevel(userID string) ([]models.GlucoseResponse, error)
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

func (s glucoseService) GetWeekGlucoseLevel(userID string) ([]models.GlucoseResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := s.repo.GetWeekGlucoseLevel(ctx, userID)
	if err != nil {
		return []models.GlucoseResponse{}, nil
	}

	var response = make([]*models.GlucoseResponse, len(resp))

	var val = make(map[int32]models.GlucoseResponse)
	var count = make(map[int32]uint)

	for k, v := range resp {
		response[k] = &models.GlucoseResponse{
			Level: v.Level,
			Type:  v.Type,
			Day:   int32(v.Date.Weekday()),
		}
		val[int32(v.Date.Weekday())] = models.GlucoseResponse{
			Level: val[int32(v.Date.Weekday())].Level + v.Level,
			Type:  v.Type,
			Day:   int32(v.Date.Weekday()),
		}

		count[response[k].Day]++
	}

	var r = make([]models.GlucoseResponse, 7)
	for k, v := range val {
		v.Level = v.Level / float32(count[v.Day])
		r[k] = v
	}

	for k, v := range r {
		if v == (models.GlucoseResponse{}) {
			r[k] = models.GlucoseResponse{
				Type:  "1",
				Level: 0,
				Day:   int32(k),
			}
		}

	}

	return r, nil
}
