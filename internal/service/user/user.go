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

func (s *Service) GetByUsername(c context.Context, username string) (*entity.User, error) {
	return s.repo.GetByUsername(c, username)
}
func (s *Service) GetByID(c context.Context, userID int) (*entity.User, error) {
	return s.repo.GetByID(c, userID)
}
func (s *Service) Update(c context.Context, userID int, ua *Update) error {
	return s.repo.Update(c, userID, ua)
}
func (s *Service) Create(c context.Context, user *entity.User) error {
	return s.repo.Create(c, user)
}

func (s *Service) Service_() {
	s.repo.Repository_()
}
