package mapdb

import (
	"context"
	"errors"

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

func (m *MapDB) GetUser(ctx context.Context, userName string) (*domain.Number, error) {
	u, exist := m.mDB[userName]
	if !exist {
		return &domain.Number{}, errors.New("username not found")
	}
	return u, nil
}

func (m *MapDB) AddUser(ctx context.Context, n *domain.Number) error {
	m.mDB[n.UserName] = n
	return nil
}

// function to check if a value present in the map
func (m *MapDB) checkForValue(userValue int, students map[string]int) bool {

	//traverse through the map
	for _, value := range students {

		//check if present value is equals to userValue
		if value == userValue {

			//if same return true
			return true
		}
	}

	//if value not found return false
	return false
}

func (m *MapDB) GetNumber(ctx context.Context, n int) (*domain.Number, error) {
	u := m.mDB["n"]
	return u, nil
}

func (m *MapDB) GetNumberByUUID(ctx context.Context, UUID string) (*domain.Number, error) {
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
