package echo

import "github.com/m-nny/go-lit/internal/lit"

type CreateUserArg struct {
	ID uint `json:"id"`

	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (u *CreateUserArg) User() (*lit.User, error) {
	hashedPassword, err := lit.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	return &lit.User{
		ID:             u.ID,
		Name:           u.Name,
		HashedPassword: hashedPassword,
	}, nil
}

type UpdateUserArg struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (u *UpdateUserArg) User() (*lit.UserUpdate, error) {
	hashedPassword, err := lit.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	return &lit.UserUpdate{
		Name:           u.Name,
		HashedPassword: hashedPassword,
	}, nil
}

type UserIdArg struct {
	UserId uint `json:"userId" param:"userId"`
}
