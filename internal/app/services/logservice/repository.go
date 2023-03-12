package logservice

import (
	"context"

	"github.com/evt/immulogapi/internal/app/models"
)

type Repo interface {
	CreateLogLine(ctx context.Context, line *models.LogLine) error
	CreateLogLines(ctx context.Context, lines []*models.LogLine) error
	GetLogLinesTotal(ctx context.Context) (uint64, error)
	GetLogLinesHistory(ctx context.Context, source string, limit int) ([]*models.LogLine, error)
}
