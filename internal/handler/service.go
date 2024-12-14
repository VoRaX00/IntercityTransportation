package handler

type Service struct {
	Place    Place
	Schedule Schedule
}

func NewService() *Service {
	return &Service{}
}
