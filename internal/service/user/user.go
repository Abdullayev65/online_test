package user

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) UserByUsername(c context.Context, username string) (*entity.User, error) {
	return s.repo.UserByUsername(c, username)
}
func (s *Service) UserByID(c context.Context, userID int) (*entity.User, error) {
	return s.repo.UserByID(c, userID)
}
func (s *Service) UserUpdate(c context.Context, userID int, ua *Update) error {
	return s.repo.UserUpdate(c, userID, ua)
}
func (s *Service) AddUser(c context.Context, user *entity.User) error {
	return s.repo.AddUser(c, user)
}
