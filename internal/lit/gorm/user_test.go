package gorm

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/go-test/deep"
	"github.com/m-nny/go-lit/internal/lit"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/DATA-DOG/go-sqlmock"
)

type UserTestSuite struct {
	suite.Suite
	gorm *gorm.DB
	mock sqlmock.Sqlmock

	userService UserService
}

func (s *UserTestSuite) SetupSuite() {

	var (
		sqlDb *sql.DB
		err   error
	)

	sqlDb, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.gorm, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDb,
	}), &gorm.Config{})
	require.NoError(s.T(), err)

	s.userService = *NewUserService(s.gorm)
}

func (s *UserTestSuite) Test_UserService_FindById() {
	const (
		id   = uint(1)
		name = "test-name"

		sqlSelectById = `SELECT * FROM "user_models" WHERE "user_models"."id" = $1`
	)
	s.mock.ExpectQuery(
		regexp.QuoteMeta(sqlSelectById),
	).
		WithArgs(id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name"}).AddRow(id, name))
	ctx := context.Background()
	res, err := s.userService.FindUserById(ctx, id)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&lit.User{ID: id, Name: name}, res))

}

// func (s *Suite) AfterTest(_, _ string) {
// 	require.NoError(s.T(), s.mock.ExpectationsWereMet())
// }

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
