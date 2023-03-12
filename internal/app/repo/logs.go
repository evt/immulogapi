package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"

	"github.com/evt/immulogapi/internal/pkg/immudb"

	"github.com/evt/immulogapi/internal/app/models"
)

type ImmuLogRepo struct {
	db *immudb.DB
}

// NewImmuLogRepo creates a new instance of ImmuDB repository
func NewImmuLogRepo(c *immudb.DB) *ImmuLogRepo {
	return &ImmuLogRepo{db: c}
}

func (r *ImmuLogRepo) CreateLogLine(ctx context.Context, line *models.LogLine) error {
	log.Printf("repo: create log line: %v", line)

	_, err := r.db.ExecContext(ctx, "INSERT INTO loglines(id, source, text) VALUES (?, ?, ?)", line.ID, line.Source, line.Text)
	if err != nil {
		return fmt.Errorf("db.ExecContext failed: %w", err)
	}
	return nil
}

func (r *ImmuLogRepo) CreateLogLines(ctx context.Context, lines []*models.LogLine) error {
	log.Printf("repo: create a batch of log lines: %s", spew.Sdump(lines))

	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("db.Begin failed: %w", err)
	}
	defer tx.Rollback()

	for _, line := range lines {
		_, err := tx.ExecContext(ctx, "INSERT INTO loglines(id, source, text) VALUES (?, ?, ?)", line.ID, line.Source, line.Text)
		if err != nil {
			return fmt.Errorf("tx.ExecContext failed: %w", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("tx.Commit failed: %w", err)
	}

	return nil
}

func (r *ImmuLogRepo) GetLogLinesTotal(ctx context.Context) (uint64, error) {
	log.Printf("repo: get log line total")

	var total uint64
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) as total FROM loglines").Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("db.QueryRowContext failed: %w", err)
	}

	return total, nil
}

func prepareHistorySQLQuery(source string, limit int) (string, []any) {
	query := "SELECT id, source, text FROM loglines"

	var args []any
	if source != "" {
		query += " WHERE source = ?"
		args = append(args, source)
	}
	query += " order by id desc"
	if limit > 0 {
		query += " LIMIT ?"
		args = append(args, limit)
	}

	return query, args
}

func (r *ImmuLogRepo) GetLogLinesHistory(ctx context.Context, source string, limit int) ([]*models.LogLine, error) {
	log.Printf("repo: get log lines history")

	query, args := prepareHistorySQLQuery(source, limit)

	spew.Dump(query, args)

	var lines []*models.LogLine
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("db.QueryContext failed: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var line models.LogLine
		err = rows.Scan(&line.ID, &line.Source, &line.Text)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed: %w", err)
		}
		lines = append(lines, &line)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration failed: %w", err)
	}

	return lines, nil
}

func (r *ImmuLogRepo) CreateLogLinesTable(ctx context.Context) error {
	sql := `
CREATE TABLE IF NOT EXISTS loglines (
    id INTEGER,
    source VARCHAR[64],
    text VARCHAR,
    PRIMARY KEY (id)
);`
	_, err := r.db.ExecContext(ctx, sql)
	if err != nil {
		return fmt.Errorf("db.ExecContext failed: %w", err)
	}
	return nil
}
