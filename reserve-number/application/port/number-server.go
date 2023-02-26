package port

import (
	"context"

	domain "github.com/devpablocristo/growuphr/reserve-number/domain"
)

type NumberReserver interface {
	AddReserveNumber(context.Context, *domain.ReservedNumber) error
	ReservedNumbers(ctx context.Context) error
}
