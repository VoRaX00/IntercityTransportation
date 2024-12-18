package models

type Ticket struct {
	Id          int64      `json:"id" db:"id"`
	Cost        int64      `json:"cost" db:"cost"`
	Schedules   []Schedule `json:"schedule" db:"schedule"`
	PhoneNumber int64      `json:"phoneNumber" db:"phone_number"`
}
