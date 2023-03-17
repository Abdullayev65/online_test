package topic_uc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
	topic_srvc "github.com/Abdullayev65/online_test/internal/service/topic"
)

type Topic interface {
	CreateTopic(c context.Context, ent *entity.Topic) error
	TopicByID(c context.Context, id int) (*entity.Topic, error)
	UpdateTopic(c context.Context, id int, update *topic_srvc.Update) error
	ListTopic(c context.Context, size, page int) ([]entity.Topic, error)
	DeleteTopic(c context.Context, id, userID int) error
	Service_()
}
