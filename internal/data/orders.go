package data

import (
	"database/sql"
	"github.com/jumagaliev1/internal/validator"
	"time"
)

type Order struct {
	ID          int64     `json:"id"`
	OrderStatus string    `json:"order_status"`
	Cart        Cart      `json:"cart_id"`
	Quantity    int       `json:"quantity"`
	TotalPrice  int       `json:"total_price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

func ValidateOrder(v *validator.Validator, order *Order) {
}

type OrderModel struct {
	DB *sql.DB
}

func (m OrderModel) Insert(order *Order) error {
	query := `
			INSERT INTO orders (cart_id, quantity, total_price)
			VALUES ($1, $2, $3)
			RETURNING id`

	args := []interface{}{order.Cart.ID, order.Quantity, order.TotalPrice}

	return m.DB.QueryRow(query, args...).Scan(&order.ID)
}
