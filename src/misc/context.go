package misc

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func NewContext(ctx context.Context, db *sqlx.DB) context.Context {
	return context.WithValue(ctx, "db", db)
}

func DBFromCtx(ctx context.Context) *sqlx.DB {
	return ctx.Value("db").(*sqlx.DB)
}
