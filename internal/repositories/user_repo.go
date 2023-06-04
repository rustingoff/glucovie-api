package repositories

import (
	"context"
	"errors"
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
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	SaveUserSettings(ctx context.Context, model models.SettingModel, userID string) error
	GetSettingsByUserId(ctx context.Context, userID string) (*models.SettingModel, error)
	GetAllUsersIDs(ctx context.Context) ([]models.User, error)
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

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user = &models.User{}
	err := r.db.
		Collection(constants.UserCollection).
		FindOne(ctx, bson.M{"email": email}).
		Decode(&user)

	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return &models.User{}, err
	}

	return user, nil
}

func (r *userRepository) SaveUserSettings(ctx context.Context, model models.SettingModel, userID string) error {
	objID, _ := primitive.ObjectIDFromHex(userID)

	_, err := r.db.Collection(constants.UserCollection).UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": model})
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetSettingsByUserId(ctx context.Context, userID string) (*models.SettingModel, error) {
	objID, _ := primitive.ObjectIDFromHex(userID)

	var settings = &models.SettingModel{}
	err := r.db.
		Collection(constants.UserCollection).
		FindOne(ctx, bson.M{"_id": objID}).
		Decode(&settings)

	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return &models.SettingModel{}, err
	}

	return settings, nil
}

func (r *userRepository) GetAllUsersIDs(ctx context.Context) ([]models.User, error) {
	var users []models.User
	cursor, err := r.db.Collection(constants.UserCollection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
