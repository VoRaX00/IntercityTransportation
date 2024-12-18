package models

type Transport struct {
	StateNumber string `json:"stateNumber" db:"state_number"`
	Model       Model  `json:"model" db:"model"`
}
