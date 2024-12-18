package auth

import "kursachDB/internal/services"

type Repo interface {
	AddUser(user services.UserRegister) error
}
