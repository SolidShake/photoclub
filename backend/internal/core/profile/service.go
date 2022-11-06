package profile

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetProfile(userID string) (Profile, error) {
	return s.repository.GetProfile(userID)
}

func (s *Service) UpdateProfile(userID, userType, logo, about string) error {
	return s.repository.UpdateProfile(userID, userType, logo, about)
}
