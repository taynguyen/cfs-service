package store

import (
	"context"
	"database/sql"
	"embed"
	"time"

	"github.com/Boostport/migration"
	"github.com/Boostport/migration/driver/mysql"
	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLStore struct {
	connStr string
	db      *sql.DB
}

func NewMySQLStore(connStr string) (IStore, error) {
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	return &MySQLStore{
		connStr: connStr,
		db:      db,
	}, nil
}

//go:embed migrations
var embedMigrationsFS embed.FS

func (s *MySQLStore) Migrate() error {
	embedSource := &migration.EmbedMigrationSource{
		EmbedFS: embedMigrationsFS,
		Dir:     "migrations",
	}

	// Create driver
	driver, err := mysql.New(s.connStr)
	if err != nil {
		return err
	}

	// Run all up migrations
	applied, err := migration.Migrate(driver, embedSource, migration.Up, 0)
	if err != nil {
		return err
	}
	logrus.Info("Migration applied: ", applied)

	return nil
}

func (s *MySQLStore) AddEvent(e *Event) error {
	ctx := context.Background()
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// createdat := time.Now().Unix()
	_, err = tx.Exec(`INSERT INTO events (id, number, agency_id, type_code, created_time, dispatch_time, response_code, inserted_at) values (?,?,?,?,?,?,?,?)`,
		e.ID, e.Number, e.AgencyID, e.TypeCode, e.CreatedTime, e.DispatchTime, e.ResponderCode, time.Now())
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
