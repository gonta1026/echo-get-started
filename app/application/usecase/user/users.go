package user

import (
	"echo-get-started/app/application/usecase/user/model/response"
	"echo-get-started/app/domain/model/user"
)

type UsersAccessor interface {
	Handle() (*[]response.User, error)
}

type Users struct {
	userRepository user.Repository
}

func NewUsers(user user.Repository) *Users {
	return &Users{
		userRepository: user,
	}
}

func (u Users) Handle() (*[]response.User, error) {
	users, err := u.userRepository.List()
	res := response.NewUsersForList(users)
	return &res, err
}
