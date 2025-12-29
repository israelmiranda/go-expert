package main

import "testing"

type MockDB struct {
	users map[int]string
}

func (m *MockDB) GetUser(id int) (string, error) {
	name, exists := m.users[id]
	if !exists {
		return "", NotFoundError{Resource: "user"}
	}
	return name, nil
}

func (m *MockDB) SaveUser(id int, name string) error {
	m.users[id] = name
	return nil
}

func TestRenameUser(t *testing.T) {
	mockDB := &MockDB{
		users: map[int]string{1: "Alice"},
	}

	service := UserService{db: mockDB}

	// Test successful rename
	err := service.RenameUser(1, "Alicia")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if mockDB.users[1] != "Alicia" {
		t.Errorf("Expected name to be Alicia, got %s", mockDB.users[1])
	}

	// Test user not found
	err = service.RenameUser(999, "Nobody")
	if _, ok := err.(NotFoundError); !ok {
		t.Errorf("Expected NotFoundError, got %v", err)
	}
}
