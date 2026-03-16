package aliasclient

import (
	"context"
	"url-shortener/aliasservice/domain"
	"url-shortener/apigateway/internal/http/dto/aliasdto"
)

type Client struct {
	service domain.AliasInterface //TODO: поправить/продолжить в самом сервисе
}

func NewClient(service domain.AliasInterface) *Client {
	return &Client{
		service: service,
	}
}

func (c *Client) SaveURL(req aliasdto.URLSaveRequest, ctx context.Context) error { //TODO: если бы нужно было что-то вернуть кроме err, правильно ли возвращать ссылку а принимать без ссылки? как лучше написать?
	return c.service.SaveURL(req, ctx)
}
