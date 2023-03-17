package question_uc

import (
	"context"
	"errors"
	question_srvc "github.com/Abdullayev65/online_test/internal/service/question"
)

type UseCase struct {
	Question Question
	Answer   Answer
}

func NewUseCase(question Question, answer Answer) *UseCase {
	return &UseCase{Question: question, Answer: answer}
}

func (u *UseCase) Create(c context.Context, create *question_srvc.Create,
	userID int) (*question_srvc.Detail, error) {
	if create.Text == nil {
		return nil, errors.New("text can not be null")
	}

	question, err := u.Question.CreateQuestion(c, create, userID)
	if err != nil {
		return nil, err
	}

	detail := question_srvc.NewDetail(question)

	for _, answerCreate := range create.Answers {
		answer, err := u.Answer.CreateAnswer(c, &answerCreate, userID, question.ID)
		if err != nil {
			return detail, err
		}
		detail.AppendAnswer(answer)
	}

	return detail, nil
}
func (u *UseCase) GetQuestionDetailByID(c context.Context, id int) (*question_srvc.Detail, error) {
	question, err := u.Question.QuestionByID(c, id)
	if err != nil {
		return nil, err
	}
	detail := question_srvc.NewDetail(question)
	answers, err := u.Answer.AnswersByQuestionID(c, id)
	if err != nil {
		return detail, err
	}
	detail.AppendAnswers(answers)
	return detail, nil
}
func (u *UseCase) UpdateQuestion(c context.Context, id int, data *question_srvc.Update, userID int) error {
	err := u.Question.UpdateQuestion(c, id, data)
	if err != nil {
		return err
	}
	if len(data.AnswersCreate) > 0 {
		for _, create := range data.AnswersCreate {
			_, err = u.Answer.CreateAnswer(c, &create, userID, id)
			if err != nil {
				return err
			}
		}
	}
	if len(data.AnswersUpdate) > 0 {
		for _, update := range data.AnswersUpdate {
			err = u.Answer.UpdateAnswer(c, &update)
			if err != nil {
				return err
			}
		}
	}
	if len(data.AnswerIDsDelete) > 0 {
		for _, answerID := range data.AnswerIDsDelete {
			err = u.Answer.DeleteAnswer(c, answerID, userID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (u *UseCase) List(c context.Context, size int, page int) ([]question_srvc.Detail, error) {
	list, err := u.Question.ListQuestion(c, size, page)
	if err != nil {
		return nil, err
	}
	dtos := make([]question_srvc.Detail, 0)
	for _, question := range list {
		t := question
		dtos = append(dtos, *question_srvc.NewDetail(&t))
	}
	return dtos, nil
}
func (u *UseCase) DeleteQuestion(c context.Context, id int, userID int) error {
	return u.Question.DeleteQuestion(c, id, userID)
}
