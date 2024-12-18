package models

import "time"

type Schedule struct {
	Id        int64     `json:"id" db:"id"`
	From      Place     `json:"from" db:"from"`
	To        Place     `json:"to" db:"to"`
	Departure time.Time `json:"departure" db:"departure"`
	Arrival   time.Time `json:"arrival" db:"arrival"`
	Vehicle   Transport `json:"vehicle" db:"state_number"`
}
