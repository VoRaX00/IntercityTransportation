package models

type Model struct {
	Model      string `json:"model" db:"model"`
	CountPlace int    `json:"countPlace" db:"count_places"`
}
