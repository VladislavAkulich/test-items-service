package client

import (
	"context"
	pgx2 "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Client interface {
	Begin(context.Context) (pgx2.Tx, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx2.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx2.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}
