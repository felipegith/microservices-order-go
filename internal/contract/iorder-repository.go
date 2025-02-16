package contract

import (
	"context"

	"www.github.com/felipegith/internal/domain"
)


type OrderRepository interface {
	Create(ctx context.Context, order *domain.Order) error
}