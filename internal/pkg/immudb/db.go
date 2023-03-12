package immudb

import (
	"database/sql"

	"github.com/codenotary/immudb/pkg/stdlib"

	"github.com/codenotary/immudb/pkg/client"
	"github.com/evt/immulogapi/internal/app/config"
)

type DB struct {
	*sql.DB
}

func New(config *config.Config) (*DB, error) {
	opts := client.DefaultOptions().
		WithUsername(config.ImmuDB.Username).
		WithPassword(config.ImmuDB.Password).
		WithDatabase(config.ImmuDB.Database).
		WithPort(config.ImmuDB.Port).
		WithAddress(config.ImmuDB.Host)

	db := stdlib.OpenDB(opts)

	return &DB{db}, nil
}
