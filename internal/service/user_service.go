package service

import (
	"context"
	"time"

	"github.com/avusulavivekchary/user-management-api/db/sqlc"
	"github.com/avusulavivekchary/user-management-api/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(name string, dob string) (sqlc.User, error) {
	parsedDOB, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return sqlc.User{}, err
	}

	var date pgtype.Date
	date.Time = parsedDOB
	date.Valid = true

	return s.repo.Queries.CreateUser(
		context.Background(),
		sqlc.CreateUserParams{
			Name: name,
			Dob:  date,
		},
	)
}

func (s *UserService) ListUsers(limit, offset int32) ([]sqlc.User, error) {
	return s.repo.Queries.ListUsers(
		context.Background(),
		sqlc.ListUsersParams{
			Limit:  limit,
			Offset: offset,
		},
	)
}
func (s *UserService) GetUserByID(id int32) (sqlc.User, error) {
	return s.repo.Queries.GetUserByID(
		context.Background(),
		id,
	)
}
func (s *UserService) UpdateUser(id int32, name string, dob string) (sqlc.User, error) {
	parsedDOB, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return sqlc.User{}, err
	}

	var date pgtype.Date
	date.Time = parsedDOB
	date.Valid = true

	return s.repo.Queries.UpdateUser(
		context.Background(),
		sqlc.UpdateUserParams{
			ID:   id,
			Name: name,
			Dob:  date,
		},
	)
}
func (s *UserService) DeleteUser(id int32) error {
	return s.repo.Queries.DeleteUser(
		context.Background(),
		id,
	)
}
