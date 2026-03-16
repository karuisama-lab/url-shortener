package handlers

type Deps struct {
	Alias *AliasHandler
}

func NewDeps(alias *AliasHandler) *Deps {
	return &Deps{
		Alias: alias,
	}
}
