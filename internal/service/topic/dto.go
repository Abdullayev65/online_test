package topic_service

import "github.com/Abdullayev65/online_test/internal/entity"

type Create struct {
	Name string `json:"name"`
}
type Update struct {
	Name *string `json:"name"`
}
type DTO struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

func NewDTO(ent *entity.Topic) *DTO {
	return &DTO{ID: &ent.ID, Name: &ent.Name}
}
