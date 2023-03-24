package question_uc

import (
	"context"
	"errors"
	answer_srvc "github.com/Abdullayev65/online_test/internal/service/answer"
	question_srvc "github.com/Abdullayev65/online_test/internal/service/question"
)

type UseCase struct {
	Question Question
	Answer   Answer
	File     File
}

func NewUseCase(question Question, answer Answer, file File) *UseCase {
	return &UseCase{Question: question, Answer: answer, File: file}
}

func (u *UseCase) Create(c context.Context, create *question_srvc.Create) (*question_srvc.Detail, error) {
	if create.Text == nil {
		return nil, errors.New("text can not be null")
	}

	filePath, err := u.File.Upload(create.ImageFile, "images")
	if err != nil {
		return nil, err
	}
	create.ImagePath = filePath

	question, err := u.Question.Create(c, create)
	if err != nil {
		return nil, err
	}

	detail := question_srvc.NewDetail(question)

	for _, answerCreate := range create.Answers {
		answerCreate.UserID = create.UserId
		answerCreate.QuestionID = &question.ID

		answer, err := u.Answer.Create(c, &answerCreate)
		if err != nil {
			return detail, err
		}
		detail.AppendAnswer(answer)
	}

	return detail, nil
}
func (u *UseCase) GetQuestionDetailByID(c context.Context, id int) (*question_srvc.Detail, error) {
	question, err := u.Question.GetByID(c, id)
	if err != nil {
		return nil, err
	}

	detail := question_srvc.NewDetail(question)
	answers, _, err := u.Answer.GetAll(c, &answer_srvc.Filter{QuestionID: &id})
	if err != nil {
		return detail, err
	}
	detail.AppendAnswers(answers)
	return detail, nil
}
func (u *UseCase) UpdateQuestion(c context.Context, data *question_srvc.Update) error {
	err := u.Question.Update(c, data)
	if err != nil {
		return err
	}
	if len(data.AnswersCreate) > 0 {
		for _, create := range data.AnswersCreate {
			create.UserID = data.UserID
			create.QuestionID = &data.ID

			_, err = u.Answer.Create(c, &create)
			if err != nil {
				return err
			}
		}
	}
	if len(data.AnswersUpdate) > 0 {
		for _, update := range data.AnswersUpdate {
			err = u.Answer.Update(c, &update)
			if err != nil {
				return err
			}
		}
	}
	if len(data.AnswerIDsDelete) > 0 {
		for _, answerID := range data.AnswerIDsDelete {

			err = u.Answer.Delete(c, answerID, *data.UserID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (u *UseCase) List(c context.Context, size int, page int) ([]question_srvc.Detail, error) {
	list, _, err := u.Question.GetAll(c, &question_srvc.Filter{Limit: &size, Offset: &page})
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
func (u *UseCase) Delete(c context.Context, id int, userID int) error {
	return u.Question.Delete(c, id, userID)
}
