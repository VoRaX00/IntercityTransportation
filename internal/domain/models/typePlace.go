package models

type TypePlace struct {
	Id   int64  `json:"id" db:"type_id"`
	Type string `json:"type" db:"type_place"`
}
