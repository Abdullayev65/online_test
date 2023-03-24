package entity

type VariantQuestionAnswer struct {
	BasicEntityID
	VariantID  *int  `bun:"variant_id"`
	QuestionID *int  `bun:"question_id"`
	AnswerID   *int  `bun:"answer_id"`
	IsCorrect  *bool `bun:"is_correct"`
}
