package mapdb

import (
	"context"
	"errors"

	"github.com/devpablocristo/growuphr/number-manager/domain"
)

type MapDB struct {
	mDB map[string]*domain.ReservedNumber
}

func NewMapDB() *MapDB {
	m := make(map[string]*domain.ReservedNumber)
	return &MapDB{
		mDB: m,
	}
}

func (m *MapDB) Create(ctx context.Context, rn *domain.ReservedNumber) error {
	m.mDB[rn.UUID] = rn
	_, err := m.Read(ctx, rn.UUID)
	if err != nil {
		return err
	}
	//fmt.Println(m.mDB[rn.UUID])

	return nil
}

func (m *MapDB) Read(ctx context.Context, UUID string) (*domain.ReservedNumber, error) {
	rn, exist := m.mDB[UUID]
	if !exist {
		return nil, errors.New("value not found")
	}
	return rn, nil
}

func (m *MapDB) CheckForUsername(ctx context.Context, checkUsr string) (*domain.ReservedNumber, bool) {
	for _, rn := range m.mDB {
		if rn.User.Username == checkUsr {
			return rn, true
		}
	}
	return nil, false
}

func (m *MapDB) CheckForNumber(ctx context.Context, checkNum int) (*domain.ReservedNumber, bool) {
	for _, rn := range m.mDB {
		if rn.Number.Number == checkNum {
			return rn, true
		}
	}
	return nil, false
}

func (m *MapDB) List(ctx context.Context) map[string]*domain.ReservedNumber {
	return m.mDB
}

func (m *MapDB) Delete(ctx context.Context, UUID string) error {
	delete(m.mDB, UUID)
	return nil
}

func (m *MapDB) Update(ctx context.Context, UUID string) error {
	return nil
}
