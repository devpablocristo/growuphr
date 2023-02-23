package port

import (
	"context"

	domain "github.com/devpablocristo/growuphr/reserve-number/domain"
)

type Number interface {
	GetNumbers(context.Context) (map[string]*domain.Number, error)
	GetNumber(context.Context, string) (*domain.Number, error)
	AddNumber(context.Context, *domain.Number) (*domain.Number, error)
	UpdateNumber(context.Context, string) error
	DeleteNumber(context.Context, string) error
}
