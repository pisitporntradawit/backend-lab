package login

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginRepository interface {
	Login(ctx context.Context, Username string) (*UserLogin, error)
}

type Service struct {
	repo LoginRepository
}

func NewService(repo LoginRepository) *Service{
	return &Service{
		repo : repo,
	}
}

func (s *Service) Login(ctx context.Context, Username string, Password string) (string, error) {
	req, err := s.repo.Login(ctx, Username)
	if err != nil {
		return "",fmt.Errorf("user not found %w", err)
	}

	if bcrypt.CompareHashAndPassword([]byte(req.Password),[]byte(Password)) != nil {
		return "" , errors.New("Invaild password")
	}

	secret := os.Getenv("JWTSECRET")
	if secret == "" {
		return "", errors.New("JWT secret not set")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username" : req.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}