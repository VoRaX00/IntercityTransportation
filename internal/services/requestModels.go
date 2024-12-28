package services

type UserLogin struct {
	PhoneNumber int64  `json:"phoneNumber"`
	FIO         string `json:"fio"`
}

type AddPlace struct {
	Name string `json:"name" validate:"required,city_name"`
	Type string `json:"type"`
}

type AddFlight struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Departure   string `json:"departure"`
	Arrival     string `json:"arrival"`
	StateNumber string `json:"stateNumber"`
}

type AddBus struct {
	StateNumber string `json:"stateNumber"`
	Model       string `json:"model"`
}

type BuyTicket struct {
	Cost        int64   `json:"cost"`
	Flights     []int64 `json:"flights"`
	PhoneNumber int64   `json:"phoneNumber"`
}

type FlightFilter struct {
	From        string `json:"from"`
	To          string `json:"to"`
	StateNumber string `json:"stateNumber"`
}
