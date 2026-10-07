package store

import (
	"context"
	"errors"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrTooMany          = errors.New("too many records found")
	ErrInvalid          = errors.New("invalid argument")
	ErrUniqueConstraint = errors.New("unique constraint violation")
)

type Stores struct {
	Projects ProjectStore
}

type Client interface {
	Transact(ctx context.Context, fn func(ctx context.Context, stores *Stores) error) error
	Stores() *Stores
}

type ProjectStore interface {
	Count(ctx context.Context) (int, error)
}
