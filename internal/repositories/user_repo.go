package repositories

import (
	"context"
	"glucovie/internal/constants"
	"glucovie/internal/models"
	"glucovie/pkg/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type UserRepositoryImpl interface {
	Register(ctx context.Context, u *models.User) error
	GetUser(ctx context.Context, userID string) (*models.User, error)
	DeleteUser(ctx context.Context, userID string) error
	UpdateUser(ctx context.Context, userID string, u map[string]interface{}) error
}

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) UserRepositoryImpl {
	return &userRepository{db: db}
}

func (r *userRepository) Register(ctx context.Context, u *models.User) error {
	_, err := r.db.
		Collection(constants.UserCollection).
		InsertOne(ctx, u)

	if err != nil {
		logger.Log.Error("failed to register user", zap.Error(err))
		return err
	}

	return nil
}

func (r *userRepository) GetUser(ctx context.Context, userID string) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return &models.User{}, err
	}

	var user *models.User
	err = r.db.
		Collection(constants.UserCollection).
		FindOne(ctx, bson.M{"_id": objID}).
		Decode(&user)

	if err != nil {
		return &models.User{}, err
	}

	return user, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, userID string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	_, err = r.db.Collection(constants.UserCollection).DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) UpdateUser(ctx context.Context, userID string, u map[string]interface{}) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	_, err = r.db.
		Collection(constants.UserCollection).
		UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": u})
	if err != nil {
		return err
	}

	return nil
}
