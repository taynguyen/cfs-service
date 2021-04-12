package store

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
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

func (s *MySQLStore) SearchEventByTimeRange(o *SearchEventOptions) (*SearchEventResult, error) {
	selectQuery := `SELECT e.id, e.number, e.agency_id, e.type_code, e.created_time, e.dispatch_time, e.response_code `
	whereQuery := `
FROM events e
WHERE e.agency_id = ?
	AND e.created_time >= ?
	AND e.created_time <= ?
`
	orderByQuery := fmt.Sprintf(`
ORDER BY e.created_time %s
`, o.Order)
	offsetQuery := buildOffsetQuery(o.PagingOpts)

	args := []interface{}{o.AgentcyID, o.From, o.To}

	// count result
	countQuery := "SELECT COUNT(1) " + whereQuery + orderByQuery + offsetQuery
	count, err := doCountQuery(s.db, countQuery, args)
	if err != nil {
		return nil, err
	}

	// Query with data
	dataQueryStr := selectQuery + whereQuery + orderByQuery + offsetQuery
	rows, err := s.db.Query(dataQueryStr, args...)
	if err != nil {
		return nil, err
	}

	// Fetch events from query result
	events := make([]*Event, 0, count)
	for rows.Next() {
		e := &Event{}
		err := rows.Scan(&e.ID, &e.Number, &e.AgencyID, &e.TypeCode, &e.CreatedTime, &e.DispatchTime, &e.ResponderCode)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	return &SearchEventResult{
		Events: events,
		PageInfo: PagingInfo{
			Total: uint64(count),
			Offet: o.PagingOpts.Offset,
		},
	}, nil
}

func doCountQuery(db *sql.DB, query string, args []interface{}) (int64, error) {
	row := db.QueryRow(query, args...)
	if row.Err() != nil {
		return 0, row.Err()
	}
	var count int64
	row.Scan(&count)
	return count, nil
}

func buildOffsetQuery(o *PagingOptions) string {
	return fmt.Sprintf(`
LIMIT %d
OFFSET %d`, o.ItemsPerPage, o.Offset)
}
