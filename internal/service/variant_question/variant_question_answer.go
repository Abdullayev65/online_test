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

func (s *Service) Create(c context.Context, data *Create) ([]entity.VariantQuestion, error) {
	return s.repo.Create(c, data)
}

func (s *Service) GetByID(c context.Context, id int) (*entity.VariantQuestion, error) {
	return s.repo.GetById(c, id)
}

func (s *Service) GetAll(c context.Context, filter *Filter) ([]entity.VariantQuestion, int, error) {
	return s.repo.GetAll(c, filter)
}

func (s *Service) Delete(c context.Context, id int) error {
	return s.repo.Delete(c, id)
}

func (s *Service) Service_() {
	s.repo.Repository_()
}
