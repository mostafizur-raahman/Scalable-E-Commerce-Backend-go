package repository

import (
	"context"
	"ecommerce/internal/database"
	"ecommerce/internal/domain/product"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type ProductRepository struct {
	db *database.DB
}

func NewProductRepository(db *database.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() ([]product.Product, error) {
	ctx := context.Background()
	rows, err := r.db.Pool.Query(ctx, "SELECT id, name, price FROM products ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("query all: %w", err)
	}
	defer rows.Close()

	var products []product.Product
	for rows.Next() {
		var p product.Product
		if err := rows.Scan(&p.Id, &p.Name, &p.Price); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}
		products = append(products, p)
	}
	return products, rows.Err()
}

func (r *ProductRepository) GetByID(id int64) (*product.Product, error) {
	var p product.Product
	err := r.db.Pool.QueryRow(context.Background(),
		"SELECT id, name, price FROM products WHERE id = $1", id).
		Scan(&p.Id, &p.Name, &p.Price)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("product not found")
	}
	return &p, err
}

func (r *ProductRepository) Create(p *product.Product) error {
	return r.db.Pool.QueryRow(context.Background(),
		"INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id",
		p.Name, p.Price).Scan(&p.Id)
}

func (r *ProductRepository) Update(p *product.Product) error {
	res, err := r.db.Pool.Exec(context.Background(),
		"UPDATE products SET name = $1, price = $2 WHERE id = $3",
		p.Name, p.Price, p.Id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return errors.New("product not found")
	}
	return nil
}

func (r *ProductRepository) Delete(id int64) error {
	res, err := r.db.Pool.Exec(context.Background(),
		"DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return errors.New("product not found")
	}
	return nil
}
