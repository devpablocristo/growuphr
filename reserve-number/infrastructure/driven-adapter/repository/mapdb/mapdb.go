package mapdb

import (
	"context"

	domain "github.com/devpablocristo/growuphr/reserve-number/domain"
)

type MapDB struct {
	mDB map[string]*domain.Number
}

func NewMapDB() *MapDB {
	m := make(map[string]*domain.Number)
	return &MapDB{
		mDB: m,
	}
}

func (m *MapDB) AddNumber(ctx context.Context, p *domain.Number) error {
	m.mDB[p.UUID] = p
	return nil
}

func (m *MapDB) GetNumber(ctx context.Context, UUID string) (*domain.Number, error) {
	p := m.mDB[UUID]
	return p, nil
}

func (m *MapDB) ListNumbers(ctx context.Context) map[string]*domain.Number {
	// var results []domain.Number
	// for _, Number := range m.mDB {
	// 	results = append(results, *Number)
	// }
	// return results, nil

	return m.mDB
}

func (m *MapDB) DeleteNumber(ctx context.Context, UUID string) error {
	delete(m.mDB, UUID)
	return nil
}

func (m *MapDB) UpdateNumber(ctx context.Context, UUID string) error {
	return nil
}
