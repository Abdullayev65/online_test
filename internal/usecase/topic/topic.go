package topic_uc

import (
	"context"
	"errors"
	"github.com/Abdullayev65/online_test/internal/entity"
	topic_srvc "github.com/Abdullayev65/online_test/internal/service/topic"
)

type UseCase struct {
	Svc Topic
}

func NewUseCase(svc Topic) *UseCase {
	return &UseCase{Svc: svc}
}

func (u *UseCase) Create(c context.Context, create *topic_srvc.Create, userID int) (*topic_srvc.TopicDetail, error) {
	if create.Name == "" {
		return nil, errors.New("name can not be null or blank")
	}
	topic := &entity.Topic{Name: create.Name}
	topic.CreatedBy = userID
	err := u.Svc.CreateTopic(c, topic)
	if err != nil {
		return nil, err
	}
	dto := topic_srvc.NewDetail(topic)
	return dto, nil
}
func (u *UseCase) TopicByID(c context.Context, id int) (*topic_srvc.TopicDetail, error) {
	topic, err := u.Svc.TopicByID(c, id)
	if err != nil {
		return nil, err
	}
	dto := topic_srvc.NewDetail(topic)
	return dto, nil
}
func (u *UseCase) UpdateTopic(c context.Context, id int, update *topic_srvc.Update) error {
	if update.Name == nil {
		return errors.New("nothing to update")
	}
	err := u.Svc.UpdateTopic(c, id, update)
	if err != nil {
		return err
	}
	return nil
}
func (u *UseCase) ListTopic(c context.Context, size int, page int) ([]topic_srvc.TopicDetail, error) {
	listTopic, err := u.Svc.ListTopic(c, size, page)
	if err != nil {
		return nil, err
	}
	dtos := make([]topic_srvc.TopicDetail, 0)
	for _, topic := range listTopic {
		t := topic
		dtos = append(dtos, *topic_srvc.NewDetail(&t))
	}
	return dtos, nil
}
func (u *UseCase) DeleteTopic(c context.Context, id int, userID int) error {
	return u.Svc.DeleteTopic(c, id, userID)
}
