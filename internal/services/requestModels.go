package services

type UserLogin struct {
	PhoneNumber string `json:"phoneNumber"`
	FIO         string `json:"fio"`
}

type AddPlace struct {
	Name string `json:"name"`
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
