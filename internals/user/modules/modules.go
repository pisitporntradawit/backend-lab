package modules

import (
	"api/internals/user/handler"
	"api/internals/user/repository"
	"github.com/jackc/pgx/v5/pgxpool"
	"api/internals/user/service"
)

type Module struct {
	Handler *handler.Handler
}

func NewModule(db *pgxpool.Pool) *Module {
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	Handler := handler.NewHandler(svc)
	return &Module{
		Handler: Handler,
	}
}
