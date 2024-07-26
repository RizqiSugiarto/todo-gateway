package repository

import (
	"context"

	"github.com/digisata/todo-gateway/domain"
	"github.com/digisata/todo-gateway/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct {
	db         mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) *UserRepository {
	return &UserRepository{
		db:         db,
		collection: collection,
	}
}

func (r UserRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	collection := r.db.Collection(r.collection)

	var user domain.User

	err := collection.FindOne(ctx, bson.M{
		"email":      email,
		"deleted_at": 0,
	}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}
