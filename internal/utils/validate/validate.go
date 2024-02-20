package validate

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(data interface{}) error {
	validate := validator.New()
	err := validate.Struct(data)
	return err
}

func PhoneNumberValidate(phoneNumber string) bool {
	if string(phoneNumber[3]) == "6" {
		r := regexp.MustCompile("[9936]{4}[1-5][0-9]{6}$")
		return r.MatchString(phoneNumber)
	}
	if string(phoneNumber[3]) == "7" {
		return string(phoneNumber[4]) == "1"
	}
	return false
}
