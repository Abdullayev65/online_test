package topic_uc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
	topic_srvc "github.com/Abdullayev65/online_test/internal/service/topic"
)

type Topic interface {
	Create(c context.Context, ent *entity.Topic) error
	GetByID(c context.Context, id int) (*entity.Topic, error)
	Update(c context.Context, id int, update *topic_srvc.Update) error
	GetAll(c context.Context, filter *topic_srvc.Filter) ([]entity.Topic, int, error)
	Delete(c context.Context, id, userID int) error
	Service_()
}
