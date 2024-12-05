package models

type User struct {
	Id         int64  `json:"id"`
	Lastname   string `json:"lastname"`
	Firstname  string `json:"firstname"`
	Patronymic string `json:"patronymic"`
}
