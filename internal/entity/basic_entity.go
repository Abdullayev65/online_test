package entity

import "time"

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
