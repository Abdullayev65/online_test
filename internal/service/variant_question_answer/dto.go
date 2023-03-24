package variant_question_answer_srvc

import (
	"github.com/Abdullayev65/online_test/internal/entity"
	answer_srvc "github.com/Abdullayev65/online_test/internal/service/answer"
)

type Create struct {
	VariantID  *int `json:"variant_id"`
	QuestionID *int `json:"question_id"`
	AnswerID   *int `json:"answer_id"`
	UserID     *int `json:"-"`
}

type Filter struct {
	Limit          *int
	Offset         *int
	Order          *string
	AllWithDeleted bool
	OnlyDeleted    bool
	UserID         *int
	VariantID      *int
}

type QuestionAnswerDTO struct {
	IsCorrect       *bool `json:"is_correct"`
	CorrectAnswerID *int  `json:"correct_answer_id"`
}

type UserVariantAnswer struct {
	Name      *string `json:"name"`
	Questions []QuestionAnswerList
}

type QuestionAnswerList struct {
	IsCorrect      *bool              `json:"is_correct"`
	ChosenAnswerID *int               `json:"chosen_answer_id"`
	Answers        []answer_srvc.List `json:"answers"`
	ID             *int               `json:"id"`
	Text           *string            `json:"text"`
	Description    *string            `json:"description"`
	TopicID        *int               `json:"topic_id,omitempty"`
}

func NewQuestionAnswerList(vqa *entity.VariantQuestionAnswer, question *entity.Question,
	answers []answer_srvc.List) *QuestionAnswerList {

	return &QuestionAnswerList{
		ID:             &question.ID,
		IsCorrect:      vqa.IsCorrect,
		ChosenAnswerID: vqa.AnswerID,
		Answers:        answers,
		Text:           question.Text,
		Description:    question.Description,
		TopicID:        question.TopicID,
	}
}
