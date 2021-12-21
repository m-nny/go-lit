package gorm

import (
	"context"
	"time"

	"github.com/m-nny/go-lit/internal/lit"
	"gorm.io/gorm"
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
	gorm *gorm.DB
}

func NewUserService(gorm *gorm.DB) *UserService {
	return &UserService{gorm: gorm}
}
func (userModel *UserModel) User() *lit.User {
	return &lit.User{
		ID:    userModel.ID,
		Name:  userModel.Name,
		Email: userModel.Email,
		// Roles:     userModel.Roles,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
	}
}
func Users(userModels []UserModel) []*lit.User {
	users := make([]*lit.User, len(userModels))
	for i, userModel := range userModels {
		users[i] = userModel.User()
	}
	return users
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(ctx context.Context, user *lit.User) error {
	userModel := newUserModel(user)
	result := s.gorm.WithContext(ctx).Create(userModel)
	return result.Error
}

// DeleteUser permanently deletes a user and all owned dials.
// Returns ENOTFOUND if user does not exist.
func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	result := s.gorm.WithContext(ctx).Delete(&UserModel{}, id)
	return result.Error
}

// FindUserByID retrieves a user by ID along with their associated auth objects.
// Returns ENOTFOUND if user does not exist.
func (s *UserService) FindUserById(ctx context.Context, id uint) (*lit.User, error) {
	var userModel UserModel
	result := s.gorm.WithContext(ctx).First(&userModel, id)
	return userModel.User(), result.Error
}

// FindUsers retrieves a list of users by filter. Also returns total count of
// matching users which may differ from returned results if filter.Limit is specified.
func (s *UserService) FindUsers(ctx context.Context, filter lit.UserFilter) ([]*lit.User, int, error) {
	var userModels []UserModel
	var count int64
	result := s.gorm.WithContext(ctx).Where(filter).Find(&userModels).Count(&count)
	return Users(userModels), int(count), result.Error
}

// UpdateUser updates a user object. Returns EUNAUTHORIZED if current user is
// not the user that is being updated. Returns ENOTFOUND if user does not exist.
func (s *UserService) UpdateUser(ctx context.Context, id uint, upd lit.UserUpdate) (*lit.User, error) {
	var userModel UserModel
	result := s.gorm.WithContext(ctx).First(&userModel, id)
	if result.Error != nil {
		return nil, result.Error
	}
	result = s.gorm.WithContext(ctx).Model(&userModel).Updates(upd)
	return userModel.User(), result.Error
}
