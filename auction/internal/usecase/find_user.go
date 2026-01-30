package usecase

import (
	"context"
	"time"

	"github.com/israelmiranda/go-expert/auction/internal/domain"
)

type UserOutput struct {
	ID        string
	Name      string
	Timestamp time.Time
}

func fromUsers(users []domain.User) []UserOutput {
	var outputs []UserOutput
	for _, user := range users {
		outputs = append(outputs, fromUser(user))
	}
	return outputs
}

func fromUser(user domain.User) UserOutput {
	return UserOutput{
		ID:        user.ID,
		Name:      user.Name,
		Timestamp: user.Timestamp,
	}
}

type FindUser interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindById(ctx context.Context, id string) (domain.User, error)
}

type FindUserUseCase struct {
	repository FindUser
}

func NewFindUserUseCase(repository FindUser) FindUserUseCase {
	return FindUserUseCase{repository}
}

func (u FindUserUseCase) FindAll(ctx context.Context) ([]UserOutput, error) {
	users, err := u.repository.FindAll(ctx)
	if err != nil {
		return []UserOutput{}, err
	}

	return fromUsers(users), nil
}

func (u FindUserUseCase) FindById(ctx context.Context, id string) (UserOutput, error) {
	user, err := u.repository.FindById(ctx, id)
	if err != nil {
		return UserOutput{}, err
	}

	return fromUser(user), nil
}
