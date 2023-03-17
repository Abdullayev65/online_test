package entity

import "github.com/uptrace/bun"

type Question struct {
	bun.BaseModel `bun:"table:question"`
	BasicEntityID
	Text        string `bun:",nullzero,notnull"`
	Description string `bun:",nullzero"`
	TopicID     int    `bun:",nullzero"`
	Chosen      int    `bun:",notnull"`
}
