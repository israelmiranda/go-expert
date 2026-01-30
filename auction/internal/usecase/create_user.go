package usecase

import (
	"context"

	"github.com/israelmiranda/go-expert/auction/internal/domain"
)

type UserInput struct {
	Name string
}

func (u UserInput) toUser() domain.User {
	return domain.CreateUser(u.Name)
}

type CreateUser interface {
	Create(ctx context.Context, user domain.User) error
}

type CreateUserUseCase struct {
	repository CreateUser
}

func NewCreateUserUseCase(repository CreateUser) CreateUserUseCase {
	return CreateUserUseCase{repository}
}

func (u CreateUserUseCase) Create(ctx context.Context, userInput UserInput) error {
	return u.repository.Create(ctx, userInput.toUser())
}
