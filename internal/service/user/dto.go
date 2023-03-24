package user

import (
	"github.com/Abdullayev65/online_test/internal/entity"
)

type List struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Type     int    `json:"type,omitempty"`
}
type Create struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
	Type     *int    `json:"type"`
}
type SignIn struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
type Update struct {
	ID       *int    `json:"id"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}
type Detail struct {
	ID       *int    `json:"id"`
	Username *string `json:"username"`
	Type     *int    `json:"type"`
}
type Filter struct {
	Limit          *int
	Offset         *int
	Order          *string
	AllWithDeleted bool
	OnlyDeleted    bool
}

func NewUserDetail(u *entity.User) *Detail {
	return &Detail{ID: &u.ID, Username: u.Username, Type: u.Type}
}
