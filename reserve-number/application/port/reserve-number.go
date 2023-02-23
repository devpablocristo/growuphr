package port

import (
	"context"

	domain "github.com/devpablocristo/growuphr/reserve-number/domain"
)

type ReserveNumber interface {
	ReserveNumber(context.Context, *domain.Number, string) error
	ReservedNumbers(ctx context.Context) error
}
