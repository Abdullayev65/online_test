package variant_uc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
	answer_srvc "github.com/Abdullayev65/online_test/internal/service/answer"
	question_srvc "github.com/Abdullayev65/online_test/internal/service/question"
	variant_srvc "github.com/Abdullayev65/online_test/internal/service/variant"
	variant_question_srvc "github.com/Abdullayev65/online_test/internal/service/variant_question"
	"github.com/Abdullayev65/online_test/internal/utill"
	"strconv"
	"time"
)

type UseCase struct {
	Variant         Variant
	VariantQuestion VariantQuestion
	Question        Question
	Answer          Answer
}

func NewUseCase(variant Variant, variantQuestion VariantQuestion,
	question Question, answer Answer) *UseCase {

	return &UseCase{Variant: variant, VariantQuestion: variantQuestion,
		Question: question, Answer: answer}
}

func (u *UseCase) GenerateVariant(c context.Context, data *variant_srvc.Create) (*variant_srvc.Detail, error) {

	if data.Name == nil {
		unix := time.Now().Unix()
		defaultName := "default_name_" + strconv.Itoa(int(unix))
		data.Name = &defaultName
	}

	variant, err := u.Variant.Create(c, data)
	if err != nil {
		return nil, err
	}

	detail := variant_srvc.NewDetail(variant)
	questionIDs := make([]int, 0)

	for _, createQuestion := range data.Questions {
		cq := createQuestion
		cq.VariantID = variant.ID
		variantQuestions, err := u.VariantQuestion.Create(c, &cq)
		if err != nil {
			return detail, err
		}
		ids := utill.Map(variantQuestions, func(v entity.VariantQuestion) int {
			return *v.QuestionID
		})
		questionIDs = append(questionIDs, ids...)
	}
	qs, _, err := u.Question.GetAll(c, &question_srvc.Filter{IDs: &questionIDs})
	if err != nil {
		return detail, err
	}

	detail.AppendQuestions(qs...)
	return detail, nil
}

func (u *UseCase) GetVariantDetailByID(c context.Context, id int) (*variant_srvc.Detail, error) {
	variant, err := u.Variant.GetByID(c, id)
	if err != nil {
		return nil, err
	}

	detail := variant_srvc.NewDetail(variant)
	variantQuestions, _, err := u.VariantQuestion.GetAll(c, &variant_question_srvc.Filter{VariantID: &id})
	if err != nil {
		return detail, err
	}

	questionIDs := utill.Map(variantQuestions, func(i entity.VariantQuestion) int {
		return *i.QuestionID
	})
	questions, _, err := u.Question.GetAll(c, &question_srvc.Filter{IDs: &questionIDs})
	if err != nil {
		return detail, err
	}

	questionDetails := make([]question_srvc.Detail, 0)
	for _, question := range questions {
		answers, _, err := u.Answer.GetAll(c, &answer_srvc.Filter{QuestionID: &question.ID})
		if err != nil {
			return nil, err
		}

		questionDetail := question_srvc.NewDetail(&question)
		questionDetail.AppendAnswers(answers)
		questionDetails = append(questionDetails, *questionDetail)
	}

	detail.Questions = questionDetails

	return detail, nil
}

func (u *UseCase) DeleteVariant(c context.Context, variantID, userID int) error {
	return u.Variant.Delete(c, variantID, userID)
}

func (u *UseCase) ListVariant(c context.Context, size, page int) ([]variant_srvc.List, error) {
	variantList, _, err := u.Variant.GetAll(c, &variant_srvc.Filter{Limit: &size, Offset: &page})
	if err != nil {
		return nil, err
	}

	lists := utill.Map(variantList, func(i entity.Variant) variant_srvc.List {
		return variant_srvc.NewList(&i)
	})

	return lists, nil
}
