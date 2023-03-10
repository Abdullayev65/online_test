package topic_service

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

func (s *Service) CreateTopic(c context.Context, ent *entity.Topic) error {
	return s.repo.CreateTopic(c, ent)
}
func (s *Service) TopicByID(c context.Context, id int) (*entity.Topic, error) {
	return s.repo.TopicByID(c, id)
}
func (s *Service) UpdateTopic(c context.Context, id int, update *Update) error {
	return s.repo.UpdateTopic(c, id, update)
}
func (s *Service) ListTopic(c context.Context, size, page int) ([]entity.Topic, error) {
	return s.repo.ListTopic(c, size, page)
}

func (s *Service) DeleteTopic(c context.Context, id int, userID int) error {
	return s.repo.DeleteTopic(c, id, userID)
}
