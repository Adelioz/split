package service

import (
	"context"
	"errors"

	"github.com/Adelioz/split/internal/models"
)

var ErrNotFound = errors.New("not found")
var ErrAlreadyExists = errors.New("already exists")

type Repository interface {
	// Users

	AddUser(ctx context.Context, user models.User) error
	UpdateUser(ctx context.Context, user models.User) error
	GetUser(ctx context.Context, id string) (models.User, error)

	// Rooms

	AddRoom(ctx context.Context, room models.Room) error
	UpdateRoom(ctx context.Context, room models.Room) error
	GetRoom(ctx context.Context, id string) (models.Room, error)

	// Expenses

	AddExpense(ctx context.Context, exp models.Expense) error
	UpdateExpense(ctx context.Context, exp models.Expense) error
	GetExpense(ctx context.Context, id string) (models.Expense, error)
	DeleteExpense(ctx context.Context, id string) (models.Expense, error)
}
