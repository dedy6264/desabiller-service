package repositories

import (
	"context"
	"database/sql"
)

type Repositories struct {
	Db *sql.DB
	// DbMongo *mongo.Client
	ctx context.Context
}

func NewRepositories(Db *sql.DB,
	// DbMongo *mongo.Client,
	ctx context.Context) Repositories {
	return Repositories{
		Db: Db,
		// DbMongo: DbMongo,
		ctx: ctx,
	}
}
