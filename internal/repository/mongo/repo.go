package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Adelioz/split/internal/models"
	"github.com/Adelioz/split/internal/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type repository struct {
	c *mongo.Client
}

// AddUser implements service.Repository.
func (r *repository) AddUser(ctx context.Context, user models.User) error {
	collection := r.usersCollection()
	_, err := collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

// GetUser implements service.Repository.
func (r *repository) GetUser(ctx context.Context, id string) (models.User, error) {
	collection := r.usersCollection()

	c := collection.FindOne(ctx, bson.M{"_id": id})
	u := models.User{}

	err := c.Decode(&u)

	if errors.Is(err, mongo.ErrNoDocuments) {
		err = fmt.Errorf("user %s: %w", id, service.ErrNotFound)
		return u, err
	}

	if err != nil {
		return u, err
	}

	return u, nil
}

// UpdateUser implements service.Repository.
func (r *repository) UpdateUser(ctx context.Context, user models.User) error {
	panic("unimplemented")
}

// AddExpense implements service.Repository.
func (r *repository) AddExpense(ctx context.Context, exp models.Expense) error {
	panic("unimplemented")
}

// AddRoom implements service.Repository.
func (r *repository) AddRoom(ctx context.Context, room models.Room) error {
	panic("unimplemented")
}

// DeleteExpense implements service.Repository.
func (r *repository) DeleteExpense(ctx context.Context, id string) (models.Expense, error) {
	panic("unimplemented")
}

// GetExpense implements service.Repository.
func (r *repository) GetExpense(ctx context.Context, id string) (models.Expense, error) {
	panic("unimplemented")
}

// GetRoom implements service.Repository.
func (r *repository) GetRoom(ctx context.Context, id string) (models.Room, error) {
	panic("unimplemented")
}

// UpdateExpense implements service.Repository.
func (r *repository) UpdateExpense(ctx context.Context, exp models.Expense) error {
	panic("unimplemented")
}

// UpdateRoom implements service.Repository.
func (r *repository) UpdateRoom(ctx context.Context, room models.Room) error {
	panic("unimplemented")
}

func NewRepository(uri string) (service.Repository, error) {
	opts := options.Client().
		ApplyURI(uri).
		SetBSONOptions(&options.BSONOptions{
			UseJSONStructTags:   true,
			OmitZeroStruct:      true,
			NilMapAsEmpty:       true,
			NilSliceAsEmpty:     true,
			NilByteSliceAsEmpty: true,
		}).
		SetMaxPoolSize(10)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	c, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("connect: %w", err)
	}

	if err := c.Ping(ctx, readpref.Primary()); err != nil {
		_ = c.Disconnect(ctx)
		return nil, fmt.Errorf("ping: %w", err)
	}

	return &repository{
		c: c,
	}, nil
}

func (r *repository) usersCollection() *mongo.Collection {
	const collection = "users"
	return r.c.Database("testing").Collection(collection)
}
