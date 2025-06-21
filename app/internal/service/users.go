package service

import (
	"fmt"
	"github.com/cweiser22/urls-ac/internal/models"
	"github.com/cweiser22/urls-ac/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	usersRepository *repository.UserRepository
}

func NewUserService(usersRepository *repository.UserRepository) *UserService {
	return &UserService{
		usersRepository: usersRepository,
	}
}

func (s *UserService) Register(email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	dto := &repository.CreateUserDTO{
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	user, err := s.usersRepository.InsertUser(dto)
	if err != nil {
		return nil, fmt.Errorf("insert user: %w", err)
	}

	return user, nil
}
