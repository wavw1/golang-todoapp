package tasks_postgres_repository

import (
	core_postgres_pool "github.com/wavw1/golang-todoapp/internal/core/repository/postgres/pool"
)

type TasksRepository struct {
	pool *core_postgres_pool.ConnectionPool
}

func NewTaskRepository(
	pool *core_postgres_pool.ConnectionPool,
) *TasksRepository {
	return &TasksRepository{
		pool: pool,
	}
}
