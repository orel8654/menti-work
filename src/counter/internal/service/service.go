package service_counter

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s Service) Concate(a, b int) int {
	return a + b
}
