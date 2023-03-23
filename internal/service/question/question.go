package question_srvc

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

func (s *Service) CreateQuestion(c context.Context, data *Create) (*entity.Question, error) {
	return s.repo.CreateQuestion(c, data)
}
func (s *Service) QuestionByID(c context.Context, id int) (*entity.Question, error) {
	return s.repo.QuestionByID(c, id)
}
func (s *Service) QuestionByIDs(c context.Context, ids []int) ([]entity.Question, error) {
	return s.repo.QuestionByIDs(c, ids)
}
func (s *Service) UpdateQuestion(c context.Context, id int, data *Update) error {
	return s.repo.UpdateQuestion(c, id, data)
}
func (s *Service) ListQuestion(c context.Context, size, page int) ([]entity.Question, error) {
	return s.repo.ListQuestion(c, size, page)
}
func (s *Service) DeleteQuestion(c context.Context, id, userID int) error {
	return s.repo.DeleteQuestion(c, id, userID)
}
func (s *Service) Service_() {
	s.repo.Repository_()
}
