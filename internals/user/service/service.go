package service

import (
	"api/internals/user/model"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// Interface ให้ test ได้โดยไม่ต้องมี DB จริง
type UserRepository interface {
	Getuser(ctx context.Context) ([]model.UserModel, error)
    CreateUser(ctx context.Context, newUser *model.UserModel) error
}

type Service struct {
	repo UserRepository
}

func NewService(repo UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Getuser(ctx context.Context) ([]model.UserModel, error) {
	resultUser, err := s.repo.Getuser(ctx)
	if err != nil {
		return nil, fmt.Errorf("service.getUser: %w", err)
	}
	return resultUser, nil
}

func (s *Service) CreateUser(ctx context.Context, newUser *model.UserModel) error {
	if newUser == nil {
		return fmt.Errorf("service.user cannot be nil")
	}

	hashPassword, err := bcrypt.GenerateFromPassword(
		[]byte(newUser.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return fmt.Errorf("service.create password failed: %w", err)
	}

	newUser.Password = string(hashPassword)
	if err = s.repo.CreateUser(ctx, newUser); err != nil {
		return fmt.Errorf("service.create user failed %w", err)
	}
    return nil
}
