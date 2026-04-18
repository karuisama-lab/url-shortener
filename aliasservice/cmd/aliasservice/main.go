package main

repo := repository.NewRepository(...)
service := usecase.NewService(repo)
handler := handler.NewAliasHandler(service)
