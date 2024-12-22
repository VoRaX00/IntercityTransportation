package models

type Ticket struct {
	Id          int64    `json:"id" db:"id"`
	Cost        int64    `json:"cost" db:"cost"`
	Schedules   []Flight `json:"flight" db:"flights"`
	PhoneNumber int64    `json:"phoneNumber" db:"users"`
}
