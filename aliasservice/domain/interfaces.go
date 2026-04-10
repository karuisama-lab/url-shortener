package domain

import "context"

type AliasRepository interface {
	SaveURL(ctx context.Context, url string, alias string) error
}

type AliasInterface interface {
	SaveURL(ctx context.Context, url string) (string, error)
}
