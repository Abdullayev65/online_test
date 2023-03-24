package entity

import "github.com/uptrace/bun"

type VariantQuestion struct {
	bun.BaseModel `bun:"variant_question"`
	ID            int  `bun:"id,pk,autoincrement"`
	VariantID     *int `bun:"variant_id"`
	QuestionID    *int `bun:"question_id"`
}
