package repository

import (
	"api/internals/user/model"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(ctx context.Context, newUser *model.UserModel) error {
	if newUser.Id == uuid.Nil {
		newUser.Id = uuid.New()
	}

	insertEmployeeSQL := `insert into users (id, name, password) values ($1, $2, $3) returning id`

	err := r.db.QueryRow(ctx, insertEmployeeSQL, newUser.Id, newUser.Name, newUser.Password).Scan(&newUser.Id)
	if err != nil {
		return fmt.Errorf("r.create user failed %w", err)
	}
	return nil

}

func (r *Repository) Getuser(ctx context.Context) ([]model.UserModel, error) {
	rows, err := r.db.Query(ctx, `
        SELECT id, name
        FROM users
    `)
	if err != nil {
		return nil, fmt.Errorf("repository.getUser query: %w", err)
	}
	defer rows.Close()

	var result []model.UserModel
	for rows.Next() {
		var p model.UserModel
		if err := rows.Scan(&p.Id, &p.Name); err != nil {
			return nil, fmt.Errorf("repository.getUser scan: %w", err)
		}
		result = append(result, p)
	}

	// ตรวจ error จาก iteration เสมอ
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repository.getUser rows: %w", err)
	}

	return result, nil
}

func (r *Repository) GetUserByID(ctx context.Context, id string) (*model.UserModel, error) {
	objID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	query := "SELECT id, name FROM users WHERE id = $1"

	var user model.UserModel
	err = r.db.QueryRow(ctx, query, objID).Scan(&user.Id, &user.Name)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	return &user, nil
}