package data

import (
	"database/sql"
	"github.com/jumagaliev1/internal/validator"
	"time"
)

type Comment struct {
	ID        int64     `json:"id"`
	User      User      `json:"user"`
	Product   Product   `json:"product"`
	Message   string    `json:"message"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

func ValidateComment(v *validator.Validator, comment *Comment) {
	v.Check(comment.Message != "", "message", "must be provided")
	v.Check(comment.Rating != 0, "rating", "must be provided")
}

type CommentModel struct {
	DB *sql.DB
}

func (m CommentModel) Insert(comment *Comment) error {
	return nil
}

func (m CommentModel) GetByProduct(p Product) ([]Comment, error) {
	return nil, nil
}
