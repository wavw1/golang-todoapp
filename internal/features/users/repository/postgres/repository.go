package users_postgres_repository

import core_postgres_pool "github.com/wavw1/golang-todoapp/internal/features/users/repository/postgres/pool"

type UsersRepository struct {
	pool core_postgres_pool.Pool
}

func NewUsersRepository(
	pool core_postgres_pool.Pool,
) *UsersRepository {
	return &UsersRepository{
		pool: pool,
	}
}
