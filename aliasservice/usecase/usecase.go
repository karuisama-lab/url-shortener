package usecase

import (
	"context"
	"url-shortener/aliasservice/domain"
	"url-shortener/aliasservice/domain/entity"
)

type Service struct {
	repo domain.AliasRepository
}

func NewService(db DBInterface) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) SaveURL(reqDto aliasdto.URLSaveRequest, ctx context.Context) error {
	req := &entity.URLSaveRequest{
		UserID: reqDto.UserID,
		URL:    reqDto.URL,
	}
	err := s.db.SaveURL(req)
	if err != nil {
		return domain.ErrNotSave
	}

	return err
}
