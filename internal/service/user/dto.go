package user

import "github.com/Abdullayev65/online_test/internal/entity"

type Create struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Type     int    `json:"type"`
}
type SignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Update struct {
	ID       *int    `json:"id"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}
type UserDTO struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Type     int    `json:"type"`
}

func NewUserDTO(u *entity.User) *UserDTO {
	return &UserDTO{ID: u.ID, Username: u.Username, Type: u.Type}
}
