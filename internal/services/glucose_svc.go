package services

import (
	"context"
	"fmt"
	"glucovie/internal/models"
	"glucovie/internal/repositories"
	"time"
)

type GlucoseServiceImpl interface {
	SaveGlucoseLevel(model *models.GlucoseLevel) error
	GetWeekGlucoseLevel() ([]*models.GlucoseResponse, error)
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

func (s glucoseService) GetWeekGlucoseLevel() ([]*models.GlucoseResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var response = make([]*models.GlucoseResponse, 7)

	resp, err := s.repo.GetWeekGlucoseLevel(ctx)
	if err != nil {
		return response, nil
	}

	if len(resp) < 7 {
		for i := len(resp); i < 7; i++ {
			resp = append(resp, &models.GlucoseLevel{
				Type:  "1",
				Level: "0",
				Date:  resp[0].Date.AddDate(0, 0, -i),
			})
		}
	}

	for k, v := range resp {
		response[len(resp)-1-k] = &models.GlucoseResponse{
			Level: v.Level,
			Type:  v.Type,
			Day:   fmt.Sprint(int(v.Date.Weekday())),
		}
	}

	return response, nil
}
