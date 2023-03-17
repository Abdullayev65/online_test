package entity

import "github.com/uptrace/bun"

type Answer struct {
	bun.BaseModel `bun:"table:answer"`
	BasicEntityID
	Text        string `bun:",nullzero"`
	Description string `bun:",nullzero"`
	IsCorrect   bool   `bun:",notnull"`
	QuestionID  int    `bun:",nullzero,notnull"`
	Chosen      int    `bun:",notnull"`
}
