package domain

import "context"

type AliasInterface interface {
	SaveURL(reqDto aliasdto.URLSaveRequest, ctx context.Context) (respDto *aliasdto.URLSaveResponse, err error)
	//GetURL
}
