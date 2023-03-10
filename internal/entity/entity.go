package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type BasicEntity struct {
	CreatedAt time.Time `bun:",notnull,nullzero,default:now()"`
	CreatedBy int       `bun:",notnull,nullzero"`
	DeletedAt time.Time `bun:",soft_delete,nullzero"`
	DeletedBy int       `bun:",nullzero"`
}
type BasicEntityID struct {
	BasicEntity
	ID int `bun:",pk,autoincrement"`
}

/*	Database entities	*/
type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            int    `bun:",pk,autoincrement"`
	Username      string `bun:",unique,notnull,nullzero"`
	Password      string `bun:",notnull,nullzero"`
	Type          int    `bun:",nullzero,notnull"`
}
type Topic struct {
	bun.BaseModel `bun:"table:topic"`
	BasicEntityID
	Name string `bun:",nullzero,notnull"`
}
type Question struct {
	bun.BaseModel `bun:"table:question"`
	BasicEntityID
	Text        string `bun:"nullzero,notnull"`
	Description string `bun:"nullzero"`
	TopicID     int    `bun:"nullzero"`
	Chosen      int    `bun:"nullzero,notnull"`
}
type Answer struct {
	bun.BaseModel `bun:"table:answer"`
	BasicEntityID
	Text        string `bun:",nullzero,notnull"`
	Description string `bun:",nullzero"`
	IsCorrect   bool   `bun:"notnull"`
	QuestionID  int    `bun:",nullzero,notnull"`
	Chosen      int    `bun:",nullzero,notnull"`
}
