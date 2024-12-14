package handler

type Service struct {
	Place    Place
	Schedule Schedule
	Auth     Auth
}

func NewService() *Service {
	return &Service{}
}
