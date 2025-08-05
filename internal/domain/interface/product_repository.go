package repository

import (
	"chekisvc/internal/domain/entity"
	"context"
)

type ProductRepository interface {
	Create(ctx context.Context, product *entity.CreateProductRequest) error
	GetByID(ctx context.Context, id uint) (*entity.Products, error)
	Update(ctx context.Context, product *entity.CreateProductRequest, productId uint) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*entity.Products, error)
	Count(ctx context.Context) (int64, error)
}
