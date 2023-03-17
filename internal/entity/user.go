package entity

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            int    `bun:",pk,autoincrement"`
	Username      string `bun:",unique,notnull,nullzero"`
	Password      string `bun:",notnull,nullzero"`
	Type          int    `bun:",nullzero,notnull"`
}
