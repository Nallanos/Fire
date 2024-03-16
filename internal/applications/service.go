package applications

import (
	"context"
	"database/sql"
	"errors"

	"github.com/nallanos/fire/internal/db"
	"github.com/nrednav/cuid2"
)

type Service struct {
	db db.Queries
}

func NewService(db db.Queries) *Service {
	return &Service{
		db: db,
	}
}

type CreateApplicationOptions struct {
	Name string
}

func (s *Service) CreateApplication(ctx context.Context, app CreateApplicationOptions) (*db.Application, error) {
	application, err := s.db.CreateApplication(ctx,
		db.CreateApplicationParams{
			ID:   cuid2.Generate(),
			Name: app.Name,
		})

	if err != nil {
		return nil, err
	}

	return &application, nil
}

func (s *Service) ListApplications(ctx context.Context) ([]db.Application, error) {
	applications, err := s.db.ListApplications(ctx)
	if err != nil {
		return nil, err
	}

	return applications, nil
}

var ErrApplicationNotFound = errors.New("application not found")

func (s *Service) GetApplication(ctx context.Context, id string) (*db.Application, error) {
	application, err := s.db.GetApplication(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrApplicationNotFound
		}
		return nil, err
	}

	return &application, nil
}
func (s *Service) DeleteApplication(ctx context.Context, id string) error {
	err := s.db.DeleteApplication(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
