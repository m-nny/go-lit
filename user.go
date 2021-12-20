package lit

import (
	"context"
	"time"
)

// User represents a user in the system.
type User struct {
	ID int `json:"id"`

	Name  string `json:"name"`
	Email string `json:"email"`

	// Timestamps for user creation & last update.
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Roles []string `json:"roles"`
}

// Validate returns an error if the user contains invalid fields.
// This only performs basic validation.
func (u *User) Validate() error {
	if u.Name == "" {
		return ErrorF(EINVALID, "User name required.")
	}
	return nil
}

// UserService represents a service for managing users.
type UserService interface {
	// Retrieves a user by ID along with their associated auth objects.
	// Returns ENOTFOUND if user does not exist.
	FindUserById(ctx context.Context, id int) (*User, error)

	// Retrieves a list of users by filter. Also returns total count of matching
	// users which may differ from returned results if filter.Limit is specified.
	FindUsers(ctx context.Context, filter UserFilter) ([]*User, int, error)

	// Creates a new user. This is only used for testing since users are typically
	// created during the OAuth creation process in AuthService.CreateAuth().
	CreateUser(ctx context.Context, user *User) error

	// Updates a user object. Returns EUNAUTHORIZED if current user is not
	// the user that is being updated. Returns ENOTFOUND if user does not exist.
	UpdateUser(ctx context.Context, id int, upd UserUpdate) (*User, error)

	// Permanently deletes a user and all owned submissions, problebs. Returns EUNAUTHORIZED
	// if current user is not the user admin. Returns ENOTFOUND if user does not exist.
	Deleteuser(ctx context.Context, id int) error
}

// UserFilter represents a filter passed to FindUsers().
type UserFilter struct {
	// Filtering fields.
	ID    *int    `json:"id"`
	Email *string `json:"email"`

	// Restrict to subset of results.
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

// UserUpdate represents a set of fields to be updated via UpdateUser().
type UserUpdate struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
}
