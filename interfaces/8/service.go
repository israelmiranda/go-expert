package main

import "fmt"

type NotFoundError struct {
	Resource string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s not found\n", e.Resource)
}

type Database interface {
	GetUser(id int) (string, error)
	SaveUser(id int, name string) error
}

type UserService struct {
	db Database
}

func (s *UserService) RenameUser(id int, newName string) error {
	name, err := s.db.GetUser(id)
	if err != nil {
		return err
	}

	if name == newName {
		return nil // No change needed
	}

	return s.db.SaveUser(id, newName)
}
