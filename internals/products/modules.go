package products

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Module struct {
	Controller *Controller
	Service    *Service
	Repository *Repository
}

func NewModule(db *pgxpool.Pool) *Module {
	//repo := NewRepository(db)
	service := NewService(&MockRepository{})
	controller := NewController(service)
	return &Module{
		Controller: controller,
	}
}
