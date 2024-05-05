package applications

import (
	"context"
	"errors"
	"fmt"

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

func (s *Service) CreateApplication(ctx context.Context, app CreateApplicationOptions) (*db.Application, error) {
	application, err := s.db.CreateApplication(ctx,
		db.CreateApplicationParams{
			ID:     cuid2.Generate(),
			Name:   app.Name,
			Status: "inactive",
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
