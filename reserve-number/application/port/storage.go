package port

import (
	"context"

	domain "github.com/devpablocristo/growuphr/reserve-number/domain"
)

type Storage interface {
	AddNumber(context.Context, *domain.Number) error
	GetNumber(context.Context, string) (*domain.Number, error)
	ListNumbers(context.Context) map[string]*domain.Number
	DeleteNumber(context.Context, string) error
	UpdateNumber(context.Context, string) error
}
