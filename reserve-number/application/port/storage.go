package port

import (
	"context"

	domain "github.com/devpablocristo/growuphr/reserve-number/domain"
)

type Storage interface {
	Create(context.Context, *domain.ReservedNumber) error
	Read(context.Context, string) (*domain.ReservedNumber, error)
	List(context.Context) map[string]*domain.ReservedNumber
	Delete(context.Context, string) error
	Update(context.Context, string) error
	CheckForUsername(string) (*domain.ReservedNumber, bool)
	CheckForNumber(int) (*domain.ReservedNumber, bool)
}
