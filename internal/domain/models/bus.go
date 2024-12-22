package models

type Bus struct {
	StateNumber string `json:"stateNumber" db:"state_number"`
	Model       Model  `json:"model" db:"models"`
}
