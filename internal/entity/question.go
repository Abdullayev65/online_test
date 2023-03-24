package entity

import "github.com/uptrace/bun"

type Question struct {
	bun.BaseModel `bun:"table:question"`
	BasicEntityID
	Text        *string `bun:"text,nullzero,notnull"`
	Description *string `bun:"description,nullzero"`
	ImagePath   *string `bun:"image_path"`
	TopicID     *int    `bun:"topic_id,nullzero"`
	Chosen      *int    `bun:"chosen,notnull"`
}
