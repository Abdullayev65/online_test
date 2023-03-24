package topic_srvc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
)

type Repository interface {
	GetAll(ctx context.Context, filter *Filter) ([]entity.Topic, int, error)
	GetByID(c context.Context, id int) (*entity.Topic, error)
	Create(c context.Context, ent *entity.Topic) error
	Update(c context.Context, id int, update *Update) error
	Delete(c context.Context, id, userID int) error
	Repository_()
}
