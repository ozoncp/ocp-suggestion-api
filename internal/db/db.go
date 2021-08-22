package db

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"google.golang.org/grpc"
)

var dbKey = "db"

// GetDBKey ...
func GetDBKey() string {
	return dbKey
}

// Connect ...
func Connect(DSN string) *sql.DB {
	db, err := sql.Open("pgx", DSN)
	if err != nil {
		log.Fatalf("connect do db error %v", err)
	}
	return db

}

// NewContext ...
func NewContext(ctx context.Context, db *sql.DB) context.Context {
	ctxDB := context.WithValue(ctx, &dbKey, db)

	return ctxDB
}

// FromContext ...
func FromContext(ctx context.Context) *sql.DB {
	client, ok := ctx.Value(&dbKey).(*sql.DB)
	if !ok {
		panic("Error getting connection from context")
	}
	return client
}

// GetDB ...
func GetDB(ctx context.Context) *sql.DB {
	return FromContext(ctx)
}

// NewInterceptorWithDB ...
func NewInterceptorWithDB(db *sql.DB) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		return handler(NewContext(ctx, db), req)
	}
}
