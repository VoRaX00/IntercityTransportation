package services

import "time"

type UserRegister struct {
	Email        string `json:"email"`
	Lastname     string `json:"lastname"`
	Firstname    string `json:"firstname"`
	Patronymic   string `json:"patronymic"`
	PasswordHash string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AddPlace struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type AddSchedule struct {
	From          string    `json:"from"`
	To            string    `json:"to"`
	Departure     time.Time `json:"departure"`
	Arrival       time.Time `json:"arrival"`
	VehicleNumber string    `json:"vehicleNumber"`
}

type AddTransport struct {
	Model       string `json:"model"`
	StateNumber string `json:"stateNumber"`
	Type        string `json:"type"`
	CountSeats  int    `json:"countSeats"`
}
