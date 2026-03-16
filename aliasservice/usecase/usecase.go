package usecase

import (
	"context"
	"url-shortener/aliasservice/domain"
	"url-shortener/aliasservice/domain/entity"
)

type Service struct {
	db DBInterface //TODO: интерфейс подменяем постгрес, pgxpool. Добавить интерфейс(+ вопрос где он должен лежать, в этом же файле? или тоже завести файл interfaces/ports отдельный в этой папке?)
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
