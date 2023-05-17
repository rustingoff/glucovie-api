package services

import (
	"context"
	apperrors "glucovie/internal/errors"
	"glucovie/internal/models"
	"glucovie/internal/repositories"
	jwtoken "glucovie/pkg/jwt"
	"time"
)

type UserServiceImpl interface {
	Login(cred *models.UserCredentials) (map[string]interface{}, error)
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

func (s *userService) Login(cred *models.UserCredentials) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	user, err := s.repo.GetUserByEmail(ctx, cred.GetMail())
	if err != nil {
		return nil, err
	}

	isMatch := jwtoken.DoPasswordsMatch(user.Password, cred.GetPassword())
	if !isMatch {
		return nil, apperrors.ErrInvalidUserCredentials
	}

	authToken, err := jwtoken.GenerateToken(*user)
	if err != nil {
		return nil, err
	}

	var tokens = map[string]interface{}{
		"at": authToken.AccessToken,
		"rt": authToken.RefreshToken,
	}

	return tokens, nil
}

func (s *userService) Register(u *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	user, err := s.repo.GetUserByEmail(ctx, u.Email)
	if err != nil {
		return err
	}

	if user.Email != "" {
		return apperrors.ErrUserAlreadyExists
	}

	hashPassword := jwtoken.GenerateHashPassword(u.Password)
	u.Password = hashPassword

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
