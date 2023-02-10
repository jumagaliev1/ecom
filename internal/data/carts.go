package data

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jumagaliev1/internal/validator"
	"time"
)

type Cart struct {
	ID       int64   `json:"id"`
	User     User    `json:"user"`
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
}

func ValidateCart(v *validator.Validator, c *Cart) {
	//TO-DO
	//v.Check(reflect.DeepEqual(c.User, User{}), "user", "must be provided")
	//v.Check(c.Product != nil, "product", "must be provided")
	//
}

type CartModel struct {
	DB *sql.DB
}

func (m CartModel) Insert(cart *Cart) error {
	query := `INSERT INTO carts (user_id, product_id, quantity) 
				VALUES ($1, $2, $3)
				RETURNING id`

	args := []interface{}{cart.User.ID, cart.Product.ID}

	return m.DB.QueryRow(query, args...).Scan(&cart.ID)
}

func (m CartModel) GetByUser(user *User) (*Cart, error) {
	query := `SELECT id, user_id, product_id, quantity 
				FROM carts
				WHERE user_id = $1`

	var cart Cart

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, user.ID).Scan(
		&cart.ID,
		&cart.User.ID,
		&cart.Product.ID,
		&cart.Quantity)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &cart, nil
}
