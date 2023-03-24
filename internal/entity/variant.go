package entity

import "github.com/uptrace/bun"

type Variant struct {
	bun.BaseModel `bun:"table:variant"`
	BasicEntityID
	Name   *string `bun:"name"`
	Number *int    `bun:"number"`
}
