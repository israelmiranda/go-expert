package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string
	Name      string
	Timestamp time.Time
}

func CreateUser(name string) User {
	return User{
		ID:        uuid.NewString(),
		Name:      name,
		Timestamp: time.Now(),
	}
}
