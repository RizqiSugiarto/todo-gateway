package controller

import (
	"context"

	"github.com/digisata/todo-gateway/domain"
)

type (
	UserUseCase interface {
		Login(ctx context.Context, req domain.User) (domain.AuthResponse, error)
		Logout(ctx context.Context) error
	}
)
