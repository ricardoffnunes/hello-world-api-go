package main

import (
    "testing"

    "go.etcd.io/bbolt"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

// MockDB is a mock for the bbolt DB
type MockDB struct {
    mock.Mock
}

func (m *MockDB) Update(fn func(*bbolt.Tx) error) error {
    args := m.Called(fn)
    return args.Error(0)
}

// Assuming createUserBucket is supposed to create a bucket in the bbolt DB, here's a simplified mock version for testing.
// In real application logic, you should replace this with the actual implementation.
func createUserBucket(db *MockDB) error {
    return db.Update(func(tx *bbolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists([]byte("Users"))
        return err
    })
}

func TestCreateUserBucket(t *testing.T) {
    mockDB := new(MockDB)
    mockDB.On("Update", mock.AnythingOfType("func(*bbolt.Tx) error")).Return(nil)

    err := createUserBucket(mockDB)
    assert.NoError(t, err)

    mockDB.AssertExpectations(t)
}