package user

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s Service) CreateUser(email, password string) error {
	return s.repository.CreateUser(email, password)
}
