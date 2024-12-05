package models

import "time"

type Schedule struct {
	Id        int64     `json:"id"`
	From      Place     `json:"from"`
	To        Place     `json:"to"`
	Departure time.Time `json:"departure"`
	Arrival   time.Time `json:"arrival"`
	Vehicle   Transport `json:"vehicle"`
}
