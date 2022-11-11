package profile

import (
	"mime/multipart"

	"github.com/SolidShake/photoclub/pkg/storage"
)

type Service struct {
	storage    *storage.FileStorage
	repository *Repository
}

func NewService(storage *storage.FileStorage, repository *Repository) *Service {
	return &Service{storage: storage, repository: repository}
}

func (s *Service) GetProfile(userID string) (Profile, error) {
	return s.repository.GetProfile(userID)
}

func (s *Service) GetLogoPath(logo string) string {
	return s.storage.GetFileLink(logo)
}

func (s *Service) SaveLogo(saveFunc storage.SaveFunc, file *multipart.FileHeader) (string, error) {
	return s.storage.SaveFile(saveFunc, file)
}

func (s *Service) UpdateProfile(userID, userType, logo, about string) error {
	return s.repository.UpdateProfile(userID, userType, logo, about)
}
