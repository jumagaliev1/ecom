package data

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jumagaliev1/internal/validator"
	"github.com/lib/pq"
	"time"
)

type Comment struct {
	ID        int64      `json:"id"`
	User      User       `json:"user"`
	Product   Product    `json:"-"`
	Message   string     `json:"message"`
	Rating    uint8      `json:"rating"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

func ValidateComment(v *validator.Validator, comment *Comment) {
	v.Check(comment.Message != "", "message", "must be provided")
	v.Check(comment.Rating != 0, "rating", "must be provided")
}

type CommentModel struct {
	DB *sql.DB
}

func (m CommentModel) Insert(comment *Comment) error {
	query := `INSERT INTO comments (user_id, product_id, comment, rating) 
			VALUES ($1,$2,$3,$4)
			RETURNING id`

	args := []interface{}{comment.User.ID, comment.Product.ID, comment.Message, comment.Rating}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&comment.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m CommentModel) GetByProduct(p *Product) ([]Comment, error) {
	query := `SELECT c.id, 
       	u.id, u.last_name, u.first_name, u.email, u.phone, u.address, u.role, u.created_at, u.updated_at, u.deleted_at,
		p.id, p.category_id, p.user_id, p.title, p.description, p.price, p.rating, p.stock, p.images, p.created_at, p.updated_at, p.deleted_at,
		c.comment, c.rating, c.created_at, c.updated_at, c.deleted_at
			FROM comments c JOIN users u on c.user_id = u.id JOIN products p on p.id = c.product_id
			WHERE product_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var comments []Comment

	rows, err := m.DB.QueryContext(ctx, query, p.ID)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.ID,
			&comment.User.ID, &comment.User.LastName, &comment.User.FirstName, &comment.User.Email, &comment.User.Phone, &comment.User.Address, &comment.User.Role, &comment.User.CreatedAt, &comment.User.UpdatedAt, &comment.User.DeletedAt,
			&comment.Product.ID, &comment.Product.Category, &comment.Product.User, &comment.Product.Title, &comment.Product.Description, &comment.Product.Price, &comment.Product.Rating, &comment.Product.Stock, pq.Array(&comment.Product.Images), &comment.Product.CreatedAt, &comment.Product.UpdatedAt, &comment.Product.DeletedAt,
			&comment.Message, &comment.Rating, &comment.CreatedAt, &comment.UpdatedAt, &comment.DeletedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
