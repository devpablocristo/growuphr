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

func (ns *NumberService) GetNumbers(ctx context.Context) (map[string]*domain.Number, error) {
	Numbers := ns.storage.ListNumbers(ctx)
	return Numbers, nil
}

func (ns *NumberService) GetNumber(ctx context.Context, UUID string) (*domain.Number, error) {
	p, err := ns.storage.GetNumber(ctx, UUID)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (ns *NumberService) AddNumber(ctx context.Context, n *domain.Number) (*domain.Number, error) {
	n.UUID = uuid.New().String()
	n.CreatedAt = time.Now()

	err := ns.storage.AddNumber(ctx, n)
	if err != nil {
		return &domain.Number{}, err
	}
	return n, nil

}

func (ns *NumberService) UpdateNumber(ctx context.Context, UUID string) error {
	return ns.storage.UpdateNumber(ctx, UUID)
}

func (ns *NumberService) DeleteNumber(ctx context.Context, UUID string) error {
	return ns.storage.DeleteNumber(ctx, UUID)
}
