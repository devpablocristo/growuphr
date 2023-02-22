package application

import (
	"context"
	"time"

	uuid "github.com/google/uuid"

	port "github.com/devpablocristo/growuphr/reserve-number/application/port"
	domain "github.com/devpablocristo/growuphr/reserve-number/domain"
)

type NumberService struct {
	storage port.Storage
}

func NewNumberService(s port.Storage) *NumberService {
	return &NumberService{
		storage: s,
	}
}

func (ps *NumberService) GetNumbers(ctx context.Context) (map[string]*domain.Number, error) {
	Numbers := ps.storage.ListNumbers(ctx)
	return Numbers, nil
}

func (ps *NumberService) GetNumber(ctx context.Context, UUID string) (*domain.Number, error) {
	p, err := ps.storage.GetNumber(ctx, UUID)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (ps *NumberService) AddNumber(ctx context.Context, n *domain.Number) (*domain.Number, error) {
	n.UUID = uuid.New().String()
	n.CreatedAt = time.Now()

	err := ps.storage.AddNumber(ctx, n)
	if err != nil {
		return &domain.Number{}, err
	}
	return n, nil
}

func (ps *NumberService) UpdateNumber(ctx context.Context, UUID string) error {
	return ps.storage.UpdateNumber(ctx, UUID)
}

func (ps *NumberService) DeleteNumber(ctx context.Context, UUID string) error {
	return ps.storage.DeleteNumber(ctx, UUID)
}
