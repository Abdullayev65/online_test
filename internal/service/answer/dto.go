package answer_srvc

import "github.com/Abdullayev65/online_test/internal/entity"

type List struct {
	ID          int    `json:"id"`
	Text        string `json:"text,omitempty"`
	Description string `json:"description,omitempty"`
	IsCorrect   bool   `json:"isCorrect"`
	Chosen      int    `json:"chosen,omitempty"`
}
type Create struct {
	Text        *string `json:"text,omitempty"`
	Description *string `json:"description,omitempty"`
	IsCorrect   *bool   `json:"isCorrect"`
}
type Update struct {
	ID          int     `json:"id"`
	Text        *string `json:"text,omitempty"`
	Description *string `json:"description,omitempty"`
	IsCorrect   *bool   `json:"isCorrect"`
}

func NewList(a *entity.Answer) *List {
	return &List{ID: a.ID, Text: a.Text, Description: a.Description,
		IsCorrect: a.IsCorrect, Chosen: a.Chosen}
}
