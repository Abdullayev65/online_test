package topic_srvc

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

func (s *Service) GetAll(c context.Context, filter *Filter) ([]entity.Topic, int, error) {
	return s.repo.GetAll(c, filter)
}
func (s *Service) GetByID(c context.Context, id int) (*entity.Topic, error) {
	return s.repo.GetByID(c, id)
}
func (s *Service) Create(c context.Context, ent *entity.Topic) error {
	return s.repo.Create(c, ent)
}
func (s *Service) Update(c context.Context, id int, update *Update) error {
	return s.repo.Update(c, id, update)
}
func (s *Service) Delete(c context.Context, id int, userID int) error {
	return s.repo.Delete(c, id, userID)
}
func (s *Service) Service_() {
	s.repo.Repository_()
}
