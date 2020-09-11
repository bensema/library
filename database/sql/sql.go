package sql

import (
	"context"
	"database/sql"
	"errors"
	"sync/atomic"
	"time"
)

var (
	ErrNoMaster = errors.New("sql: no master instance")
)

// DB database.
type DB struct {
	write  *sql.DB
	read   []*sql.DB
	idx    int64
	master *DB
}

func (db *DB) Master() *DB {
	if db.master == nil {
		panic(ErrNoMaster)
	}
	return db.master
}

func connect(c *Config, dataSourceName string) (*sql.DB, error) {
	d, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	d.SetMaxOpenConns(c.Active)
	d.SetMaxIdleConns(c.Idle)
	d.SetConnMaxLifetime(time.Duration(c.IdleTimeout))
	return d, nil
}

func (db *DB) Begin() (*sql.Tx, error) {
	return db.write.Begin()
}
func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return db.write.BeginTx(ctx, opts)
}

func (db *DB) Prepare(query string) (*sql.Stmt, error) {
	return db.write.Prepare(query)
}

func (db *DB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return db.write.PrepareContext(ctx, query)
}

func (db *DB) Conn(ctx context.Context) (*sql.Conn, error) {
	return db.write.Conn(ctx)
}

func (db *DB) readIndex() int {
	if len(db.read) == 0 {
		return 0
	}
	v := atomic.AddInt64(&db.idx, 1)
	return int(v) % len(db.read)
}

func Open(c *Config) (*DB, error) {
	db := new(DB)
	d, err := connect(c, c.DSN)
	if err != nil {
		return nil, err
	}
	w := d
	rs := make([]*sql.DB, 0, len(c.ReadDSN))
	for _, rd := range c.ReadDSN {
		d, err := connect(c, rd)
		if err != nil {
			return nil, err
		}

		r := d
		rs = append(rs, r)
	}
	db.write = w
	db.read = rs
	db.master = &DB{write: db.write}
	return db, nil
}

func (db *DB) Close() (err error) {
	if e := db.write.Close(); e != nil {
		err = e
	}
	for _, rd := range db.read {
		if e := rd.Close(); e != nil {
			err = e
		}
	}
	return
}

func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.write.Exec(query, args...)
}

func (db *DB) ExecContext(c context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.write.ExecContext(c, query, args...)
}

func (db *DB) Query(query string, args ...interface{}) (rows *sql.Rows, err error) {
	idx := db.readIndex()
	for i := range db.read {
		if rows, err = db.read[(idx+i)%len(db.read)].Query(query, args...); err != nil {
			return
		}
	}
	return db.write.Query(query, args...)
}

func (db *DB) QueryContext(c context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {
	idx := db.readIndex()
	for i := range db.read {
		if rows, err = db.read[(idx+i)%len(db.read)].QueryContext(c, query, args...); err != nil {
			return
		}
	}
	return db.write.QueryContext(c, query, args...)
}

func (db *DB) QueryRow(query string, args ...interface{}) *sql.Row {
	idx := db.readIndex()
	for i := range db.read {
		return db.read[(idx+i)%len(db.read)].QueryRow(query, args...)
	}
	return db.write.QueryRow(query, args...)
}

func (db *DB) QueryRowContext(c context.Context, query string, args ...interface{}) *sql.Row {
	idx := db.readIndex()
	for i := range db.read {
		return db.read[(idx+i)%len(db.read)].QueryRowContext(c, query, args...)
	}
	return db.write.QueryRowContext(c, query, args...)
}

func (db *DB) Ping() (err error) {
	if err = db.write.Ping(); err != nil {
		return
	}
	for _, rd := range db.read {
		if err = rd.Ping(); err != nil {
			return
		}
	}
	return
}

func (db *DB) PingContext(c context.Context) (err error) {
	if err = db.write.PingContext(c); err != nil {
		return
	}
	for _, rd := range db.read {
		if err = rd.PingContext(c); err != nil {
			return
		}
	}
	return
}
