package topic_repo

import (
	"context"
	"errors"
	"github.com/Abdullayev65/online_test/internal/entity"
	topic_service "github.com/Abdullayev65/online_test/internal/service/topic"
	"github.com/uptrace/bun"
	"time"
)

type Repository struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *Repository {
	return &Repository{DB: DB}
}

func (r *Repository) Repository_() {
}
func (r *Repository) TopicByID(c context.Context, id int) (*entity.Topic, error) {
	ent := new(entity.Topic)
	ent.ID = id
	err := r.DB.NewSelect().Model(ent).WherePK().Scan(c)
	if err != nil {
		return nil, err
	}
	return ent, nil
}
func (r *Repository) UpdateTopic(c context.Context, id int, update *topic_service.Update) error {
	ent, err := r.TopicByID(c, id)
	if err != nil {
		return errors.New("topic not found")
	}
	if update.Name != nil {
		ent.Name = *update.Name
	}
	_, err = r.DB.NewUpdate().Model(ent).WherePK().Exec(c)
	return err
}
func (r *Repository) CreateTopic(c context.Context, ent *entity.Topic) error {
	_, err := r.DB.NewInsert().Model(ent).Exec(c)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) ListTopic(c context.Context, size, page int) ([]entity.Topic, error) {
	slice := make([]entity.Topic, 0)
	err := r.DB.NewSelect().Model(&slice).
		Offset(size * (page - 1)).Limit(size).
		Order("id ASC").Scan(c)
	if err != nil {
		return nil, err
	}
	return slice, nil
}
func (r *Repository) DeleteTopic(c context.Context, id int, userID int) error {
	topic, err := r.TopicByID(c, id)
	if err != nil {
		return err
	}
	topic.DeletedBy = userID
	topic.DeletedAt = time.Now()
	_, err = r.DB.NewUpdate().Model(topic).WherePK().Exec(c)
	if err != nil {
		return err
	}
	return nil
}
