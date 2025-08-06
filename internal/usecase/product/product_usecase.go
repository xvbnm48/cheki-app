package product

import (
	"chekisvc/internal/domain/entity"
	repository "chekisvc/internal/domain/interface"
	"chekisvc/pkg/logger"
	"context"
	"time"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, product *entity.CreateProductRequest) error
	GetProductByID(ctx context.Context, id uint) (*entity.Products, error)
	UpdateProduct(ctx context.Context, product *entity.CreateProductRequest, productId uint) error
	DeleteProduct(ctx context.Context, id uint) error
	ListProducts(ctx context.Context, limit, offset int) ([]*entity.Products, error)
	CountProducts(ctx context.Context) (int64, error)
}

type productUsecase struct {
	productRepo repository.ProductRepository
	logger      logger.Logger
	timeout     time.Duration
}

func NewProductUsecase(productRepo repository.ProductRepository, logger logger.Logger) ProductUsecase {
	return &productUsecase{
		productRepo: productRepo,
		logger:      logger,
		timeout:     time.Second * 10,
	}
}

// CountProducts implements ProductUsecase.
func (p *productUsecase) CountProducts(ctx context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()
	return p.productRepo.Count(ctx)
}

// DeleteProduct implements ProductUsecase.
func (p *productUsecase) DeleteProduct(ctx context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()
	return p.productRepo.Delete(ctx, id)
}

// GetProductByID implements ProductUsecase.
func (p *productUsecase) GetProductByID(ctx context.Context, id uint) (*entity.Products, error) {
	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()
	return p.productRepo.GetByID(ctx, id)
}

// ListProducts implements ProductUsecase.
func (p *productUsecase) ListProducts(ctx context.Context, limit int, offset int) ([]*entity.Products, error) {
	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()

	return p.productRepo.List(ctx, limit, offset)
}

// UpdateProduct implements ProductUsecase.
func (p *productUsecase) UpdateProduct(ctx context.Context, product *entity.CreateProductRequest, productId uint) error {
	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()
	return p.productRepo.Update(ctx, product, productId)
}

func (p *productUsecase) CreateProduct(ctx context.Context, product *entity.CreateProductRequest) error {
	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()

	return p.productRepo.Create(ctx, product)
}
