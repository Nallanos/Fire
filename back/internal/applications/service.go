package applications

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/nallanos/fire/internal/db"
	"github.com/nallanos/fire/internal/deployment"
	"github.com/nrednav/cuid2"
)

type Service struct {
	db         *db.Queries
	deployment *deployment.ContainerService
}

func NewService(db *db.Queries, s *deployment.ContainerService) *Service {
	return &Service{
		db:         db,
		deployment: s,
	}
}

type CreateApplicationOptions struct {
	Name string
}

func (s *Service) getRandomPortInRange(start, end int) int {
	apps, err := s.ListApplications(context.Background())
	if err != nil {
		slog.Error("error while listing applications: %w", err)
	}

	usedPorts := make(map[int]bool)
	for _, element := range apps {
		portInt, err := strconv.Atoi(element.Port)
		if err != nil {
			slog.Error("error while converting appPort into int : %w", err)
		} else {
			usedPorts[portInt] = true
		}
	}

	for i := start; i <= end; i++ {
		if !usedPorts[i] {
			return i
		}
	}

	return -1
}

func (s *Service) CreateApplication(ctx context.Context, app CreateApplicationOptions) (*db.Application, error) {
	randomPort := s.getRandomPortInRange(3000, 4000)
	application, err := s.db.CreateApplication(ctx,
		db.CreateApplicationParams{
			ID:     cuid2.Generate(),
			Name:   app.Name,
			Status: "inactive",
			Port:   fmt.Sprintf("%d", randomPort),
		})
	if err != nil {
		return nil, fmt.Errorf("error while creating application: %w", err)
	}
	return &application, nil
}

func (s *Service) ListApplications(ctx context.Context) ([]db.Application, error) {
	applications, err := s.db.ListApplications(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while listing applications: %w", err)
	}

	return applications, nil
}

var ErrApplicationNotFound = errors.New("application not found")

func (s *Service) GetApplication(ctx context.Context, id string) (*db.Application, error) {
	application, err := s.db.GetApplication(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error while getting application: %w", err)
	}

	return &application, nil
}
func (s *Service) DeleteApplication(ctx context.Context, id string) error {
	err := s.db.DeleteApplication(ctx, id)
	if err != nil {
		return fmt.Errorf("error while deleting application: %w", err)
	}
	return nil
}
