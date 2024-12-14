package handler

type Service struct {
	Auth      Auth
	Place     Place
	Schedule  Schedule
	Transport Transport
}

func NewService() *Service {
	return &Service{}
}
