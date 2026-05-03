package login

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct{
	db *pgxpool.Pool // เก็บ connection pool ไว้ใช้คุยกับ Database
}

func NewRepository(db *pgxpool.Pool) *Repository{
	return &Repository{
		db : db, 
	}
}