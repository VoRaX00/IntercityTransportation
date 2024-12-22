package models

type Place struct {
	Id        int64     `json:"id" db:"id"`
	NamePlace string    `json:"name" db:"place_name"`
	Type      TypePlace `json:"type" db:"types_places"`
}
