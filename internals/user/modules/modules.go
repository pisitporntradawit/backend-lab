package modules

import (
	"api/internals/user/handler"
	"api/internals/user/repository"
	"github.com/jackc/pgx/v5/pgxpool"
	"api/internals/user/service"
)

type Module struct {
	Controller *handler.Controller
}

func NewModule(db *pgxpool.Pool) *Module {
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	controller := handler.NewController(svc)
	return &Module{
		Controller: controller,
	}
}
