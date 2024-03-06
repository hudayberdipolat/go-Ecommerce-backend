package models

type ProductComment struct {
	ID             int     `json:"id"`
	ProductComment string  `json:"product_Comment"`
	UserID         int     `json:"user_id"`
	ProductID      int     `json:"product_id"`
	User           User    `json:"user"`
	Product        Product `json:"product"`
}

func (*ProductComment) TableName() string {
	return "product_comments"
}
