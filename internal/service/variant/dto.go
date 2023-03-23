package variant_srvc

import (
	"github.com/Abdullayev65/online_test/internal/entity"
	question_srvc "github.com/Abdullayev65/online_test/internal/service/question"
	variant_question "github.com/Abdullayev65/online_test/internal/service/variant_question"
)

type Create struct {
	Name      *string                   `json:"name,omitempty"`
	Questions []variant_question.Create `json:"questions"`
}

type List struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Number int    `json:"number"`
}

type Detail struct {
	ID        int                    `json:"id"`
	Name      string                 `json:"name"`
	Questions []question_srvc.Detail `json:"questions"`
	Number    int                    `json:"number"`
}

func NewDetail(m *entity.Variant) *Detail {
	return &Detail{ID: m.ID, Name: m.Name, Number: m.Number}
}
func NewList(m *entity.Variant) List {
	return List{ID: m.ID, Name: m.Name, Number: m.Number}
}

func (d *Detail) AppendQuestion(a *entity.Question) {
	d.Questions = append(d.Questions, *question_srvc.NewDetail(a))
}
func (d *Detail) AppendQuestions(questions ...entity.Question) {
	for _, q := range questions {
		d.AppendQuestion(&q)
	}
}
