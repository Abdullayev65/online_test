package variant_question_answer_uc

import (
	"context"
	"errors"
	answer_srvc "github.com/Abdullayev65/online_test/internal/service/answer"
	srvc "github.com/Abdullayev65/online_test/internal/service/variant_question_answer"
)

type UseCase struct {
	Svc      VariantQuestionAnswer
	Question Question
	Answer   Answer
	Variant  Variant
}

func NewUseCase(svc VariantQuestionAnswer, question Question,
	answer Answer, variant Variant) *UseCase {
	return &UseCase{Svc: svc, Answer: answer, Variant: variant, Question: question}
}

func (u *UseCase) Create(c context.Context, create *srvc.Create,
	userID int) (*srvc.QuestionAnswerDTO, error) {

	if create.AnswerID == nil {
		return nil, errors.New("answer_id can not be null")
	}
	if create.QuestionID == nil {
		return nil, errors.New("question_id can not be null")
	}
	if create.VariantID == nil {
		return nil, errors.New("variant_id can not be null")
	}

	questionAnswer, err := u.Svc.Create(c, create, userID)
	if err != nil {
		return nil, err
	}

	dto := &srvc.QuestionAnswerDTO{IsCorrect: questionAnswer.IsCorrect,
		CorrectAnswerID: questionAnswer.AnswerID}

	if !dto.IsCorrect {
		correctAnswer, err := u.Answer.CorrectAnswerByQuestionID(c, *create.QuestionID)
		if err != nil {
			return dto, err
		}
		dto.CorrectAnswerID = correctAnswer.ID
	}

	return dto, nil
}

func (u *UseCase) VariantAnswerByUserIDAndVariantID(c context.Context, userID,
	variantID int) (*srvc.UserVariantAnswer, error) {

	variant, err := u.Variant.VariantByID(c, variantID)
	if err != nil {
		return nil, err
	}

	variantQuestionAnswers, err := u.Svc.GetByUserIDAndVariantID(c, userID, variantID)
	if err != nil {
		return nil, err
	}

	questionAnswerList := make([]srvc.QuestionAnswerList, 0)

	for _, variantQuestionAnswer := range variantQuestionAnswers {
		question, err := u.Question.QuestionByID(c, variantQuestionAnswer.QuestionID)
		if err != nil {
			return nil, err
		}

		answers, err := u.Answer.AnswersByQuestionID(c, variantQuestionAnswer.QuestionID)
		if err != nil {
			return nil, err
		}

		answersList := answer_srvc.NewLists(answers...)
		answerList := srvc.NewQuestionAnswerList(&variantQuestionAnswer, question, answersList)

		questionAnswerList = append(questionAnswerList, *answerList)
	}

	dto := srvc.UserVariantAnswer{Name: variant.Name, Questions: questionAnswerList}

	return &dto, nil
}
