package usecase

import (
	"learn-golang/syntax/day_5/domain"
	"time"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(name string) error {
	time.Sleep(3 * time.Second)
	user := domain.User{Name: name}
	return s.repo.Save(user)
}
