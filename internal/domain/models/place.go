package models

type Place struct {
	Id        int64     `json:"id" db:"id"`
	NamePlace string    `json:"name" db:"name_place"`
	Type      TypePlace `json:"type" db:"type"`
}
