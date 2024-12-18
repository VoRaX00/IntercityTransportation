package models

type User struct {
	PhoneNumber int64  `json:"phoneNumber" db:"phone_number"`
	FIO         string `json:"fio" db:"fio"`
}
