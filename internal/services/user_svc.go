package services

import (
	"context"
	"glucovie/internal/models"
	"glucovie/internal/repositories"
	"time"
)

type UserServiceImpl interface {
	Login() error
	Register(u *models.User) error
	GetUser(userID string) (*models.User, error)
	DeleteUser(userID string) error
	UpdateUser(userID string, u *models.User) error
}

type userService struct {
	repo repositories.UserRepositoryImpl
}

func NewUserService(repo repositories.UserRepositoryImpl) UserServiceImpl {
	return &userService{repo: repo}
}

// TODO: implement
func (s *userService) Login() error {
	return nil
}

func (s *userService) Register(u *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return s.repo.Register(ctx, u)
}

func (s *userService) GetUser(userID string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return s.repo.GetUser(ctx, userID)
}

// TODO: implement
func (s *userService) DeleteUser(userID string) error {
	panic("TODO: implement")
}

// TODO: implement
func (s *userService) UpdateUser(userID string, u *models.User) error {
	panic("TODO: implement")

}
