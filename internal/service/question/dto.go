package question_srvc

import (
	"github.com/Abdullayev65/online_test/internal/entity"
	answer_srvc "github.com/Abdullayev65/online_test/internal/service/answer"
	"mime/multipart"
)

type List struct {
	ID          *int    `json:"id"`
	Text        *string `json:"text"`
	Description *string `json:"description"`
	TopicID     *int    `json:"topic_id,omitempty"`
	Chosen      *int    `json:"chosen,omitempty"`
	ImagePath   *string `json:"image_path"`
}

type Detail struct {
	ID          *int               `json:"id"`
	Answers     []answer_srvc.List `json:"answers"`
	Text        *string            `json:"text"`
	Description *string            `json:"description"`
	TopicID     *int               `json:"topic_id"`
	Chosen      *int               `json:"chosen"`
	ImagePath   *string            `json:"image_path"`
}

type Create struct {
	Answers     []answer_srvc.Create  `json:"answers" form:"answers"`
	Text        *string               `json:"text,omitempty" form:"text"`
	Description *string               `json:"description,omitempty" form:"description"`
	TopicID     *int                  `json:"topic_id,omitempty" form:"topic_id"`
	UserId      *int                  `json:"-" form:"-"`
	ImagePath   string                `json:"-" form:"-"`
	ImageFile   *multipart.FileHeader `json:"-" form:"image_file"`
}

type Update struct {
	ID              int                  `json:"-"`
	AnswersCreate   []answer_srvc.Create `json:"answers_create"`
	AnswersUpdate   []answer_srvc.Update `json:"answers_update"`
	AnswerIDsDelete []int                `json:"answer_ids_delete"`
	Text            *string              `json:"text,omitempty"`
	Description     *string              `json:"description,omitempty"`
	TopicID         *int                 `json:"topic_id,omitempty"`
	UserID          *int                 `json:"-"`
}

type Filter struct {
	Limit          *int
	Offset         *int
	Order          *string
	CreatedBy      *int
	TopicID        *int
	IDs            []int
	AllWithDeleted bool
	OnlyDeleted    bool
}

type AnswerList struct {
	ID          int    `json:"id"`
	Text        string `json:"text,omitempty"`
	Description string `json:"description,omitempty"`
	IsCorrect   bool   `json:"isCorrect"`
	Chosen      int    `json:"chosen,omitempty"`
}

func NewDetail(q *entity.Question) *Detail {
	return &Detail{ID: &q.ID, Text: q.Text, Description: q.Description,
		TopicID: q.TopicID, Chosen: q.Chosen, ImagePath: q.ImagePath}
}

func (d *Detail) AppendAnswer(a *entity.Answer) {
	d.Answers = append(d.Answers, answer_srvc.NewList(a))
}
func (d *Detail) AppendAnswers(answers []entity.Answer) {
	for _, a := range answers {
		d.AppendAnswer(&a)
	}
}

func NewList(q *entity.Question) *List {
	return &List{ID: &q.ID, Text: q.Text, Description: q.Description, TopicID: q.TopicID,
		Chosen: q.Chosen, ImagePath: q.ImagePath}
}
