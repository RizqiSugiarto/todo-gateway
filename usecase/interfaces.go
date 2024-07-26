package usecase

import (
	"context"

	"github.com/digisata/todo-gateway/domain"
)

type (
	UserRepository interface {
		GetByEmail(ctx context.Context, email string) (domain.User, error)
	}

	CacheRepository interface {
		Set(req domain.CacheItem) error
		Get(key string) (domain.CacheItem, error)
		Delete(key string) error
	}
)
