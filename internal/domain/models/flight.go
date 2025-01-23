package models

import "time"

type Flight struct {
	Id        int64     `json:"id" db:"id"`
	From      Place     `json:"from" db:"places"`
	To        Place     `json:"to" db:"to_places"`
	Departure time.Time `json:"departure" db:"departure"`
	Arrival   time.Time `json:"arrival" db:"arrival"`
	Bus       Bus       `json:"bus" db:"buses"`
}
