package user

import (
	"errors"

	"github.com/SolidShake/photoclub/db"
	"golang.org/x/crypto/bcrypt"
)

var errorEmailUsed = errors.New("email already exists")
var errorInternal = errors.New("internal error")
var errorUserNotFound = errors.New("user not found")
var errorInvalidPassword = errors.New("invalid password")

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s Service) CreateUser(email, password string) error {
	if _, err := s.repository.GetUserByEmail(email); err != db.ErrNoMatch {
		return errorEmailUsed
	}

	password, err := s.hashPassword(password)
	if err != nil {
		return errorInternal
	}

	return s.repository.CreateUser(email, password)
}

func (s Service) GetUser(email, password string) (*User, error) {
	user, err := s.repository.GetUserByEmail(email)
	if err == db.ErrNoMatch {
		return nil, errorUserNotFound
	}
	if err != nil {
		return nil, errors.New("internal error")
	}

	if !s.CheckPassword(user.Password, password) {
		return nil, errorInvalidPassword
	}

	return &user, nil
}

func (s Service) CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s Service) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
