package variant_question_srvc

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

func (s *Service) CreateVariantQuestion(c context.Context, data *Create, variantID int) ([]entity.VariantQuestion, error) {
	return s.repo.CreateVariantQuestion(c, data, variantID)
}

func (s *Service) VariantQuestionByID(c context.Context, id int) (*entity.VariantQuestion, error) {
	return s.repo.VariantQuestionByID(c, id)
}

func (s *Service) VariantQuestionByVariantID(c context.Context, variantID int) ([]entity.VariantQuestion, error) {
	return s.repo.VariantQuestionByVariantID(c, variantID)
}

func (s *Service) DeleteVariantQuestion(c context.Context, id int) error {
	return s.repo.DeleteVariantQuestion(c, id)
}

func (s *Service) Service_() {
	s.repo.Repository_()
}
