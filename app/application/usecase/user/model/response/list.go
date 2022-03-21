package response

import "echo-get-started/app/domain/model/user"

func NewUsersForList(users []user.Entity) []User {
	if len(users) == 0 {
		newUsers := make([]User, 0)
		return newUsers
	}

	newUsers := make([]User, 0, len(users))
	for _, user := range users {
		newUser := User{
			ID:   user.ID,
			Name: user.Name,
		}
		newUsers = append(newUsers, newUser)
	}
	return newUsers
}
