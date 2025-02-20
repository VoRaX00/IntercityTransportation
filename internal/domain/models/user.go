package models

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
	"strings"
)

type User struct {
	PhoneNumber int64  `json:"phoneNumber" db:"phone_number" validate:"required,phone"`
	FIO         string `json:"fio" db:"fio" validate:"required,fullname"`
}

func PhoneValidator(fl validator.FieldLevel) bool {
	number := strconv.FormatInt(fl.Field().Int(), 10)
	if len(number) != 11 {
		return false
	}

	return strings.HasPrefix(number, "79") || strings.HasPrefix(number, "89")
}

func FullNameValidator(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[А-ЯA-Z][а-яa-z]+\s[А-ЯA-Z][а-яa-z]+\s[А-ЯA-Z][а-яa-z]+$`)
	return re.MatchString(fl.Field().String())
}
