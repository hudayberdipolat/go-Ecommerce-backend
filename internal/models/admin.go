package models

type Admin struct {
	ID            int    `json:"id"`
	Username      string `json:"username"`
	FullName      string `json:"full_name"`
	PhoneNumber   string `json:"phone_number"`
	Email         string `jsin:"email"`
	AdminImageURL string `json:"user_image_url"`
	AdminStatys   string `json:"admin_status"`
	Password      string `json:"password"`
	CrearedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
