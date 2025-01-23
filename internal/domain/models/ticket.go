package models

type Ticket struct {
	Id     int64    `json:"id" db:"id"`
	Cost   int64    `json:"cost" db:"cost"`
	Flight []Flight `json:"flights" db:"flights"`
	User   User     `json:"user" db:"users"`
}
