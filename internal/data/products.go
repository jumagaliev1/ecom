package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jumagaliev1/internal/validator"
	"github.com/lib/pq"
	"time"
)

type Product struct {
	ID          int64     `json:"id"`
	Category    int32     `json:"category,omitempty"`
	User        int64     `json:"user"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       Price     `json:"price,omitempty"`
	Rating      uint8     `json:"rating,omitempty"`
	Stock       int       `json:"stock"`
	Images      []string  `json:"images"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	DeletedAt   time.Time `json:"-"`
}

func ValidateProduct(v *validator.Validator, p *Product) {
	v.Check(p.Title != "", "title", "must be provided")
	v.Check(len(p.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(p.Price != 0, "price", "must be provided")
	v.Check(p.Price >= 100, "price", "must be greater than 100")

	v.Check(p.Category != 0, "category", "must be provided")
	v.Check(p.Category > 0, "category", "must be a positive integer")
}

type ProductModel struct {
	DB *sql.DB
}

func (m ProductModel) Insert(product *Product) error {
	query := `
			INSERT INTO products (title, category_id, user_id, description, price, images)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id, created_at`

	args := []interface{}{product.Title, product.Category, product.User, product.Description, product.Price, pq.Array(product.Images)}

	return m.DB.QueryRow(query, args...).Scan(&product.ID, &product.CreatedAt)
}

func (m ProductModel) Get(id int64) (*Product, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
			SELECT id, category_id, user_id, title, description, price, rating, stock, images, created_at
			FROM products 
			WHERE id = $1`

	var product Product

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&product.ID,
		&product.Category,
		&product.User,
		&product.Title,
		&product.Description,
		&product.Price,
		&product.Rating,
		&product.Stock,
		pq.Array(&product.Images),
		&product.CreatedAt)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &product, nil
}

func (m ProductModel) Update(product *Product) error {
	query := `
		UPDATE products
		SET title = $1, category_id = $2, user_id = $3, description = $4, price = $5, rating = $6, stock = $7, images = $8, updated_at = now()
		WHERE id = $9 
		RETURNING updated_at`

	args := []interface{}{
		product.Title,
		product.Category,
		product.User,
		product.Description,
		product.Price,
		product.Rating,
		product.Stock,
		pq.Array(product.Images),
		product.ID,
	}

	return m.DB.QueryRow(query, args...).Scan(&product.UpdatedAt)
}

func (m ProductModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM products
		WHERE id = $1`

	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (m ProductModel) GetAll(title string, category int, filters Filters) ([]*Product, error) {
	query := fmt.Sprintf(`
			SELECT id, category_id, user_id, title, description, price, rating, stock, images, created_at
			FROM products
			WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
			AND (category_id = $2 or $2 = 0)
			ORDER BY %s %s, id ASC
			LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, title, category, filters.limit(), filters.offset())
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []*Product{}

	for rows.Next() {
		var product Product
		err := rows.Scan(
			&product.ID,
			&product.Category,
			&product.User,
			&product.Title,
			&product.Description,
			&product.Price,
			&product.Rating,
			&product.Stock,
			pq.Array(&product.Images),
			&product.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}
