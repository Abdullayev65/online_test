package topic_srvc

import "github.com/Abdullayev65/online_test/internal/entity"

type Create struct {
	Name string `json:"name"`
}
type Update struct {
	Name *string `json:"name"`
}
type TopicDetail struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

func NewDetail(ent *entity.Topic) *TopicDetail {
	return &TopicDetail{ID: &ent.ID, Name: &ent.Name}
}
