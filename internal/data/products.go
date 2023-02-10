package data

import (
	"database/sql"
	"github.com/jumagaliev1/internal/validator"
	"github.com/lib/pq"
	"time"
)

type Product struct {
	ID          int64          `json:"id"`
	Category    int32          `json:"category,omitempty"`
	User        int64          `json:"user"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Price       Price          `json:"price,omitempty"`
	Rating      uint8          `json:"rating,omitempty"`
	Stock       int            `json:"stock"`
	Images      pq.StringArray `json:"images"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   time.Time      `json:"-"`
}

func ValidateItem(v *validator.Validator, p *Product) {
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
	return nil, nil
}

func (m ProductModel) Update(movie *Product) error {
	return nil
}

func (m ProductModel) Delete(id int64) error {
	return nil
}
