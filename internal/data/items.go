package data

import (
	"github.com/jumagaliev1/internal/validator"
	"time"
)

type Item struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"-"`
	Title       string    `json:"title"`
	Price       Price     `json:"price,omitempty"`
	IsPurchased bool      `json:"isPurchased"`
	Category    int32     `json:"category,omitempty"`
	Rating      uint8     `json:"rating,omitempty"`
	Version     int32     `json:"version"`
}

func ValidateItem(v *validator.Validator, item *Item) {
	v.Check(item.Title != "", "title", "must be provided")
	v.Check(len(item.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(item.Price != 0, "price", "must be provided")
	v.Check(item.Price >= 100, "price", "must be greater than 100")

	v.Check(item.Category != 0, "category", "must be provided")
	v.Check(item.Category > 0, "category", "must be a positive integer")

}
