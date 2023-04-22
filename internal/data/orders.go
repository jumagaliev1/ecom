package data

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jumagaliev1/internal/validator"
	"time"
)

const (
	OrderStatusCreated = "CREATED"
	OrderStatusCancel  = "CANCEL"
	OrderStatusFinish  = "FINISH"
)

type Order struct {
	ID          int64      `json:"id"`
	OrderStatus string     `json:"order_status"`
	Cart        Cart       `json:"cart_id"`
	Quantity    int        `json:"quantity"`
	TotalPrice  int        `json:"total_price"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type OrderReq struct {
	CartID   int `json:"cart_id"`
	Quantity int `json:"quantity"`
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
func (m OrderModel) GetByID(ID int) (*Order, error) {
	query := `SELECT id, cart_id, order_status, quantity, total_price, created_at, updated_at, deleted_at
				FROM orders 
				WHERE id = $1`

	var order Order

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, ID).Scan(
		&order.ID,
		&order.Cart.ID,
		&order.OrderStatus,
		&order.Quantity,
		&order.TotalPrice,
		&order.CreatedAt,
		&order.UpdatedAt,
		&order.DeletedAt)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &order, nil
}
func (m OrderModel) GetAll() ([]Order, error) {
	var orders []Order

	query := `
			SELECT id, cart_id, order_status, quantity, total_price, created_at, updated_at, deleted_at 
			FROM orders`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var order Order
		err := rows.Scan(
			&order.ID,
			&order.Cart.ID,
			&order.Quantity,
			&order.TotalPrice,
			&order.CreatedAt,
			&order.UpdatedAt,
			&order.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
