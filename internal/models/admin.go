package models

type Admin struct {
	ID          int    `json:"id"`
	Firstname   string `json:"firstname"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phone_number"`
	AdminRole   string `json:"admin_role"`
	AdminStatus string `json:"admin_status"`
	Password    string `json:"password"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (*Admin) TableName() string {
	return "admins"
}
