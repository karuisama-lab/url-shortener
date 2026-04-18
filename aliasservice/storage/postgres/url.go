package postgres

import "context"

type Postgres struct {
	// db *sql.DB
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (p *Postgres) SaveURL(ctx context.Context, url string, alias string) error {
	_ = ctx
	_ = url
	_ = alias

	// здесь потом будет реальное сохранение в БД
	return nil
}
