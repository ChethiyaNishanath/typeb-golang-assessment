package greeting

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Greet(name string) (string, error) {
	if err := ValidateName(name); err != nil {
		return "", err
	}
	return "Hello " + name, nil
}
