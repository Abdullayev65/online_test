package entity

import "time"

type BasicEntity struct {
	CreatedAt time.Time `bun:"created_at,notnull,nullzero,default:now()"`
	CreatedBy *int      `bun:"created_by,notnull,nullzero"`
	DeletedAt time.Time `bun:"deleted_at,soft_delete,nullzero"`
	DeletedBy *int      `bun:"deleted_by,nullzero"`
}
type BasicEntityID struct {
	BasicEntity
	ID int `bun:",pk,autoincrement"`
}
