package user

import (
	"context"
	"errors"
	"github.com/Abdullayev65/online_test/internal/entity"
	user_service "github.com/Abdullayev65/online_test/internal/service/user"
	"github.com/Abdullayev65/online_test/internal/utill"
)

type UseCase struct {
	User     User
	TokenJWT *utill.TokenJWT
}

func NewUseCase(svc User, tokenJWT *utill.TokenJWT) *UseCase {
	return &UseCase{User: svc, TokenJWT: tokenJWT}
}

func (u *UseCase) GetUserDetail(c context.Context, userID int) (*user_service.Detail, error) {
	m, err := u.User.GetByID(c, userID)
	if err != nil {
		return nil, err
	}

	detail := user_service.NewUserDetail(m)
	return detail, nil
}
func (u *UseCase) CreateUser(c context.Context, cu *user_service.Create) (*entity.User, error) {
	if cu.Username == nil {
		return nil, errors.New("username is invalid")
	}
	if cu.Password == nil {
		return nil, errors.New("password is invalid")
	}
	if *cu.Type == 0 {
		return nil, errors.New("type is invalid")
	}
	if *cu.Type != 2 && *cu.Type != 3 {
		return nil, errors.New("type can be 2 or 3")
	}
	user := entity.User{Username: cu.Username, Password: cu.Password, Type: cu.Type}
	err := u.User.Create(c, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (u *UseCase) UpdateUser(c context.Context, userID int, update *user_service.Update) error {
	if !utill.AnyNotNil(update.Username, update.Password) {
		return errors.New("nothing to update")
	}
	return u.User.Update(c, userID, update)
}

func (u *UseCase) GetUserByUsername(c context.Context, username string) (*entity.User, error) {
	return u.User.GetByUsername(c, username)
}
func (u *UseCase) UserGenerateToken(c context.Context, sign *user_service.SignIn) (string, error) {
	if sign.Username == nil || sign.Password == nil {
		return "", errors.New("username or password invalid")
	}
	user, err := u.GetUserByUsername(c, *sign.Username)
	if err != nil || *user.Password != *sign.Password {
		return "", errors.New("username or password wrong")
	}
	token, err := u.TokenJWT.Generate(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
