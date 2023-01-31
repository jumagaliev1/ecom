package data

import (
	"time"
)

type Item struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"-"`
	Title       string    `json:"title"`
	Price       int64     `json:"price,omitempty,string"`
	IsPurchased bool      `json:"isPurchased"`
	Category    int64     `json:"category,omitempty"`
	Rating      uint8     `json:"rating,omitempty"`
	Version     int32     `json:"version"`
}
