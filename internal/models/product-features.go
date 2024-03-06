package models

type ProductFeature struct {
	ID        int     `json:"id"`
	Feature   string  `json:"feature"`
	ProductID int     `json:"product_id"`
	Product   Product `json:"product"`
}

func (*ProductFeature) TableName() string {
	return "product_features"
}
