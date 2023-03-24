package entity

import "github.com/uptrace/bun"

type Topic struct {
	bun.BaseModel `bun:"table:topic"`
	BasicEntityID
	Name *string `bun:"name,notnull"`
}
