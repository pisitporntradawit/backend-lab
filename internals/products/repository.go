package products

import (
    "context"
    "fmt"

    "github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
    db *pgxpool.Pool 
}

func NewRepository(db *pgxpool.Pool) *Repository {
    return &Repository{db: db}
}

func (r *Repository) GetProducts(ctx context.Context) ([]ProductsModel, error) {
    rows, err := r.db.Query(ctx, `
        SELECT idnumber, name, image
        FROM products
    `)
    if err != nil {
        return nil, fmt.Errorf("repository.GetProducts query: %w", err)
    }
    defer rows.Close()

    var products []ProductsModel
    for rows.Next() {
        var p ProductsModel
        if err := rows.Scan(&p.Idnumber, &p.Name, &p.Image); err != nil {
            return nil, fmt.Errorf("repository.GetProducts scan: %w", err)
        }
        products = append(products, p)
    }

    // ตรวจ error จาก iteration เสมอ
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("repository.GetProducts rows: %w", err)
    }

    return products, nil
}
