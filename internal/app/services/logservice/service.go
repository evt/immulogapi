package logservice

import (
	"context"
	"log"
	"time"

	"github.com/evt/immulogapi/internal/app/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/evt/immulogapi/proto/v1"
)

type Service struct {
	v1.UnimplementedLogServiceServer
	repo Repo
}

func New(r Repo) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateLogLine(ctx context.Context, r *v1.CreateLogLineRequest) (*v1.CreateLogLineResponse, error) {
	log.Printf("service: create log line: %v", r)

	line := prepareLine(r)

	err := s.repo.CreateLogLine(ctx, line)
	if err != nil {
		log.Printf("failed to create log line in repository: %v", err)
		return nil, status.Error(codes.Internal, "failed to create log line in repository")
	}

	return &v1.CreateLogLineResponse{
		Id: line.ID,
	}, nil
}

func (s *Service) CreateLogLines(ctx context.Context, r *v1.CreateLogLinesRequest) (*v1.CreateLogLinesResponse, error) {
	log.Printf("service: create a batch of log lines: %v", r)

	lines := make([]*models.LogLine, 0, len(r.Lines))
	ids := make([]uint64, 0, len(r.Lines))
	for _, l := range r.Lines {
		line := prepareLine(l)
		lines = append(lines, line)
		ids = append(ids, line.ID)
	}

	err := s.repo.CreateLogLines(ctx, lines)
	if err != nil {
		log.Printf("failed to create log line in repository: %v", err)
		return nil, status.Error(codes.Internal, "failed to create log line in repository")
	}

	return &v1.CreateLogLinesResponse{
		Id: ids,
	}, nil
}

func (s *Service) GetLogLineTotal(ctx context.Context, r *emptypb.Empty) (*v1.GetLogLineTotalResponse, error) {
	log.Printf("service: get log line total")

	total, err := s.repo.GetLogLinesTotal(ctx)
	if err != nil {
		log.Printf("failed to get total log lines from repository: %v", err)
		return nil, status.Error(codes.Internal, "failed to get total log lines from repository")
	}

	return &v1.GetLogLineTotalResponse{
		Total: total,
	}, nil
}

func (s *Service) GetLogLinesHistory(ctx context.Context, r *v1.GetLogLinesHistoryRequest) (*v1.GetLogLinesHistoryResponse, error) {
	log.Printf("service: get log line history")

	lines, err := s.repo.GetLogLinesHistory(ctx, r.Source, int(r.Limit))
	if err != nil {
		log.Printf("failed to get log lines history from repository: %v", err)
		return nil, status.Error(codes.Internal, "failed to get log lines history from repository")
	}

	historyLines := toHistoryLines(lines)

	return &v1.GetLogLinesHistoryResponse{
		Lines: historyLines,
	}, nil
}

func prepareLine(r *v1.CreateLogLineRequest) *models.LogLine {
	createdAt := time.Now().UTC()
	if r.CreatedAt.IsValid() {
		createdAt = r.CreatedAt.AsTime()
	}

	id := createdAt.UnixNano()

	return &models.LogLine{
		ID:     uint64(id),
		Source: r.Source,
		Text:   r.Text,
	}
}

func toHistoryLines(lines []*models.LogLine) []*v1.LogLineHistoryRec {
	historyLines := make([]*v1.LogLineHistoryRec, 0, len(lines))
	for _, l := range lines {
		historyLines = append(historyLines, &v1.LogLineHistoryRec{
			Id:     l.ID,
			Source: l.Source,
			Text:   l.Text,
		})
	}
	return historyLines
}
