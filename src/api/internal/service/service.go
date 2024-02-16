package service_api

type ServiceCounter interface {
	Concate(a, b int) int
}

type Service struct {
	counter ServiceCounter
}

func NewService(counter ServiceCounter) *Service {
	return &Service{
		counter: counter,
	}
}

func (s *Service) ConcateLogic(a, b int) int {
	result := s.counter.Concate(a, b)
	return result
}
