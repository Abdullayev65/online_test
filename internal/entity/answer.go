package entity

import "github.com/uptrace/bun"

type Answer struct {
	bun.BaseModel `bun:"table:answer"`
	BasicEntityID
	Text        *string `bun:"text,nullzero"`
	Description *string `bun:"description,nullzero"`
	IsCorrect   *bool   `bun:"is_correct,notnull"`
	QuestionID  *int    `bun:"question_id,nullzero,notnull"`
	Chosen      *int    `bun:"chosen,notnull"`
}
