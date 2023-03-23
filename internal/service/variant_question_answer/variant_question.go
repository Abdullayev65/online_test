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

func (s *Service) Create(c context.Context, data *Create,
	userID int) (*entity.VariantQuestionAnswer, error) {

	return s.repo.Create(c, data, userID)
}

func (s *Service) GetByUserIDAndVariantID(c context.Context, userID,
	variantID int) ([]entity.VariantQuestionAnswer, error) {
	return s.repo.GetByUserIDAndVariantID(c, userID, variantID)
}

func (s *Service) Service_() {
	s.repo.Repository_()
}
