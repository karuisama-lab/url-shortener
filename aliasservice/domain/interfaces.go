package domain

import "context"

type AliasInterface interface {
	SaveURL(reqDto aliasdto.URLSaveRequest, ctx context.Context) error
	//GetURL
}
