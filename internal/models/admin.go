package models

type Admin struct {
}

func (*Admin) TableName() string {
	return "admins"
}
