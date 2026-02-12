package misc

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type ctxKey struct{}

var dbKey ctxKey

func NewContext(ctx context.Context, db *sqlx.DB) context.Context {
	return context.WithValue(ctx, dbKey, db)
}

func DBFromCtx(ctx context.Context) *sqlx.DB {
	return ctx.Value(dbKey).(*sqlx.DB)
}
