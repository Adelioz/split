package service

import (
	"context"

	"github.com/Adelioz/split/internal/models"
	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// User

func (s *Service) AddUser(ctx context.Context, user models.User) error {
	err := s.repo.AddUser(ctx, user)
	id := uuid.New()
	user.ID = id.String()
	return err
}

func (s *Service) UpdateUser(ctx context.Context, user models.User) error {
	err := s.repo.UpdateUser(ctx, user)
	return err
}

func (s *Service) GetUser(ctx context.Context, id string) (models.User, error) {
	user, err := s.repo.GetUser(ctx, id)
	return user, err
}

// Room

func (s *Service) AddRoom(ctx context.Context, room models.Room) error {
	err := s.repo.AddRoom(ctx, room)
	return err
}

func (s *Service) UpdateRoom(ctx context.Context, room models.Room) error {
	err := s.repo.UpdateRoom(ctx, room)
	return err
}

func (s *Service) GetRoom(ctx context.Context, id string) (models.Room, error) {
	room, err := s.repo.GetRoom(ctx, id)
	return room, err
}

// Expes

func (s *Service) AddExpense(ctx context.Context, exp models.Expense) error {
	err := s.repo.AddExpense(ctx, exp)
	return err
}

func (s *Service) UpdateExpense(ctx context.Context, exp models.Expense) error {
	err := s.repo.UpdateExpense(ctx, exp)
	return err
}

func (s *Service) GetExpense(ctx context.Context, id string) (models.Expense, error) {
	exp, err := s.repo.GetExpense(ctx, id)
	return exp, err
}

func (s *Service) DeleteExpense(ctx context.Context, id string) (models.Expense, error) {
	exp, err := s.repo.DeleteExpense(ctx, id)
	return exp, err
}
