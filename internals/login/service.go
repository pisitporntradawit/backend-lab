package login

import (

)

type LoginRepository interface {

}

type Service struct {
	repo LoginRepository
}

func NewService(repo *Repository) *Service{
	return &Service{
		repo : repo,
	}
}