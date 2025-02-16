package repositories

import (
	"context"
	"database/sql"

	"www.github.com/felipegith/internal/domain"
)


type OrderRepositoryDatabase struct {
	db *sql.DB
}

func Constructor(db *sql.DB) * OrderRepositoryDatabase{
	return &OrderRepositoryDatabase{db: db}
}


func (r *OrderRepositoryDatabase) Create(ctx context.Context, order *domain.Order) error {
	query := `INSERT INTO orders (id, name, email) VALUES (?, ?, ?)`
	_, err:= r.db.ExecContext(ctx, query, order.Id, order.Name, order.Email)
	return err
}