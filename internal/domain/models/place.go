package models

type Place struct {
	Id   int64     `json:"id" db:"id"`
	Name string    `json:"name" db:"place_name"`
	Type TypePlace `json:"type" db:"type_place"`
}
