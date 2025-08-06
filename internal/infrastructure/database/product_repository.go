package database

import (
	"chekisvc/internal/domain/entity"
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type ProductPostgresRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductPostgresRepository {
	return &ProductPostgresRepository{
		db: db,
	}

}

func (r *ProductPostgresRepository) List(ctx context.Context, limit, offset int) ([]*entity.Products, error) {
	var products []*entity.Products

	query := `SELECT * FROM products LIMIT $1 OFFSET $2`
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err

	}
	defer rows.Close()
	for rows.Next() {
		var product entity.Products
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (r *ProductPostgresRepository) Create(ctx context.Context, product *entity.CreateProductRequest) error {
	query := `INSERT INTO products (name, description, price, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW())`
	_, err := r.db.ExecContext(ctx, query, product.Name, product.Description, product.Price)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductPostgresRepository) GetByID(ctx context.Context, id uint) (*entity.Products, error) {
	var product entity.Products
	err := r.db.QueryRowContext(ctx, "SELECT id, name, description, price, category_id, created_at, updated_at FROM products WHERE id = $1", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductPostgresRepository) Update(ctx context.Context, product *entity.CreateProductRequest, id uint) error {
	query := `UPDATE products SET name = $1, description = $2, price = $3, updated_at = NOW() WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, product.Name, product.Description, product.Price, id)
	if err != nil {
		return err
	}

	return nil
}

// Add other methods as required by the repository.ProductRepository interface

// Count returns the total number of products in the database.
func (r *ProductPostgresRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM products").Scan(&count); err != nil {
		return 0, err

	}
	return count, nil
}

func (r *ProductPostgresRepository) Delete(ctx context.Context, id uint) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
