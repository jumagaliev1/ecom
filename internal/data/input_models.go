package data

type InputCreateProduct struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       int      `json:"price"`
	Category    int32    `json:"category"`
	Stock       int      `json:"stock"`
	Images      []string `json:"images"`
}

type InputListProducts struct {
	Title    string
	Category int
	Filters
}

type InputComment struct {
	ProductID int    `json:"product_id"`
	Message   string `json:"message"`
	Rating    uint8  `json:"rating"`
}

type InputAuthUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type InputUpdateProduct struct {
	Category    *int32   `json:"category"`
	Title       *string  `json:"title"`
	Description *string  `json:"description"`
	Price       *int     `json:"price"`
	Rating      *float32 `json:"rating"`
	Stock       *int     `json:"stock"`
	Images      []string `json:"images"`
}
