package usecase

import (
	"context"
	"url-shortener/aliasservice/domain"
)

type Service struct {
	repo domain.AliasRepository
}

func NewService(repo domain.AliasRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) SaveURL(ctx context.Context, url string) (string, error) {
	alias := "some-alias"
	// заменить заглушку some-alias на реальную генерацию alias
	err := s.repo.SaveURL(ctx, url, alias)
	if err != nil {
		return "", domain.ErrNotSave
	}

	return alias, nil
}
