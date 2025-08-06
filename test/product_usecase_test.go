package test

import (
	"context"
	"testing"
	"time"

	"chekisvc/internal/domain/entity"
	"chekisvc/internal/usecase/product"
	"chekisvc/pkg/logger"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockProductRepository is a mock type for the ProductRepository type
type MockProductRepository struct {
	mock.Mock
}

// Create is a mock method
func (m *MockProductRepository) Create(ctx context.Context, product *entity.CreateProductRequest) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

// GetByID is a mock method
func (m *MockProductRepository) GetByID(ctx context.Context, id uint) (*entity.Products, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.Products), args.Error(1)
}

// Update is a mock method
func (m *MockProductRepository) Update(ctx context.Context, product *entity.CreateProductRequest, productID uint) error {
	args := m.Called(ctx, product, productID)
	return args.Error(0)
}

// Delete is a mock method
func (m *MockProductRepository) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// List is a mock method
func (m *MockProductRepository) List(ctx context.Context, limit, offset int) ([]*entity.Products, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]*entity.Products), args.Error(1)
}

// Count is a mock method
func (m *MockProductRepository) Count(ctx context.Context) (int64, error) {
	args := m.Called(ctx)
	return args.Get(0).(int64), args.Error(1)
}

func TestProductUsecase_CreateProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	logger := logger.NewLogger() // assuming you have a logger constructor
	uc := product.NewProductUsecase(mockRepo, logger)

	product := &entity.CreateProductRequest{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       100,
	}

	mockRepo.On("Create", mock.Anything, product).Return(nil)

	err := uc.CreateProduct(context.Background(), product)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductUsecase_GetProductByID(t *testing.T) {
	mockRepo := new(MockProductRepository)
	logger := logger.NewLogger()
	uc := product.NewProductUsecase(mockRepo, logger)

	product := &entity.Products{
		ID:          1,
		Name:        "Test Product",
		Description: "Test Description",
		Price:       100,
		CategoryID:  1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mockRepo.On("GetByID", mock.Anything, uint(1)).Return(product, nil)

	result, err := uc.GetProductByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestProductUsecase_UpdateProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	logger := logger.NewLogger()
	uc := product.NewProductUsecase(mockRepo, logger)

	product := &entity.CreateProductRequest{
		Name:        "Updated Product",
		Description: "Updated Description",
		Price:       200,
	}

	mockRepo.On("Update", mock.Anything, product, uint(1)).Return(nil)

	err := uc.UpdateProduct(context.Background(), product, 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductUsecase_DeleteProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	logger := logger.NewLogger()
	uc := product.NewProductUsecase(mockRepo, logger)

	mockRepo.On("Delete", mock.Anything, uint(1)).Return(nil)

	err := uc.DeleteProduct(context.Background(), 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductUsecase_ListProducts(t *testing.T) {
	mockRepo := new(MockProductRepository)
	logger := logger.NewLogger()
	uc := product.NewProductUsecase(mockRepo, logger)

	products := []*entity.Products{
		{
			ID:          1,
			Name:        "Test Product 1",
			Description: "Test Description 1",
			Price:       100,
			CategoryID:  1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Name:        "Test Product 2",
			Description: "Test Description 2",
			Price:       200,
			CategoryID:  1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	mockRepo.On("List", mock.Anything, 10, 0).Return(products, nil)

	result, err := uc.ListProducts(context.Background(), 10, 0)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}

func TestProductUsecase_CountProducts(t *testing.T) {
	mockRepo := new(MockProductRepository)
	logger := logger.NewLogger()
	uc := product.NewProductUsecase(mockRepo, logger)

	mockRepo.On("Count", mock.Anything).Return(int64(10), nil)

	count, err := uc.CountProducts(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, int64(10), count)
	mockRepo.AssertExpectations(t)
}
