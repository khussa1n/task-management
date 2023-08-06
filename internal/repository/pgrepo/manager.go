package pgrepo

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	usersTable        = "users"
	statusesTable     = "statuses"
	prioritiesTable   = "priorities"
	tasksTable        = "tasks"
	rolesTable        = "roles"
	membersTasksTable = "members_tasks"
	actionsTable      = "actions"
	eventsTable       = "events"
	taskLogsTable     = "task_logs"
	commentsTable     = "comments"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Postgres {
	return &Postgres{
		Pool: pool,
	}
}
