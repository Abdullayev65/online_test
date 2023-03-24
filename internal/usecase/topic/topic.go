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

func (u *UseCase) ListTopic(c context.Context, offset int, limit int) ([]topic_srvc.TopicDetail, error) {
	listTopic, count, err := u.Svc.GetAll(c, &topic_srvc.Filter{Offset: &offset, Limit: &limit})
	if err != nil {
		return nil, err
	}
	_ = count

	dtos := make([]topic_srvc.TopicDetail, 0)
	for _, topic := range listTopic {
		dto := topic_srvc.TopicDetail{ID: topic.ID, Name: topic.Name}
		if dto.Name == nil {
			s := "I13"
			dto.Name = &s
		}
		dtos = append(dtos, dto) //NewDetail(&topic))
	}
	return dtos, nil
}
func (u *UseCase) TopicByID(c context.Context, id int) (*topic_srvc.TopicDetail, error) {
	topic, err := u.Svc.GetByID(c, id)
	if err != nil {
		return nil, err
	}
	dto := topic_srvc.NewDetail(topic)
	return dto, nil
}
func (u *UseCase) Create(c context.Context, create *topic_srvc.Create, userID int) (*topic_srvc.TopicDetail, error) {
	if create.Name == nil {
		return nil, errors.New("name can not be null or blank")
	}
	topic := &entity.Topic{Name: create.Name}
	topic.CreatedBy = &userID
	err := u.Svc.Create(c, topic)
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
	err := u.Svc.Update(c, id, update)
	if err != nil {
		return err
	}
	return nil
}

/* Admin */
func (u *UseCase) AdminDeleteTopic(c context.Context, id int, userID int) error {
	return u.Svc.Delete(c, id, userID)
}
