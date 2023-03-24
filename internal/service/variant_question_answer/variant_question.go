package variant_question_answer_srvc

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

func (s *Service) GetAll(c context.Context, filter *Filter) ([]entity.VariantQuestionAnswer, int, error) {
	return s.repo.GetAll(c, filter)
}

func (s *Service) Create(c context.Context, data *Create) (*entity.VariantQuestionAnswer, error) {
	return s.repo.Create(c, data)
}

func (s *Service) Service_() {
	s.repo.Repository_()
}
