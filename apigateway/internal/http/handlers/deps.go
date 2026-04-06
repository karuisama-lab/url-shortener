package handlers

import "url-shortener/apigateway/internal/http/handlers/aliashandlers"

type Deps struct {
	Alias *aliashandlers.AliasHandler
}

func NewDeps(alias *aliashandlers.AliasHandler) *Deps {
	return &Deps{
		Alias: alias,
	}
}
