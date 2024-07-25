package repository

import (
	"context"
	"time"

	"github.com/digisata/auth-service/domain"
	"github.com/digisata/auth-service/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (r UserRepository) Create(ctx context.Context, req domain.User) error {
	collection := r.db.Collection(r.collection)
	user := req

	now := time.Now().Local().Unix()
	user.CreatedAt = now
	user.UpdatedAt = now
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (r UserRepository) GetAll(ctx context.Context, req domain.GetAllUserRequest) ([]domain.User, error) {
	collection := r.db.Collection(r.collection)
	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})

	// Base filter
	filter := bson.M{
		"deleted_at": 0,
		"role":       bson.M{"$ne": domain.ADMIN},
	}

	if req.Search != nil && *req.Search != "" {
		pattern := ".*" + *req.Search + ".*"
		filter["$or"] = []interface{}{
			bson.M{"name": bson.M{"$regex": pattern, "$options": "i"}},
			bson.M{"email": bson.M{"$regex": pattern, "$options": "i"}},
		}
	}

	if req.IsActive != nil {
		filter["is_active"] = *req.IsActive
	}

	// Adding role condition if req.Role is not nil
	if req.Role != nil {
		filter = bson.M{
			"$and": []interface{}{
				filter,
				bson.M{"role": *req.Role},
			},
		}
	}

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var users []domain.User
	err = cursor.All(ctx, &users)
	if users == nil {
		return users, err
	}

	return users, nil
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

func (r UserRepository) GetByID(ctx context.Context, id string) (domain.User, error) {
	collection := r.db.Collection(r.collection)

	var user domain.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = collection.FindOne(ctx, bson.M{
		"_id":        idHex,
		"deleted_at": 0,
	}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r UserRepository) Update(ctx context.Context, req domain.UpdateUser) error {
	collection := r.db.Collection(r.collection)

	// Create a map to hold the fields to be updated
	updateFields := bson.M{}

	// Check each field in req and add it to the map if it is non-nil
	if req.CustomerID != nil {
		updateFields["customer_id"] = req.CustomerID
	}

	if req.Name != nil {
		updateFields["name"] = req.Name
	}

	if req.IsActive != nil {
		updateFields["is_active"] = req.IsActive
	}

	if req.Note != nil {
		updateFields["note"] = req.Note
	}

	// Set the UpdatedAt field
	updateFields["updated_at"] = time.Now().Local().Unix()

	// Perform the update
	_, err := collection.UpdateOne(ctx, bson.M{"_id": req.ID}, bson.M{"$set": updateFields})
	if err != nil {
		return err
	}

	return nil
}

func (r UserRepository) Delete(ctx context.Context, req domain.DeleteUser) error {
	collection := r.db.Collection(r.collection)

	deleteUser := req
	now := time.Now().Local().Unix()
	deleteUser.IsActive = false
	deleteUser.UpdatedAt = now
	deleteUser.DeletedAt = now

	_, err := collection.UpdateOne(ctx, bson.M{"_id": req.ID}, bson.M{"$set": deleteUser})
	if err != nil {
		return err
	}

	return nil
}
