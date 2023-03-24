package answer_srvc

import (
	"github.com/Abdullayev65/online_test/internal/entity"
	utill "github.com/Abdullayev65/online_test/internal/utill"
)

type List struct {
	ID          *int    `json:"id"`
	Text        *string `json:"text"`
	Description *string `json:"description,omitempty"`
	IsCorrect   *bool   `json:"is_correct,omitempty"`
}
type Create struct {
	Text        *string `json:"text,omitempty"`
	Description *string `json:"description,omitempty"`
	IsCorrect   *bool   `json:"is_correct"`
	UserID      *int    `json:"-"`
	QuestionID  *int    `json:"-"`
}
type Update struct {
	ID          int     `json:"id"`
	Text        *string `json:"text,omitempty"`
	Description *string `json:"description,omitempty"`
	IsCorrect   *bool   `json:"is_correct"`
}
type Filter struct {
	Limit          *int
	Offset         *int
	Order          *string
	CreatedBy      *int
	QuestionID     *int
	AllWithDeleted bool
	OnlyDeleted    bool
}

func NewList(a *entity.Answer) List {
	return List{ID: &a.ID, Text: a.Text, Description: a.Description,
		IsCorrect: a.IsCorrect}
}
func NewLists(ms ...entity.Answer) []List {
	lists := utill.Map(ms, func(i entity.Answer) List {
		return NewList(&i)
	})
	return lists
}
