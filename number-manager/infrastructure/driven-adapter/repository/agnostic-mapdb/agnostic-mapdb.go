package mapdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type MapDB struct {
	mDB map[string]any
}

func NewMapDB() *MapDB {
	m := make(map[string]any)
	return &MapDB{
		mDB: m,
	}
}

func (m *MapDB) Create(ctx context.Context, v any) error {
	UUID := uuid.New().String()
	m.mDB[UUID] = v
	_, err := m.Read(ctx, UUID)
	if err != nil {
		return err
	}
	return nil
}

func (m *MapDB) Read(ctx context.Context, UUID string) (any, error) {
	v, exist := m.mDB[UUID]
	if !exist {
		return nil, errors.New("value not found")
	}
	return v, nil
}

// func (m *MapDB) CheckForValue(value any) bool {
// 	for _, v := range m.mDB {
// 		if v == value {
// 			return true
// 		}
// 	}
// 	return false
// }

func (m *MapDB) List(ctx context.Context) map[string]any {
	// var results []domain.Number
	// for _, Number := range m.mDB {
	// 	results = append(results, *Number)
	// }
	// return results, nil

	fmt.Println("storage: ")
	fmt.Println(m)
	return m.mDB
}

func (m *MapDB) Delete(ctx context.Context, UUID string) error {
	delete(m.mDB, UUID)
	return nil
}

func (m *MapDB) Update(ctx context.Context, UUID string) error {
	return nil
}
