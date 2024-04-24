package mongodb

import (
	"auth-service/internal/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{
		Collection: collection,
	}
}

func (ur *UserRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	objID, errId := primitive.ObjectIDFromHex(id)
	if errId != nil {
		return nil, errId
	}
	if ur.Collection == nil {
		return nil, errors.New("collection is nil")
	}
	var user models.User

	filter := bson.M{"_id": objID}
	err := ur.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) AddUser(ctx context.Context, username, hashedPassword string) (*models.User, error) {
	now := time.Now().Format(time.RFC3339)
	user := models.User{
		Username:       username,
		HashedPassword: hashedPassword,
		CreatedAt:      now,
		LastLogin:      now,
	}

	result, err := ur.Collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to convert InsertedID to ObjectID")
	}
	user.ID = oid.Hex()

	return &user, nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, id, username, email string) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update := bson.M{"$set": bson.M{"username": username, "email": email}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedUser models.User
	err = ur.Collection.FindOneAndUpdate(ctx, bson.M{"_id": objID}, update, opts).Decode(&updatedUser)
	if err != nil {
		return nil, err
	}
	return &updatedUser, nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = ur.Collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func (ur *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	filter := bson.M{"username": username}
	err := ur.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
