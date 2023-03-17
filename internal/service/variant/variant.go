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

func (s *Service) CreateVariant(c context.Context, data *Create, userID int) (*entity.Variant, error) {
	return s.repo.CreateVariant(c, data, userID)
}

func (s *Service) VariantByID(c context.Context, id int) (*entity.Variant, error) {
	return s.repo.VariantByID(c, id)
}

func (s *Service) ListVariant(c context.Context, size, page int) ([]entity.Variant, error) {
	return s.repo.ListVariant(c, size, page)
}

func (s *Service) DeleteVariant(c context.Context, id, userID int) error {
	return s.repo.DeleteVariant(c, id, userID)
}

func (s *Service) Service_() {
	s.repo.Repository_()
}
