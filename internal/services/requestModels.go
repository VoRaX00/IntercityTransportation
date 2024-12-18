package services

import "time"

type UserRegister struct {
	PhoneNumber string `json:"phoneNumber"`
	FIO         string `json:"fio"`
}

type UserLogin struct {
	PhoneNumber string `json:"phoneNumber"`
	FIO         string `json:"fio"`
}

type AddPlace struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type AddSchedule struct {
	From        string    `json:"from"`
	To          string    `json:"to"`
	Departure   time.Time `json:"departure"`
	Arrival     time.Time `json:"arrival"`
	StateNumber string    `json:"stateNumber"`
}

type AddTransport struct {
	StateNumber string `json:"stateNumber"`
	Model       string `json:"model"`
}
