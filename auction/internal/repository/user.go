package repository

import (
	"context"
	"errors"
	"time"

	"github.com/israelmiranda/go-expert/auction/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongo struct {
	ID        string `bson:"_id"`
	Name      string `bson:"name"`
	Timestamp int64  `bson:"timestamp"`
}

func fromUser(user domain.User) UserMongo {
	return UserMongo{
		ID:        user.ID,
		Name:      user.Name,
		Timestamp: user.Timestamp.Unix(),
	}
}

func (u UserMongo) toUser() domain.User {
	return domain.User{
		ID:        u.ID,
		Name:      u.Name,
		Timestamp: time.Unix(u.Timestamp, 0),
	}
}

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) UserRepository {
	return UserRepository{
		collection: database.Collection("users"),
	}
}

func (r UserRepository) FindAll(ctx context.Context) ([]domain.User, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return []domain.User{}, err
	}
	defer cursor.Close(ctx)

	var usersMongo []UserMongo
	if err := cursor.All(ctx, &usersMongo); err != nil {
		return []domain.User{}, err
	}

	var users []domain.User
	for _, userMongo := range usersMongo {
		users = append(users, userMongo.toUser())
	}

	return users, nil
}

func (r UserRepository) FindById(ctx context.Context, id string) (domain.User, error) {
	filter := bson.M{"_id": id}

	var userMongo UserMongo
	err := r.collection.FindOne(ctx, filter).Decode(&userMongo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.User{}, err
		}

		return domain.User{}, err
	}

	return userMongo.toUser(), nil
}

func (r UserRepository) Create(ctx context.Context, user domain.User) error {
	_, err := r.collection.InsertOne(ctx, fromUser(user))
	if err != nil {
		return err
	}

	return nil
}
