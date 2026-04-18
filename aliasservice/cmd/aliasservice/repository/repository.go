package repository

import "context"

type Repository struct {
	// db *sql.DB
}

func NewRepository( /* db */) *Repository {
	return &Repository{}
}

func (r *Repository) SaveURL(ctx context.Context, url string, alias string) error {
	// здесь потом будет сохранение в БД
	return nil
}
