package topic_srvc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
)

type Repository interface {
	CreateTopic(c context.Context, ent *entity.Topic) error
	TopicByID(c context.Context, id int) (*entity.Topic, error)
	UpdateTopic(c context.Context, id int, update *Update) error
	ListTopic(c context.Context, size, page int) ([]entity.Topic, error)
	DeleteTopic(c context.Context, id, userID int) error
	Repository_()
}
