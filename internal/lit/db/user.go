package db

import (
	"context"
	"time"

	"github.com/m-nny/go-lit/internal/lit"
)

type UserModel struct {
	ID    uint `gorm:"primarykey"`
	Name  string
	Email string
	// Roles []string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func newUserModel(user *lit.User) *UserModel {
	return &UserModel{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		// Roles:     user.Roles,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// Ensure service implements interface.
var _ lit.UserService = (*UserService)(nil)

type UserService struct {
	db *DB
}

func NewUserService(db *DB) *UserService {
	return &UserService{db: db}
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(ctx context.Context, user *lit.User) error {
	userModel := newUserModel(user)
	result := s.db.gDb.Create(userModel)
	return result.Error
}

// DeleteUser permanently deletes a user and all owned dials.
// Returns EUNAUTHORIZED if current user is not the user being deleted.
// Returns ENOTFOUND if user does not exist.
func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	result := s.db.gDb.Delete(&UserModel{}, id)
	return result.Error
}

// FindUserByID retrieves a user by ID along with their associated auth objects.
// Returns ENOTFOUND if user does not exist.
func (s *UserService) FindUserById(ctx context.Context, id uint) (*lit.User, error) {
	return nil, lit.Errorf(lit.ENOTIMPLEMENTED, "method not implemented yet")
}

// FindUsers retrieves a list of users by filter. Also returns total count of
// matching users which may differ from returned results if filter.Limit is specified.
func (s *UserService) FindUsers(ctx context.Context, filter lit.UserFilter) ([]*lit.User, int, error) {
	return nil, 0, lit.Errorf(lit.ENOTIMPLEMENTED, "method not implemented yet")
}

// UpdateUser updates a user object. Returns EUNAUTHORIZED if current user is
// not the user that is being updated. Returns ENOTFOUND if user does not exist.
func (s *UserService) UpdateUser(ctx context.Context, id uint, upd lit.UserUpdate) (*lit.User, error) {
	return nil, lit.Errorf(lit.ENOTIMPLEMENTED, "method not implemented yet")
}
