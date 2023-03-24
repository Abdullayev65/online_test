package entity

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            int     `bun:"id,pk,autoincrement"`
	Username      *string `bun:"username,unique,notnull,nullzero"`
	Password      *string `bun:"password,notnull,nullzero"`
	Type          *int    `bun:"type,notnull"`
}
