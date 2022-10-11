package tickets

import (
	"context"

	"github.com/TomasCambiasso/desafio-goweb-tomascambiasso/internal/tickets/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	ps, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	ps, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	ps, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return -1, err
	}
	all, err := s.repository.GetAll(ctx)
	if err != nil {
		return -1, err
	}
	avg := float64(len(ps)) / float64((len(all)))
	return avg, nil
}
