package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRole int8

const (
	ADMIN UserRole = iota + 1
	CUSTOMER
	COMMITTEE

	USER_COLLECTION string = "users"
)

type (
	User struct {
		ID        primitive.ObjectID `bson:"_id"`
		Name      string             `bson:"name"`
		Email     string             `bson:"email"`
		Password  string             `bson:"password"`
		Note      *string            `bson:"note"`
		CreatedAt int64              `bson:"created_at"`
		UpdatedAt int64              `bson:"updated_at"`
		DeletedAt int64              `bson:"deleted_at"`
	}

	AuthResponse struct {
		AccessToken string
	}
)
