package answer_srvc

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

func (s *Service) CreateAnswer(c context.Context, data *Create, userID, questionID int) (*entity.Answer, error) {
	return s.repo.CreateAnswer(c, data, userID, questionID)
}
func (s *Service) AnswerByID(c context.Context, id int) (*entity.Answer, error) {
	return s.repo.AnswerByID(c, id)
}
func (s *Service) UpdateAnswer(c context.Context, data *Update) error {
	return s.repo.UpdateAnswer(c, data)
}
func (s *Service) AnswersByQuestionID(c context.Context, questionID int) ([]entity.Answer, error) {
	return s.repo.AnswersByQuestionID(c, questionID)
}
func (s *Service) DeleteAnswer(c context.Context, id, userID int) error {
	return s.repo.DeleteAnswer(c, id, userID)
}
func (s *Service) CorrectAnswerByQuestionID(c context.Context, questionID int) (*entity.Answer, error) {
	return s.repo.CorrectAnswerByQuestionID(c, questionID)
}
func (s *Service) Service_() {
	s.repo.Repository_()
}
