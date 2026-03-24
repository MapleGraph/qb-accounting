package postgres

import (
	"context"

	qbpostgres "github.com/MapleGraph/qb-core/v2/pkg/postgres"
)

func readContext(ctx context.Context) context.Context {
	if qbpostgres.GetTransaction(ctx) != nil {
		return ctx
	}
	return qbpostgres.WithRead(ctx)
}

func writeContext(ctx context.Context) context.Context {
	if qbpostgres.GetTransaction(ctx) != nil {
		return ctx
	}
	return qbpostgres.WithWrite(ctx)
}
