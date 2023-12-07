package psql

import (
	"github.com/jmoiron/sqlx"
	"todoapp/config"

	_ "github.com/lib/pq"
)

const (
	defaultMaxConnLifeTime = 0
	defaultMaxConnIdleTime = 0
	defaultMaxIdleConns    = 2
	defaultMaxOpenConns    = 0
)

type Database struct {
	*sqlx.DB
}

func NewDatabase(config *config.Schema) *Database {
	db, err := sqlx.Connect(config.Vendor.Database.Driver, config.Vendor.Database.DataSource)
	if err != nil {
		panic(err)
	}
	maxConnLifeTime := config.Vendor.Database.MaxConnLifeTime
	if maxConnLifeTime == 0 {
		maxConnLifeTime = defaultMaxConnLifeTime
	}

	maxConnIdleTime := config.Vendor.Database.MaxConnIdleTime
	if maxConnIdleTime == 0 {
		maxConnIdleTime = defaultMaxConnIdleTime
	}

	maxIdleConns := config.Vendor.Database.MaxIdleConns
	if maxIdleConns == 0 {
		maxIdleConns = defaultMaxIdleConns
	}

	maxOpenConns := config.Vendor.Database.MaxOpenConns
	if maxOpenConns == 0 {
		maxOpenConns = defaultMaxOpenConns
	}

	db.SetConnMaxLifetime(maxConnLifeTime)
	db.SetConnMaxIdleTime(maxConnIdleTime)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)

	return &Database{db}
}
