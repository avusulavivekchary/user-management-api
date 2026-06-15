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

// ADD THIS BELOW CreateUser
func (s *UserService) ListUsers() ([]sqlc.User, error) {
	return s.repo.Queries.ListUsers(context.Background())
}
