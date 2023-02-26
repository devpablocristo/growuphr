package port

import (
	"context"

	domain "github.com/devpablocristo/growuphr/reserve-number/domain"
)

type Number interface {
	GetNumbers(context.Context) (map[string]*domain.ReservedNumber, error)
	GetNumber(context.Context, string) (*domain.ReservedNumber, error)
	AddNumber(context.Context, *domain.ReservedNumber) (*domain.ReservedNumber, error)
	UpdateNumber(context.Context, string) error
	DeleteNumber(context.Context, string) error
}
