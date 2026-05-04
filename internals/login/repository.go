package login

import (
	"context"
	"errors"
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

func (r *Repository) Login(ctx context.Context, Username string) (*UserLogin, error) {
	var req UserLogin
	err := r.db.QueryRow(ctx,"select name, password from users where name = $1",Username).Scan(&req.Username, &req.Password)
	if err != nil{
		return nil, errors.New("User not found")
	}
	return &req,nil
}