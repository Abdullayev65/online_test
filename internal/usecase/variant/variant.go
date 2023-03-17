package variant_uc

import (
	"context"
	"github.com/Abdullayev65/online_test/internal/entity"
	variant_srvc "github.com/Abdullayev65/online_test/internal/service/variant"
	"github.com/Abdullayev65/online_test/internal/utill"
	"strconv"
	"time"
)

type UseCase struct {
	Variant         Variant
	VariantQuestion VariantQuestion
	Question        Question
}

func NewUseCase(variant Variant, variantQuestion VariantQuestion, question Question) *UseCase {
	return &UseCase{Variant: variant,
		VariantQuestion: variantQuestion, Question: question}
}

func (u *UseCase) GenerateVariant(c context.Context, data *variant_srvc.Create,
	userID int) (*variant_srvc.Detail, error) {

	if data.Name == nil {
		unix := time.Now().Unix()
		defaultName := "default_name_" + strconv.Itoa(int(unix))
		data.Name = &defaultName
	}

	variant, err := u.Variant.CreateVariant(c, data, userID)
	if err != nil {
		return nil, err
	}

	detail := variant_srvc.NewDetail(variant)
	questionIDs := make([]int, 0)

	for _, createQuestion := range data.Questions {
		cq := createQuestion
		variantQuestions, err := u.VariantQuestion.CreateVariantQuestion(c, &cq, variant.ID)
		if err != nil {
			return detail, err
		}
		ids := utill.Map(variantQuestions, func(v entity.VariantQuestion) int {
			return v.QuestionID
		})
		questionIDs = append(questionIDs, ids...)
	}
	qs, err := u.Question.QuestionByIDs(c, questionIDs)
	if err != nil {
		return detail, err
	}

	detail.AppendQuestions(qs...)
	return detail, nil
}

func (u *UseCase) GetVariantDetailByID(c context.Context, id int) (*variant_srvc.Detail, error) {
	variant, err := u.Variant.VariantByID(c, id)
	if err != nil {
		return nil, err
	}

	detail := variant_srvc.NewDetail(variant)
	variantQuestions, err := u.VariantQuestion.VariantQuestionByVariantID(c, id)
	if err != nil {
		return detail, err
	}

	questionIDs := utill.Map(variantQuestions, func(i entity.VariantQuestion) int {
		return i.QuestionID
	})
	questions, err := u.Question.QuestionByIDs(c, questionIDs)
	if err != nil {
		return detail, err
	}

	detail.AppendQuestions(questions...)

	return detail, nil
}

func (u *UseCase) DeleteVariant(c context.Context, variantID, userID int) error {
	return u.Variant.DeleteVariant(c, variantID, userID)
}

func (u *UseCase) ListVariant(c context.Context, size, page int) ([]variant_srvc.List, error) {
	variantList, err := u.Variant.ListVariant(c, size, page)
	if err != nil {
		return nil, err
	}

	lists := utill.Map(variantList, func(i entity.Variant) variant_srvc.List {
		return variant_srvc.NewList(&i)
	})

	return lists, nil
}
