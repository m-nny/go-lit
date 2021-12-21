package gorm

import (
	"context"
	"errors"
	"time"

	"github.com/m-nny/go-lit/internal/lit"
	"gorm.io/gorm"
)

type UserModel struct {
	ID             uint `gorm:"primarykey"`
	Name           string
	HashedPassword string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func newUserModel(user *lit.User) *UserModel {
	return &UserModel{
		ID:             user.ID,
		Name:           user.Name,
		HashedPassword: user.HashedPassword,

		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// Ensure service implements interface.
var _ lit.UserService = (*UserService)(nil)

type UserService struct {
	gormDb *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{gormDb: db}
}
func (userModel *UserModel) User() *lit.User {
	return &lit.User{
		ID:             userModel.ID,
		Name:           userModel.Name,
		HashedPassword: userModel.HashedPassword,

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
func (s *UserService) CreateUser(ctx context.Context, user *lit.User) (*lit.User, error) {
	userModel := newUserModel(user)
	result := s.gormDb.Create(userModel)
	return userModel.User(), result.Error
}

// DeleteUser permanently deletes a user and all owned dials.
// Returns ENOTFOUND if user does not exist.
func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	result := s.gormDb.Delete(&UserModel{}, id)
	return result.Error
}

// FindUserByID retrieves a user by ID along with their associated auth objects.
// Returns ENOTFOUND if user does not exist.
func (s *UserService) FindUserById(ctx context.Context, id uint) (*lit.User, error) {
	var userModel UserModel
	result := s.gormDb.First(&userModel, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, lit.Errorf(lit.ENOTFOUND, `User with id "%d" not found`, id)
	}
	return userModel.User(), result.Error
}

// FindUsers retrieves a list of users by filter. Also returns total count of
// matching users which may differ from returned results if filter.Limit is specified.
func (s *UserService) FindUsers(ctx context.Context, filter *lit.UserFilter) ([]*lit.User, int, error) {
	userModels := []UserModel{}
	// var userModels []UserModel
	var count int64
	result := s.gormDb.WithContext(ctx).Where(filter).Find(&userModels).Count(&count)
	return Users(userModels), int(count), result.Error
}

// UpdateUser updates a user object. Returns EUNAUTHORIZED if current user is
// not the user that is being updated. Returns ENOTFOUND if user does not exist.
func (s *UserService) UpdateUser(ctx context.Context, id uint, upd *lit.UserUpdate) (*lit.User, error) {
	var userModel UserModel
	result := s.gormDb.First(&userModel, id)
	if result.Error != nil {
		return nil, result.Error
	}
	result = s.gormDb.Model(&userModel).Updates(upd)
	return userModel.User(), result.Error
}
