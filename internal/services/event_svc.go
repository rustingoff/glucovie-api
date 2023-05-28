package services

import (
	"context"
	"glucovie/internal/models"
	"glucovie/internal/repositories"
	"time"
)

type EventServiceImpl interface {
	SaveEvent(model *models.EventModel) error
	GetEvents(userID string) ([]*models.EventResponse, error)
	DeleteEvent(id string) error
}

type eventService struct {
	repo repositories.EventRepositoryImpl
}

func NewEventService(repo repositories.EventRepositoryImpl) EventServiceImpl {
	return &eventService{repo: repo}
}

func (s eventService) SaveEvent(model *models.EventModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return s.repo.SaveEvent(ctx, model)
}

func (s eventService) GetEvents(userID string) ([]*models.EventResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return s.repo.GetEvents(ctx, userID)
}

func (s eventService) DeleteEvent(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return s.repo.DeleteEvent(ctx, id)
}
