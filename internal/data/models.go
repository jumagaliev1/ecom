package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Products ProductModel
	Users    UserModel
	Carts    CartModel
	Orders   OrderModel
	Comments CommentModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Products: ProductModel{DB: db},
		Users:    UserModel{DB: db},
		Carts:    CartModel{DB: db},
		Orders:   OrderModel{DB: db},
		Comments: CommentModel{DB: db},
	}
}
