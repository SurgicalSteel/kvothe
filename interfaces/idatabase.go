package interfaces

//go:generate mockgen --destination=../mocks/mock_idatabase.go --package=mocks --source=idatabase.go
import (
	"context"
	"database/sql"

	"github.com/SurgicalSteel/kvothe/resources"

	"github.com/jmoiron/sqlx"
)

// IDatabase is interface for database
type IDatabase interface {
	ConnectDB(dbAccRead *resources.DBAccount, dbAccWrite *resources.DBAccount)
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	DriverName() string
	Close()
	Begin() (IDBTx, error)
	In(query string, params ...interface{}) (string, []interface{}, error)
	Rebind(query string) string
	Select(dest interface{}, query string, args ...interface{}) error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error

	// DB transaction.
	IDatabaseTx
	Commit() error
	Rollback() error

	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowSqlx(query string, args ...interface{}) *sqlx.Row
}

// IDatabaseTx is interface for only transaction database
type IDatabaseTx interface {
	BeginTx() (*sqlx.Tx, error)
	TransactionBlock(tx *sqlx.Tx, fc func(tx *sqlx.Tx) error) error
}

type IDBTx interface {
	Rollback()
	Commit()
}

type IRows interface {
	MapScan(dest map[string]interface{}) error
	SliceScan() ([]interface{}, error)
	StructScan(dest interface{}) error
	Next() bool
}
