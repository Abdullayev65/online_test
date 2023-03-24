package variant_srvc

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

func (s *Service) GetAll(c context.Context, filter *Filter) ([]entity.Variant, int, error) {
	return s.repo.GetAll(c, filter)
}

func (s *Service) GetByID(c context.Context, id int) (*entity.Variant, error) {
	return s.repo.GetByID(c, id)
}

func (s *Service) Create(c context.Context, data *Create) (*entity.Variant, error) {
	return s.repo.Create(c, data)
}

func (s *Service) Delete(c context.Context, id, userID int) error {
	return s.repo.Delete(c, id, userID)
}

func (s *Service) Service_() {
	s.repo.Repository_()
}
