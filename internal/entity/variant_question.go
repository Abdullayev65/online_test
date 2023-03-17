package entity

import "github.com/uptrace/bun"

type VariantQuestion struct {
	bun.BaseModel `bun:"variant_question"`
	ID            int `bun:",pk,autoincrement"`
	VariantID     int
	QuestionID    int
}
