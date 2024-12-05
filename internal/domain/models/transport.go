package models

type Transport struct {
	Model       string        `json:"model"`
	StateNumber string        `json:"stateNumber"`
	Type        TypeTransport `json:"type"`
	CountSeats  int           `json:"countSeats"`
}
