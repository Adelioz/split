package mock

import (
	"context"

	"github.com/Adelioz/split/internal/models"
	"github.com/Adelioz/split/internal/service"
)

type repository struct{}

// AddUser implements service.Repository.
func (r *repository) AddUser(ctx context.Context, user models.User) error {
	return nil
}

// GetUser implements service.Repository.
func (r *repository) GetUser(ctx context.Context, id string) (models.User, error) {
	panic("unimplemented")
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

func NewRepository() (service.Repository, error) {
	return &repository{}, nil
}
