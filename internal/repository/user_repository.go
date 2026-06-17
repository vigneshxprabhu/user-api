package repository

import (
	"context"

	database "user-api/db/sqlc/generated"
)

type UserRepository struct {
	Queries *database.Queries
}

func NewUserRepository(q *database.Queries) *UserRepository {
	return &UserRepository{
		Queries: q,
	}
}
func (r *UserRepository) ListUsers() ([]database.User, error) {

	return r.Queries.ListUsers(context.Background())
}
func (r *UserRepository) GetUser(id int32) (database.User, error) {
	return r.Queries.GetUser(context.Background(), id)
}
func (r *UserRepository) CreateUser(
	params database.CreateUserParams,
) (database.User, error) {

	return r.Queries.CreateUser(
		context.Background(),
		params,
	)
}
func (r *UserRepository) UpdateUser(
	params database.UpdateUserParams,
) (database.User, error) {

	return r.Queries.UpdateUser(
		context.Background(),
		params,
	)
}
func (r *UserRepository) DeleteUser(id int32) error {
	return r.Queries.DeleteUser(
		context.Background(),
		id,
	)
}
