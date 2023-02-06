package data

import (
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
