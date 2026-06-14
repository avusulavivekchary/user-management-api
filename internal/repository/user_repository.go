package repository

import (
	"github.com/avusulavivekchary/user-management-api/db/sqlc"
)

type UserRepository struct {
	Queries *sqlc.Queries
}

func NewUserRepository(q *sqlc.Queries) *UserRepository {
	return &UserRepository{
		Queries: q,
	}
}
